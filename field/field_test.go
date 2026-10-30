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
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/stretchr/testify/require"
)

func TestParseFields(t *testing.T) {
	type model struct {
		ID        bson.ObjectID `bson:"_id,omitempty" mongox:"autoID"`
		CreatedAt time.Time     `bson:"created_at"`
		UpdatedAt time.Time     `bson:"updated_at"`
		DeletedAt time.Time     `bson:"deleted_at,omitempty"`
	}

	testCases := []struct {
		name string
		doc  any
		want []*Filed
	}{
		{
			name: "nil",
			doc:  nil,
			want: nil,
		},
		{
			name: "not struct",
			doc:  1,
			want: nil,
		},
		{
			name: "empty struct",
			doc:  struct{}{},
			want: []*Filed{},
		},
		{
			name: "none inlined struct",
			doc: &struct {
				ID               bson.ObjectID `bson:"_id,omitempty" mongox:"autoID"`
				Name             string        `bson:"name"`
				CreatedAt        time.Time     `bson:"created_at"`
				UpdatedAt        time.Time     `bson:"updated_at"`
				DeletedAt        time.Time     `bson:"deleted_at,omitempty"`
				CreateSecondTime int64         `bson:"create_second_time" mongox:"autoCreateTime:second"`
				UpdateSecondTime int64         `bson:"update_second_time" mongox:"autoUpdateTime:second"`
				CreateMilliTime  int64         `bson:"create_milli_time" mongox:"autoCreateTime:milli"`
				UpdateMilliTime  int64         `bson:"update_milli_time" mongox:"autoUpdateTime:milli"`
				CreateNanoTime   int64         `bson:"create_nano_time" mongox:"autoCreateTime:nano"`
				UpdateNanoTime   int64         `bson:"update_nano_time" mongox:"autoUpdateTime:nano"`

				NoneBsonTagField    string
				InvalidBsonTagField string    `bson:",omitempty"`
				InvalidTimeTagField time.Time `bson:"invalid_time_tag_field" mongox:"autoCreateTime:time"`
			}{},
			want: []*Filed{
				{
					Name:       "ID",
					MongoField: "_id",
					AutoID:     true,
					FieldType:  reflect.TypeFor[bson.ObjectID](),
				},
				{
					Name:       "Name",
					MongoField: "name",
					FieldType:  reflect.TypeFor[string](),
				},
				{
					Name:           "CreatedAt",
					MongoField:     "created_at",
					FieldType:      reflect.TypeFor[time.Time](),
					AutoCreateTime: UnixTime,
				},
				{
					Name:           "UpdatedAt",
					MongoField:     "updated_at",
					FieldType:      reflect.TypeFor[time.Time](),
					AutoUpdateTime: UnixTime,
				},
				{
					Name:           "DeletedAt",
					MongoField:     "deleted_at",
					FieldType:      reflect.TypeFor[time.Time](),
					AutoDeleteTime: UnixTime,
				},
				{
					Name:           "CreateSecondTime",
					MongoField:     "create_second_time",
					FieldType:      reflect.TypeFor[int64](),
					AutoCreateTime: UnixSecond,
				},
				{
					Name:           "UpdateSecondTime",
					MongoField:     "update_second_time",
					FieldType:      reflect.TypeFor[int64](),
					AutoUpdateTime: UnixSecond,
				},
				{
					Name:           "CreateMilliTime",
					MongoField:     "create_milli_time",
					FieldType:      reflect.TypeFor[int64](),
					AutoCreateTime: UnixMillisecond,
				},
				{
					Name:           "UpdateMilliTime",
					MongoField:     "update_milli_time",
					FieldType:      reflect.TypeFor[int64](),
					AutoUpdateTime: UnixMillisecond,
				},
				{
					Name:           "CreateNanoTime",
					MongoField:     "create_nano_time",
					FieldType:      reflect.TypeFor[int64](),
					AutoCreateTime: UnixNanosecond,
				},
				{
					Name:           "UpdateNanoTime",
					MongoField:     "update_nano_time",
					FieldType:      reflect.TypeFor[int64](),
					AutoUpdateTime: UnixNanosecond,
				},
				{
					Name:       "NoneBsonTagField",
					MongoField: "nonebsontagfield",
					FieldType:  reflect.TypeFor[string](),
				},
				{
					Name:       "InvalidBsonTagField",
					MongoField: "invalidbsontagfield",
					FieldType:  reflect.TypeFor[string](),
				},
				{
					Name:           "InvalidTimeTagField",
					MongoField:     "invalid_time_tag_field",
					FieldType:      reflect.TypeFor[time.Time](),
					AutoCreateTime: 0,
				},
			},
		},
		{
			name: "inlined struct",
			doc: struct {
				model            `bson:",inline"`
				Name             string `bson:"name"`
				CreateSecondTime int64  `bson:"create_second_time" mongox:"autoCreateTime"`
				UpdateSecondTime int64  `bson:"update_second_time" mongox:"autoUpdateTime"`
				CreateMilliTime  int64  `bson:"create_milli_time" mongox:"autoCreateTime:milli"`
				UpdateMilliTime  int64  `bson:"update_milli_time" mongox:"autoUpdateTime:milli"`
				CreateNanoTime   int64  `bson:"create_nano_time" mongox:"autoCreateTime:nano"`
				UpdateNanoTime   int64  `bson:"update_nano_time" mongox:"autoUpdateTime:nano"`
			}{},
			want: []*Filed{
				{
					Name:      "model",
					FieldType: reflect.TypeFor[model](),
					InlinedFields: []*Filed{
						{
							Name:       "ID",
							MongoField: "_id",
							AutoID:     true,
							FieldType:  reflect.TypeFor[bson.ObjectID](),
						},
						{
							Name:           "CreatedAt",
							MongoField:     "created_at",
							FieldType:      reflect.TypeFor[time.Time](),
							AutoCreateTime: UnixTime,
						},
						{
							Name:           "UpdatedAt",
							MongoField:     "updated_at",
							FieldType:      reflect.TypeFor[time.Time](),
							AutoUpdateTime: UnixTime,
						},
						{
							Name:           "DeletedAt",
							MongoField:     "deleted_at",
							FieldType:      reflect.TypeFor[time.Time](),
							AutoDeleteTime: UnixTime,
						},
					},
				},
				{
					Name:       "Name",
					MongoField: "name",
					FieldType:  reflect.TypeFor[string](),
				},
				{
					Name:           "CreateSecondTime",
					MongoField:     "create_second_time",
					FieldType:      reflect.TypeFor[int64](),
					AutoCreateTime: UnixSecond,
				},
				{
					Name:           "UpdateSecondTime",
					MongoField:     "update_second_time",
					FieldType:      reflect.TypeFor[int64](),
					AutoUpdateTime: UnixSecond,
				},
				{
					Name:           "CreateMilliTime",
					MongoField:     "create_milli_time",
					FieldType:      reflect.TypeFor[int64](),
					AutoCreateTime: UnixMillisecond,
				},
				{
					Name:           "UpdateMilliTime",
					MongoField:     "update_milli_time",
					FieldType:      reflect.TypeFor[int64](),
					AutoUpdateTime: UnixMillisecond,
				},
				{
					Name:           "CreateNanoTime",
					MongoField:     "create_nano_time",
					FieldType:      reflect.TypeFor[int64](),
					AutoCreateTime: UnixNanosecond,
				},
				{
					Name:           "UpdateNanoTime",
					MongoField:     "update_nano_time",
					FieldType:      reflect.TypeFor[int64](),
					AutoUpdateTime: UnixNanosecond,
				},
			},
		},
		{
			name: "Not time.Time type for default time field",
			doc: struct {
				CreatedAt int64 `bson:"created_at"`
				UpdatedAt int   `bson:"updated_at"`
			}{},
			want: []*Filed{
				{
					Name:           "CreatedAt",
					MongoField:     "created_at",
					FieldType:      reflect.TypeFor[int64](),
					AutoCreateTime: UnixSecond,
				},
				{
					Name:           "UpdatedAt",
					MongoField:     "updated_at",
					FieldType:      reflect.TypeFor[int](),
					AutoUpdateTime: UnixSecond,
				},
			},
		},
		{
			name: "mongox and bson tag for default time field with a value in unix seconds.",
			doc: struct {
				CreatedAt int64 `mongox:"autoCreateTime" bson:"created_at"`
				UpdatedAt int   `bson:"updated_at"`
			}{},
			want: []*Filed{
				{
					Name:           "CreatedAt",
					MongoField:     "created_at",
					FieldType:      reflect.TypeFor[int64](),
					AutoCreateTime: UnixSecond,
				},
				{
					Name:           "UpdatedAt",
					MongoField:     "updated_at",
					FieldType:      reflect.TypeFor[int](),
					AutoUpdateTime: UnixSecond,
				},
			},
		},
		{
			name: "mongox tag for default time.Time fields",
			doc: struct {
				CreatedAt time.Time `bson:"created_at" mongox:"autoCreateTime"`
				UpdatedAt time.Time `bson:"updated_at" mongox:"autoUpdateTime"`
			}{},
			want: []*Filed{
				{
					Name:           "CreatedAt",
					MongoField:     "created_at",
					FieldType:      reflect.TypeFor[time.Time](),
					AutoCreateTime: UnixTime,
				},
				{
					Name:           "UpdatedAt",
					MongoField:     "updated_at",
					FieldType:      reflect.TypeFor[time.Time](),
					AutoUpdateTime: UnixTime,
				},
			},
		},
		{
			name: "mongox timestamp precision tags on time.Time fields",
			doc: struct {
				CreateSecondTime time.Time `bson:"create_second_time" mongox:"autoCreateTime:second"`
				UpdateMilliTime  time.Time `bson:"update_milli_time" mongox:"autoUpdateTime:milli"`
				CreateNanoTime   time.Time `bson:"create_nano_time" mongox:"autoCreateTime:nano"`
			}{},
			want: []*Filed{
				{
					Name:           "CreateSecondTime",
					MongoField:     "create_second_time",
					FieldType:      reflect.TypeFor[time.Time](),
					AutoCreateTime: 0,
				},
				{
					Name:           "UpdateMilliTime",
					MongoField:     "update_milli_time",
					FieldType:      reflect.TypeFor[time.Time](),
					AutoUpdateTime: 0,
				},
				{
					Name:           "CreateNanoTime",
					MongoField:     "create_nano_time",
					FieldType:      reflect.TypeFor[time.Time](),
					AutoCreateTime: 0,
				},
			},
		},
		{
			name: "mongox and bson tag for default time field with a value in unix milli seconds.",
			doc: struct {
				CreatedAt int64 `bson:"created_at" mongox:"autoCreateTime:milli"`
				UpdatedAt int   `bson:"updated_at"`
			}{},
			want: []*Filed{
				{
					Name:           "CreatedAt",
					MongoField:     "created_at",
					FieldType:      reflect.TypeFor[int64](),
					AutoCreateTime: UnixMillisecond,
				},
				{
					Name:           "UpdatedAt",
					MongoField:     "updated_at",
					FieldType:      reflect.TypeFor[int](),
					AutoUpdateTime: UnixSecond,
				},
			},
		},
		{
			name: "mongox timestamp precision tags on int fields",
			doc: struct {
				CreateMilliTime int `bson:"create_milli_time" mongox:"autoCreateTime:milli"`
				UpdateNanoTime  int `bson:"update_nano_time" mongox:"autoUpdateTime:nano"`
			}{},
			want: []*Filed{
				{
					Name:           "CreateMilliTime",
					MongoField:     "create_milli_time",
					FieldType:      reflect.TypeFor[int](),
					AutoCreateTime: UnixMillisecond,
				},
				{
					Name:           "UpdateNanoTime",
					MongoField:     "update_nano_time",
					FieldType:      reflect.TypeFor[int](),
					AutoUpdateTime: UnixNanosecond,
				},
			},
		},
		{
			name: "invalid type 4 default time field",
			doc: struct {
				CreatedAt string `bson:"created_at"`
				UpdatedAt bool   `bson:"updated_at"`
			}{},
			want: []*Filed{
				{
					Name:           "CreatedAt",
					MongoField:     "created_at",
					FieldType:      reflect.TypeFor[string](),
					AutoCreateTime: 0,
				},
				{
					Name:           "UpdatedAt",
					MongoField:     "updated_at",
					FieldType:      reflect.TypeFor[bool](),
					AutoUpdateTime: 0,
				},
			},
		},
		{
			name: "invalid type for mongox time tag",
			doc: struct {
				CreatedAt string `bson:"created_at" mongox:"autoCreateTime"`
				UpdatedAt bool   `bson:"updated_at" mongox:"autoUpdateTime:milli"`
			}{},
			want: []*Filed{
				{
					Name:           "CreatedAt",
					MongoField:     "created_at",
					FieldType:      reflect.TypeFor[string](),
					AutoCreateTime: 0,
				},
				{
					Name:           "UpdatedAt",
					MongoField:     "updated_at",
					FieldType:      reflect.TypeFor[bool](),
					AutoUpdateTime: 0,
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := ParseFields(tc.doc)
			require.ElementsMatch(t, tc.want, got)
		})
	}
}

func TestParseInlineDeletedAtOptions(t *testing.T) {
	type Audit struct {
		DeletedAt int64 `bson:"removed_at" mongox:"autoDeleteTime:milli"`
	}
	tests := []struct {
		name    string
		doc     any
		inlined bool
	}{
		{"inline", struct {
			Audit `bson:",inline"`
		}{}, true},
		{"inline before omitempty", struct {
			Audit `bson:",inline,omitempty"`
		}{}, true},
		{"inline after omitempty", struct {
			Audit `bson:",omitempty,inline"`
		}{}, true},
		{"inline between options", struct {
			Audit `bson:",omitempty,inline,minsize"`
		}{}, true},
		{"no tag", struct{ Audit }{}, false},
		{"no inline option", struct {
			Audit `bson:",omitempty"`
		}{}, false},
		{"inline in field name", struct {
			Audit `bson:"inline,omitempty"`
		}{}, false},
		{"inline option prefix", struct {
			Audit `bson:",inline_extra"`
		}{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := ParseFields(tt.doc)
			require.Len(t, fields, 1)
			if !tt.inlined {
				require.Nil(t, fields[0].InlinedFields)
				return
			}
			require.Len(t, fields[0].InlinedFields, 1)
			deleted := fields[0].InlinedFields[0]
			require.Equal(t, "removed_at", deleted.MongoField)
			require.Equal(t, UnixMillisecond, deleted.AutoDeleteTime)
			encoded, err := bson.Marshal(tt.doc)
			require.NoError(t, err)
			require.Equal(t, bson.TypeInt64, bson.Raw(encoded).Lookup(deleted.MongoField).Type)
		})
	}
}
