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
		if v.Kind() == reflect.Pointer {
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

// setTimeField sets a zero-valued auto time field.
func setTimeField(dest reflect.Value, timeType field.TimeType, currentTime time.Time, fieldType reflect.Type) {
	if !dest.IsZero() {
		return
	}
	value, ok := getTimeValue(timeType, currentTime, fieldType)
	if !ok {
		return
	}
	reflectValue := reflect.ValueOf(value)
	if reflectValue.Type().AssignableTo(dest.Type()) {
		dest.Set(reflectValue)
		return
	}
	if reflectValue.Type().ConvertibleTo(dest.Type()) {
		dest.Set(reflectValue.Convert(dest.Type()))
	}
}

func beforeUpdate(dest any, currentTime time.Time, fields []*field.Filed, _ ...any) error {
	updates, exist := dest.(bson.M)
	if !exist || updates == nil {
		return nil
	}

	updatedFields := findAdditionalFields(currentTime, fields, findUpdatedFields)
	if len(updatedFields) > 0 {
		if updates["$set"] == nil {
			updates["$set"] = bson.M{}
		}

		setFields, ok := updates["$set"].(bson.M)
		if ok {
			for k, v := range updatedFields {
				if _, exit := setFields[k]; !exit {
					setFields[k] = v
				}
			}
		}
	}

	return nil
}

func beforeUpsert(dest any, currentTime time.Time, fields []*field.Filed, _ ...any) error {
	updates, exist := dest.(bson.M)
	if !exist || updates == nil {
		return nil
	}

	updatedTimes := findAdditionalFields(currentTime, fields, findUpdatedFields)

	if len(updatedTimes) > 0 {
		if updates["$set"] == nil {
			updates["$set"] = bson.M{}
		}

		setFields, ok := updates["$set"].(bson.M)
		if ok {
			for k, v := range updatedTimes {
				if _, exit := setFields[k]; !exit {
					setFields[k] = v
				}
			}
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

// FilterNotDeleted adds the model's deletion scope without changing the caller's
// filter. Missing fields, nulls and zero times all represent live documents.
// The filter must have been validated by Execute.
func FilterNotDeleted(filter any, fields []*field.Filed) any {
	return filterByDeleteScope(filter, fields, operation.SoftDeleteScopeNotDeleted)
}

func filterByDeleteScope(filter any, fields []*field.Filed, scope operation.SoftDeleteScope) any {
	if scope == operation.SoftDeleteScopeWithDeleted {
		return filter
	}
	conditions := bson.D{}
	visitDeletedFields(fields, func(fd *field.Filed) {
		zero := reflect.Zero(fd.FieldType).Interface()
		conditions = append(conditions, bson.E{Key: fd.MongoField, Value: bson.D{
			{Key: "$in", Value: bson.A{nil, zero}},
		}})
	})
	if scope == operation.SoftDeleteScopeOnlyDeleted {
		// Negate the entire live scope, including missing/null/zero values.
		// Without deletion fields, the live scope is {}, so no documents match.
		conditions = bson.D{{Key: "$nor", Value: bson.A{conditions}}}
	} else if len(conditions) == 0 {
		return filter
	}
	// Like mongo.Collection, interpret byte slice filters as raw BSON.
	// Invalid filters are left for the driver/server to reject, not repaired.
	if raw, ok := filter.([]byte); ok {
		filter = bson.Raw(raw)
	}
	return bson.D{{Key: "$and", Value: bson.A{filter, conditions}}}
}

func visitDeletedFields(fields []*field.Filed, visit func(*field.Filed)) {
	for _, fd := range fields {
		if fd.InlinedFields != nil {
			visitDeletedFields(fd.InlinedFields, visit)
		} else if fd.AutoDeleteTime != 0 {
			visit(fd)
		}
	}
}

func beforeDelete(opCtx *operation.OpContext) error {
	setFields := findAdditionalFields(opCtx.StartTime, opCtx.Fields, func(fd *field.Filed, now time.Time) (string, any) {
		if fd.AutoDeleteTime != 0 {
			if value, ok := getTimeValue(fd.AutoDeleteTime, now, fd.FieldType); ok {
				return fd.MongoField, value
			}
		}
		return "", nil
	})
	opCtx.SoftDelete = len(setFields) > 0
	if !opCtx.SoftDelete {
		return nil
	}
	// Prepare the update document without executing update callbacks.
	opCtx.Updates = bson.M{"$set": bson.M(setFields)}
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
		value, ok := getTimeValue(fd.AutoCreateTime, currentTime, fd.FieldType)
		if ok {
			return fd.MongoField, value
		}
	}
	return "", nil
}

func findUpdatedFields(fd *field.Filed, currentTime time.Time) (string, any) {
	if fd.AutoUpdateTime != 0 {
		value, ok := getTimeValue(fd.AutoUpdateTime, currentTime, fd.FieldType)
		if ok {
			return fd.MongoField, value
		}
	}
	return "", nil
}

func getTimeValue(timeType field.TimeType, currentTime time.Time, fieldType reflect.Type) (any, bool) {
	if timeType == 0 {
		return nil, false
	}
	switch fieldType.Kind() {
	case reflect.Struct:
		if fieldType == reflect.TypeFor[time.Time]() && timeType == field.UnixTime {
			return currentTime, true
		}
	case reflect.Int:
		switch timeType {
		case field.UnixSecond:
			return int(currentTime.Unix()), true
		case field.UnixMillisecond:
			return int(currentTime.UnixMilli()), true
		case field.UnixNanosecond:
			return int(currentTime.UnixNano()), true
		}
	case reflect.Int64:
		switch timeType {
		case field.UnixSecond:
			return currentTime.Unix(), true
		case field.UnixMillisecond:
			return currentTime.UnixMilli(), true
		case field.UnixNanosecond:
			return currentTime.UnixNano(), true
		}
	default:
	}
	return nil, false
}
