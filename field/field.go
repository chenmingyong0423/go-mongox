// Copyright 2025 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package field

import (
	"reflect"
	"strings"
	"time"
)

type Filed struct {
	Name string
	// the field name in mongo
	MongoField     string
	AutoID         bool
	FieldType      reflect.Type
	AutoCreateTime TimeType
	AutoUpdateTime TimeType

	InlinedFields []*Filed
}

type (
	// TimeType MONGOX time type
	TimeType int64
)

// Mongox time types
const (
	UnixTime        TimeType = 1
	UnixSecond      TimeType = 2
	UnixMillisecond TimeType = 3
	UnixNanosecond  TimeType = 4
)

const (
	CreatedAt      = "CreatedAt"
	UpdatedAt      = "UpdatedAt"
	AutoCreateTime = "autoCreateTime"
	AutoUpdateTime = "autoUpdateTime"
)

var (
	timeType = reflect.TypeOf(time.Time{})
)

func ParseFields[T any](doc T) []*Filed {
	docType := reflect.TypeOf(doc)
	if docType == nil {
		return nil
	}
	if docType.Kind() == reflect.Ptr {
		docType = docType.Elem()
	}
	if docType.Kind() != reflect.Struct {
		return nil
	}
	numField := docType.NumField()
	fields := make([]*Filed, 0, numField)

	for i := 0; i < numField; i++ {
		structField := docType.Field(i)
		fd := &Filed{Name: structField.Name, FieldType: structField.Type}

		bsonTag := structField.Tag.Get("bson")
		if structField.Anonymous {
			if bsonTag == ",inline" {
				fields = append(fields, &Filed{Name: structField.Name, FieldType: structField.Type, InlinedFields: ParseFields(reflect.New(structField.Type).Elem().Interface())})
				continue
			}
		}

		fd.MongoField = getMongoField(bsonTag, structField.Name)

		tag := structField.Tag.Get("mongox")
		if len(tag) > 0 {
			parseTag(tag, fd)
		} else if structField.Name == CreatedAt {
			parseDefaultTimeType(structField, fd, func(timeType TimeType) {
				fd.AutoCreateTime = timeType
			})
		} else if structField.Name == UpdatedAt {
			parseDefaultTimeType(structField, fd, func(timeType TimeType) {
				fd.AutoUpdateTime = timeType
			})
		}

		fields = append(fields, fd)
	}

	return fields
}

func parseDefaultTimeType(structField reflect.StructField, fd *Filed, set func(timeType TimeType)) {
	switch structField.Type.Kind() {
	case reflect.Struct:
		if structField.Type == timeType {
			set(UnixTime)
		}
	case reflect.Int64, reflect.Int:
		set(UnixSecond)
	default:
	}
}

func getMongoField(bsonTag string, defaultValue string) string {
	if bsonTag == "" {
		return defaultValue
	}
	split := strings.Split(bsonTag, ",")
	if split[0] == "" {
		return defaultValue
	}
	return split[0]
}

func parseTag(tag string, fd *Filed) {
	split := strings.Split(tag, ",")
	for _, s := range split {
		switch {
		case s == "autoID":
			fd.AutoID = true
		case strings.HasPrefix(s, AutoCreateTime):
			fd.AutoCreateTime = parseTimeType(s)
		case strings.HasPrefix(s, AutoUpdateTime):
			fd.AutoUpdateTime = parseTimeType(s)
		}
	}
}

func parseTimeType(tag string) TimeType {
	if strings.Contains(tag, ":") {
		timeType := strings.Split(tag, ":")[1]
		switch timeType {
		case "milli":
			return UnixMillisecond
		case "nano":
			return UnixNanosecond
		case "second":
			return UnixSecond
		default:
			return 0
		}
	}
	return UnixSecond
}
