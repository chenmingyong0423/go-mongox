// Copyright 2024 chenmingyong0423

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
	"time"

	"github.com/chenmingyong0423/go-mongox/v2/field"

	"github.com/chenmingyong0423/go-mongox/v2/operation"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var strategies = map[operation.OpType]func(dest any, currentTime time.Time, fields []*field.Filed, opts ...any) error{
	operation.OpTypeBeforeInsert: beforeInsert,
	operation.OpTypeBeforeUpdate: beforeUpdate,
	operation.OpTypeBeforeUpsert: beforeUpsert,
}

func beforeInsert(dest any, currentTime time.Time, fields []*field.Filed, _ ...any) error {
	if v, ok := dest.(reflect.Value); ok {
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		return processFields4Insert(v, currentTime, fields)
	}
	return nil
}

func processFields4Insert(dest reflect.Value, currentTime time.Time, fields []*field.Filed) error {
	for idx, fd := range fields {
		value := dest.Field(idx)
		if fd.InlinedFields != nil {
			err := processFields4Insert(value, currentTime, fd.InlinedFields)
			if err != nil {
				return err
			}
		} else {
			if fd.AutoID {
				if value.IsZero() {
					value.Set(reflect.ValueOf(bson.NewObjectID()))
				}
			} else {
				handleTimeField(value, fd, currentTime)
			}
		}
	}
	return nil
}

// 设置时间字段
func handleTimeField(dest reflect.Value, fd *field.Filed, currentTime time.Time) {
	switch {
	case fd.AutoCreateTime != 0:
		setTimeField(dest, fd.AutoCreateTime, currentTime, fd.FieldType)
	case fd.AutoUpdateTime != 0:
		setTimeField(dest, fd.AutoUpdateTime, currentTime, fd.FieldType)
	}
}

// 设置具体的时间值
func setTimeField(dest reflect.Value, timeType field.TimeType, currentTime time.Time, fieldType reflect.Type) {
	if !dest.IsZero() {
		return
	}
	switch timeType {
	case field.UnixTime:
		dest.Set(reflect.ValueOf(currentTime))
	case field.UnixSecond:
		switch fieldType.Kind() {
		case reflect.Int:
			dest.Set(reflect.ValueOf(int(currentTime.Unix())))
		default:
			dest.Set(reflect.ValueOf(currentTime.Unix()))
		}
	case field.UnixMillisecond:
		dest.Set(reflect.ValueOf(currentTime.UnixMilli()))
	case field.UnixNanosecond:
		dest.Set(reflect.ValueOf(currentTime.UnixNano()))
	}
}

func beforeUpdate(dest any, currentTime time.Time, fields []*field.Filed, _ ...any) error {
	updates, ok := dest.(bson.M)
	if !ok || updates == nil {
		return nil
	}

	setFields, ok := updates["$set"].(bson.M)
	if !ok {
		return nil
	}

	updatedFields := findAdditionalFields(currentTime, fields, findUpdatedFields)

	for k, v := range updatedFields {
		if _, exit := setFields[k]; !exit {
			setFields[k] = v
		}
	}

	return nil
}

func beforeUpsert(dest any, currentTime time.Time, fields []*field.Filed, _ ...any) error {
	updates, ok := dest.(bson.M)
	if !ok || updates == nil {
		return nil
	}
	setFields, ok := updates["$set"].(bson.M)
	if !ok {
		return nil
	}

	updatedTimes := findAdditionalFields(currentTime, fields, findUpdatedFields)

	for k, v := range updatedTimes {
		if _, exit := setFields[k]; !exit {
			setFields[k] = v
		}
	}

	idAndCreateFields := findAdditionalFields(currentTime, fields, findUpsertFields)
	if len(idAndCreateFields) > 0 {
		if updates["$setOnInsert"] == nil {
			updates["$setOnInsert"] = bson.M{}
		}

		if setOnInsertFields, ok := updates["$setOnInsert"].(bson.M); ok {
			for k, v := range idAndCreateFields {
				if _, exit := setOnInsertFields[k]; !exit {
					setOnInsertFields[k] = v
				}
			}
		}
	}

	return nil
}

// 通用字段处理
func findAdditionalFields(currentTime time.Time, fields []*field.Filed, handler func(field *field.Filed, currentTime time.Time) (string, any)) map[string]any {
	result := make(map[string]any, len(fields))
	for _, fd := range fields {
		if fd.InlinedFields != nil {
			inlinedFields := findAdditionalFields(currentTime, fd.InlinedFields, handler)
			for k, v := range inlinedFields {
				result[k] = v
			}
		} else {
			if key, value := handler(fd, currentTime); key != "" {
				result[key] = value
			}
		}
	}
	return result
}

func findUpsertFields(fd *field.Filed, currentTime time.Time) (string, any) {
	if fd.AutoID {
		return fd.MongoField, bson.NewObjectID()
	}

	if fd.AutoCreateTime != 0 {
		return fd.MongoField, getTimeValue(fd.AutoCreateTime, currentTime)
	}
	return "", nil
}

func findUpdatedFields(fd *field.Filed, currentTime time.Time) (string, any) {
	if fd.AutoUpdateTime != 0 {
		return fd.MongoField, getTimeValue(fd.AutoUpdateTime, currentTime)
	}
	return "", nil
}

func getTimeValue(timeType field.TimeType, currentTime time.Time) any {
	switch timeType {
	case field.UnixTime:
		return currentTime
	case field.UnixSecond:
		return currentTime.Unix()
	case field.UnixMillisecond:
		return currentTime.UnixMilli()
	case field.UnixNanosecond:
		return currentTime.UnixNano()
	}
	return nil
}
