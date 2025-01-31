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
					FieldType:  reflect.TypeOf(bson.ObjectID{}),
				},
				{
					Name:       "Name",
					MongoField: "name",
					FieldType:  reflect.TypeOf(""),
				},
				{
					Name:           "CreatedAt",
					MongoField:     "created_at",
					FieldType:      reflect.TypeOf(time.Time{}),
					AutoCreateTime: UnixTime,
				},
				{
					Name:           "UpdatedAt",
					MongoField:     "updated_at",
					FieldType:      reflect.TypeOf(time.Time{}),
					AutoUpdateTime: UnixTime,
				},
				{
					Name:       "DeletedAt",
					MongoField: "deleted_at",
					FieldType:  reflect.TypeOf(time.Time{}),
				},
				{
					Name:           "CreateSecondTime",
					MongoField:     "create_second_time",
					FieldType:      reflect.TypeOf(int64(0)),
					AutoCreateTime: UnixSecond,
				},
				{
					Name:           "UpdateSecondTime",
					MongoField:     "update_second_time",
					FieldType:      reflect.TypeOf(int64(0)),
					AutoUpdateTime: UnixSecond,
				},
				{
					Name:           "CreateMilliTime",
					MongoField:     "create_milli_time",
					FieldType:      reflect.TypeOf(int64(0)),
					AutoCreateTime: UnixMillisecond,
				},
				{
					Name:           "UpdateMilliTime",
					MongoField:     "update_milli_time",
					FieldType:      reflect.TypeOf(int64(0)),
					AutoUpdateTime: UnixMillisecond,
				},
				{
					Name:           "CreateNanoTime",
					MongoField:     "create_nano_time",
					FieldType:      reflect.TypeOf(int64(0)),
					AutoCreateTime: UnixNanosecond,
				},
				{
					Name:           "UpdateNanoTime",
					MongoField:     "update_nano_time",
					FieldType:      reflect.TypeOf(int64(0)),
					AutoUpdateTime: UnixNanosecond,
				},
				{
					Name:       "NoneBsonTagField",
					MongoField: "NoneBsonTagField",
					FieldType:  reflect.TypeOf(""),
				},
				{
					Name:       "InvalidBsonTagField",
					MongoField: "InvalidBsonTagField",
					FieldType:  reflect.TypeOf(""),
				},
				{
					Name:           "InvalidTimeTagField",
					MongoField:     "invalid_time_tag_field",
					FieldType:      reflect.TypeOf(time.Time{}),
					AutoCreateTime: 0,
				},
			},
		},
		{
			name: "inlined struct",
			doc: struct {
				model            `bson:",inline"`
				Name             string `bson:"name"`
				CreateSecondTime int64  `bson:"create_second_time" mongox:"autoCreateTime:second"`
				UpdateSecondTime int64  `bson:"update_second_time" mongox:"autoUpdateTime:second"`
				CreateMilliTime  int64  `bson:"create_milli_time" mongox:"autoCreateTime:milli"`
				UpdateMilliTime  int64  `bson:"update_milli_time" mongox:"autoUpdateTime:milli"`
				CreateNanoTime   int64  `bson:"create_nano_time" mongox:"autoCreateTime:nano"`
				UpdateNanoTime   int64  `bson:"update_nano_time" mongox:"autoUpdateTime:nano"`
			}{},
			want: []*Filed{
				{
					Name:      "model",
					FieldType: reflect.TypeOf(model{}),
					InlinedFields: []*Filed{
						{
							Name:       "ID",
							MongoField: "_id",
							AutoID:     true,
							FieldType:  reflect.TypeOf(bson.ObjectID{}),
						},
						{
							Name:           "CreatedAt",
							MongoField:     "created_at",
							FieldType:      reflect.TypeOf(time.Time{}),
							AutoCreateTime: UnixTime,
						},
						{
							Name:           "UpdatedAt",
							MongoField:     "updated_at",
							FieldType:      reflect.TypeOf(time.Time{}),
							AutoUpdateTime: UnixTime,
						},
						{
							Name:       "DeletedAt",
							MongoField: "deleted_at",
							FieldType:  reflect.TypeOf(time.Time{}),
						},
					},
				},
				{
					Name:       "Name",
					MongoField: "name",
					FieldType:  reflect.TypeOf(""),
				},
				{
					Name:           "CreateSecondTime",
					MongoField:     "create_second_time",
					FieldType:      reflect.TypeOf(int64(0)),
					AutoCreateTime: UnixSecond,
				},
				{
					Name:           "UpdateSecondTime",
					MongoField:     "update_second_time",
					FieldType:      reflect.TypeOf(int64(0)),
					AutoUpdateTime: UnixSecond,
				},
				{
					Name:           "CreateMilliTime",
					MongoField:     "create_milli_time",
					FieldType:      reflect.TypeOf(int64(0)),
					AutoCreateTime: UnixMillisecond,
				},
				{
					Name:           "UpdateMilliTime",
					MongoField:     "update_milli_time",
					FieldType:      reflect.TypeOf(int64(0)),
					AutoUpdateTime: UnixMillisecond,
				},
				{
					Name:           "CreateNanoTime",
					MongoField:     "create_nano_time",
					FieldType:      reflect.TypeOf(int64(0)),
					AutoCreateTime: UnixNanosecond,
				},
				{
					Name:           "UpdateNanoTime",
					MongoField:     "update_nano_time",
					FieldType:      reflect.TypeOf(int64(0)),
					AutoUpdateTime: UnixNanosecond,
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
