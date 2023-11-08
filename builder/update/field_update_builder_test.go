// Copyright 2023 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package update

import (
	"testing"
	"time"

	"github.com/chenmingyong0423/go-mongox/converter"

	"github.com/chenmingyong0423/go-mongox/pkg/utils"

	"github.com/chenmingyong0423/go-mongox/types"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_fieldUpdateBuilder_Set(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "bson",
			value: bson.D{{Key: "name", Value: "cmy"}},
			want:  bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: "cmy"}}}},
		},
		{
			name:  "map",
			value: map[string]any{"name": "cmy"},
			want:  bson.D{{Key: "$set", Value: map[string]any{"name": "cmy"}}},
		},
		{
			name:  "struct",
			value: types.UserName{Name: "cmy"},
			want:  bson.D{{Key: "$set", Value: types.UserName{Name: "cmy"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().Set(tc.value).Build())
		})
	}
}

func Test_fieldUpdateBuilder_Unset(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "$unset", Value: bson.D{}}}, BsonBuilder().Unset().Build())
	assert.Equal(t, bson.D{{Key: "$unset", Value: bson.D{bson.E{Key: "name", Value: ""}}}}, BsonBuilder().Unset("name").Build())
	assert.Equal(t, bson.D{{Key: "$unset", Value: bson.D{bson.E{Key: "name", Value: ""}, bson.E{Key: "age", Value: ""}}}}, BsonBuilder().Unset("name", "age").Build())
}

func Test_fieldUpdateBuilder_SetOnInsert(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "bson",
			value: bson.D{{Key: "name", Value: "cmy"}},
			want:  bson.D{{Key: "$setOnInsert", Value: bson.D{{Key: "name", Value: "cmy"}}}},
		},
		{
			name:  "map",
			value: map[string]any{"name": "cmy"},
			want:  bson.D{{Key: "$setOnInsert", Value: map[string]any{"name": "cmy"}}},
		},
		{
			name:  "struct",
			value: types.UserName{Name: "cmy"},
			want:  bson.D{{Key: "$setOnInsert", Value: types.UserName{Name: "cmy"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().SetOnInsert(tc.value).Build())
		})
	}
}

func Test_fieldUpdateBuilder_CurrentDateKeyValues(t *testing.T) {
	testCases := []struct {
		name string
		data []types.KeyValue[any]

		want bson.D
	}{
		{
			name: "nil params",
			want: bson.D{bson.E{Key: "$currentDate", Value: bson.D{}}},
		},
		{
			name: "empty params",
			data: []types.KeyValue[any]{},
			want: bson.D{bson.E{Key: "$currentDate", Value: bson.D{}}},
		},
		{
			name: "normal params",
			data: []types.KeyValue[any]{converter.KeyValue[any]("lastModified", true), converter.KeyValue[any]("cancellation.date", "timestamp")},
			want: bson.D{bson.E{Key: "$currentDate", Value: bson.D{bson.E{Key: "lastModified", Value: true}, bson.E{Key: "cancellation.date", Value: bson.M{"$type": "timestamp"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().CurrentDateKeyValues(tc.data...).Build())
		})
	}
}

func Test_fieldUpdateBuilder_CurrentDateForMap(t *testing.T) {
	testCases := []struct {
		name string
		data map[string]any

		want bson.D
	}{
		{
			name: "nil values",
			want: bson.D{bson.E{Key: "$currentDate", Value: bson.D{}}},
		},
		{
			name: "empty values",
			data: map[string]any{},
			want: bson.D{bson.E{Key: "$currentDate", Value: bson.D{}}},
		},
		{
			name: "normal values",
			data: map[string]any{"lastModified": true, "cancellation.date": "timestamp"},
			want: bson.D{bson.E{Key: "$currentDate", Value: bson.D{bson.E{Key: "lastModified", Value: true}, bson.E{Key: "cancellation.date", Value: bson.M{"$type": "timestamp"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, utils.EqualBSONDElements(tc.want, BsonBuilder().CurrentDateForMap(tc.data).Build()))
		})
	}
}

func Test_fieldUpdateBuilder_Inc(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "bson",
			value: bson.D{{Key: "comments", Value: 1}, {Key: "score", Value: 1}},
			want:  bson.D{{Key: "$inc", Value: bson.D{{Key: "comments", Value: 1}, {Key: "score", Value: 1}}}},
		},
		{
			name:  "map",
			value: map[string]any{"comments": 1, "score": 1},
			want:  bson.D{{Key: "$inc", Value: map[string]any{"comments": 1, "score": 1}}},
		},
		{
			name: "struct",
			value: struct {
				Comments int `bson:"comments"`
				Score    int `bson:"score"`
			}{Comments: 1, Score: 1},
			want: bson.D{{Key: "$inc", Value: struct {
				Comments int `bson:"comments"`
				Score    int `bson:"score"`
			}{Comments: 1, Score: 1}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().Inc(tc.value).Build())
		})
	}
}

func Test_fieldUpdateBuilder_Min(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "bson",
			value: bson.D{{Key: "stock", Value: 100}, {Key: "dateExpired", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, {Key: "score", Value: -50}},
			want:  bson.D{{Key: "$min", Value: bson.D{{Key: "stock", Value: 100}, {Key: "dateExpired", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, {Key: "score", Value: -50}}}},
		},
		{
			name:  "map",
			value: map[string]any{"stock": 100, "dateExpired": time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC), "score": -50},
			want:  bson.D{{Key: "$min", Value: map[string]any{"stock": 100, "dateExpired": time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC), "score": -50}}},
		},
		{
			name: "struct",
			value: struct {
				Stock       int       `bson:"stock"`
				DateExpired time.Time `bson:"dateExpired"`
				Score       int       `bson:"score"`
			}{Stock: 100, DateExpired: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC), Score: -50},
			want: bson.D{{Key: "$min", Value: struct {
				Stock       int       `bson:"stock"`
				DateExpired time.Time `bson:"dateExpired"`
				Score       int       `bson:"score"`
			}{Stock: 100, DateExpired: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC), Score: -50}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().Min(tc.value).Build())
		})
	}
}

func Test_fieldUpdateBuilder_Max(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "bson",
			value: bson.D{{Key: "stock", Value: 100}, {Key: "dateExpired", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, {Key: "score", Value: -50}},
			want:  bson.D{{Key: "$max", Value: bson.D{{Key: "stock", Value: 100}, {Key: "dateExpired", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, {Key: "score", Value: -50}}}},
		},
		{
			name:  "map",
			value: map[string]any{"stock": 100, "dateExpired": time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC), "score": -50},
			want:  bson.D{{Key: "$max", Value: map[string]any{"stock": 100, "dateExpired": time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC), "score": -50}}},
		},
		{
			name: "struct",
			value: struct {
				Stock       int       `bson:"stock"`
				DateExpired time.Time `bson:"dateExpired"`
				Score       int       `bson:"score"`
			}{Stock: 100, DateExpired: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC), Score: -50},
			want: bson.D{{Key: "$max", Value: struct {
				Stock       int       `bson:"stock"`
				DateExpired time.Time `bson:"dateExpired"`
				Score       int       `bson:"score"`
			}{Stock: 100, DateExpired: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC), Score: -50}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().Max(tc.value).Build())
		})
	}
}

func Test_fieldUpdateBuilder_MaxKeyValues(t *testing.T) {
	testCases := []struct {
		name         string
		bsonElements []types.KeyValue[any]
		want         bson.D
	}{
		{
			name: "nil params",
			want: bson.D{bson.E{Key: "$max", Value: bson.D{}}},
		},
		{
			name:         "empty params",
			bsonElements: []types.KeyValue[any]{},
			want:         bson.D{bson.E{Key: "$max", Value: bson.D{}}},
		},
		{
			name:         "normal params",
			bsonElements: []types.KeyValue[any]{converter.KeyValue[any]("stock", 100), converter.KeyValue[any]("dateExpired", time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)), converter.KeyValue[any]("score", -50)},
			want:         bson.D{bson.E{Key: "$max", Value: bson.D{bson.E{Key: "stock", Value: 100}, bson.E{Key: "dateExpired", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "score", Value: -50}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().MaxKeyValues(tc.bsonElements...).Build())
		})
	}
}

func Test_fieldUpdateBuilder_Mul(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name: "bson",
			value: bson.D{{Key: "price", Value: 1.25}, {Key: "qty", Value: 2}, {Key: "score", Value: -1},
				{Key: "n", Value: -1.1}},
			want: bson.D{bson.E{Key: "$mul", Value: bson.D{bson.E{Key: "price", Value: 1.25}, bson.E{Key: "qty", Value: 2}, bson.E{Key: "score", Value: -1}, bson.E{Key: "n", Value: -1.1}}}},
		},
		{
			name:  "map",
			value: map[string]any{"price": 1.25, "qty": 2, "score": -1, "n": -1.1},
			want:  bson.D{bson.E{Key: "$mul", Value: map[string]any{"price": 1.25, "qty": 2, "score": -1, "n": -1.1}}},
		},
		{
			name: "struct",
			value: struct {
				Price float64 `bson:"price"`
				Qty   int     `bson:"qty"`
				Score int     `bson:"score"`
				N     float64 `bson:"n"`
			}{Price: 1.25, Qty: 2, Score: -1, N: -1.1},
			want: bson.D{bson.E{Key: "$mul", Value: struct {
				Price float64 `bson:"price"`
				Qty   int     `bson:"qty"`
				Score int     `bson:"score"`
				N     float64 `bson:"n"`
			}{Price: 1.25, Qty: 2, Score: -1, N: -1.1}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().Mul(tc.value).Build())
		})
	}
}
func Test_fieldUpdateBuilder_MulKeyValues(t *testing.T) {
	testCases := []struct {
		name         string
		bsonElements []types.KeyValue[any]
		want         bson.D
	}{
		{
			name: "nil params",
			want: bson.D{bson.E{Key: "$mul", Value: bson.D{}}},
		},
		{
			name:         "empty params",
			bsonElements: []types.KeyValue[any]{},
			want:         bson.D{bson.E{Key: "$mul", Value: bson.D{}}},
		},
		{
			name:         "values contain non-number",
			bsonElements: []types.KeyValue[any]{converter.KeyValue[any]("price", 1.25), converter.KeyValue[any]("qty", "2.5")},
			want:         bson.D{bson.E{Key: "$mul", Value: bson.D{bson.E{Key: "price", Value: 1.25}}}},
		},
		{
			name:         "normal params",
			bsonElements: []types.KeyValue[any]{converter.KeyValue[any]("price", 1.25), converter.KeyValue[any]("qty", 2), converter.KeyValue[any]("score", -1), converter.KeyValue[any]("n", -1.1)},
			want:         bson.D{bson.E{Key: "$mul", Value: bson.D{bson.E{Key: "price", Value: 1.25}, bson.E{Key: "qty", Value: 2}, bson.E{Key: "score", Value: -1}, bson.E{Key: "n", Value: -1.1}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().MulKeyValues(tc.bsonElements...).Build())
		})
	}
}

func Test_fieldUpdateBuilder_Rename(t *testing.T) {
	testCases := []struct {
		name  string
		value any

		want bson.D
	}{
		{
			name:  "bson",
			value: bson.D{{Key: "nmae", Value: "name"}, {Key: "name.first", Value: "name.last"}},
			want:  bson.D{bson.E{Key: "$rename", Value: bson.D{bson.E{Key: "nmae", Value: "name"}, bson.E{Key: "name.first", Value: "name.last"}}}},
		},
		{
			name:  "map",
			value: map[string]any{"nmae": "name", "name.first": "name.last"},
			want:  bson.D{bson.E{Key: "$rename", Value: map[string]any{"nmae": "name", "name.first": "name.last"}}},
		},
		{
			name: "struct",
			value: struct {
				Nmae      string `bson:"nmae"`
				NameFirst string `bson:"name.first"`
			}{Nmae: "name", NameFirst: "name.last"},
			want: bson.D{bson.E{Key: "$rename", Value: struct {
				Nmae      string `bson:"nmae"`
				NameFirst string `bson:"name.first"`
			}{Nmae: "name", NameFirst: "name.last"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().Rename(tc.value).Build())
		})
	}
}

func Test_fieldUpdateBuilder_RenameKeyValues(t *testing.T) {
	testCases := []struct {
		name         string
		bsonElements []types.KeyValue[string]

		want bson.D
	}{
		{
			name: "nil params",
			want: bson.D{bson.E{Key: "$rename", Value: bson.D{}}},
		},
		{
			name:         "empty params",
			bsonElements: []types.KeyValue[string]{},
			want:         bson.D{bson.E{Key: "$rename", Value: bson.D{}}},
		},
		{
			name:         "normal params",
			bsonElements: []types.KeyValue[string]{converter.KeyValue[string]("nmae", "name"), converter.KeyValue[string]("name.first", "name.last")},
			want:         bson.D{bson.E{Key: "$rename", Value: bson.D{bson.E{Key: "nmae", Value: "name"}, bson.E{Key: "name.first", Value: "name.last"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().RenameKeyValues(tc.bsonElements...).Build())
		})
	}
}

func Test_fieldUpdateBuilder_SetKeyValues(t *testing.T) {
	testCases := []struct {
		name         string
		bsonElements []types.KeyValue[any]
		want         bson.D
	}{
		{
			name:         "zero params",
			bsonElements: nil,
			want:         bson.D{bson.E{Key: "$set", Value: bson.D{}}},
		},
		{
			name:         "normal params",
			bsonElements: []types.KeyValue[any]{converter.KeyValue[any]("name", "cmy"), converter.KeyValue[any]("age", 24)},
			want:         bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: "cmy"}, bson.E{Key: "age", Value: 24}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().SetKeyValues(tc.bsonElements...).Build())
		})
	}
}

func Test_fieldUpdateBuilder_SetOnInsertKeyValues(t *testing.T) {
	testCases := []struct {
		name         string
		bsonElements []types.KeyValue[any]
		want         bson.D
	}{
		{
			name: "nil bsonElements",
			want: bson.D{bson.E{Key: "$setOnInsert", Value: bson.D{}}},
		},
		{
			name:         "empty bsonElements",
			bsonElements: []types.KeyValue[any]{},
			want:         bson.D{bson.E{Key: "$setOnInsert", Value: bson.D{}}},
		},
		{
			name:         "normal bsonElements",
			bsonElements: []types.KeyValue[any]{converter.KeyValue[any]("name", "cmy"), converter.KeyValue[any]("age", 24)},
			want:         bson.D{bson.E{Key: "$setOnInsert", Value: bson.D{bson.E{Key: "name", Value: "cmy"}, bson.E{Key: "age", Value: 24}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().SetOnInsertKeyValues(tc.bsonElements...).Build())
		})
	}
}

func Test_fieldUpdateBuilder_CurrentDate(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "bson",
			value: bson.D{{Key: "lastModified", Value: true}, {Key: "cancellation.date", Value: bson.M{"$type": "timestamp"}}},
			want:  bson.D{{Key: "$currentDate", Value: bson.D{{Key: "lastModified", Value: true}, {Key: "cancellation.date", Value: bson.M{"$type": "timestamp"}}}}},
		},
		{
			name:  "map",
			value: map[string]any{"lastModified": true, "cancellation.date": bson.M{"$type": "timestamp"}},
			want:  bson.D{{Key: "$currentDate", Value: map[string]any{"lastModified": true, "cancellation.date": bson.M{"$type": "timestamp"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().CurrentDate(tc.value).Build())
		})
	}
}

func Test_fieldUpdateBuilder_IncKeyValues(t *testing.T) {
	testCases := []struct {
		name         string
		bsonElements []types.KeyValue[int]
		want         bson.D
	}{
		{
			name: "nil params",
			want: bson.D{bson.E{Key: "$inc", Value: bson.D{}}},
		},
		{
			name:         "empty params",
			bsonElements: []types.KeyValue[int]{},
			want:         bson.D{bson.E{Key: "$inc", Value: bson.D{}}},
		},
		{
			name:         "normal params",
			bsonElements: []types.KeyValue[int]{converter.KeyValue[int]("read", 1), converter.KeyValue[int]("likes", 1), converter.KeyValue[int]("comments", 1), converter.KeyValue[int]("score", -1)},
			want:         bson.D{bson.E{Key: "$inc", Value: bson.D{bson.E{Key: "read", Value: 1}, bson.E{Key: "likes", Value: 1}, bson.E{Key: "comments", Value: 1}, bson.E{Key: "score", Value: -1}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().IncKeyValues(tc.bsonElements...).Build())
		})
	}
}

func Test_fieldUpdateBuilder_MinKeyValues(t *testing.T) {
	testCases := []struct {
		name         string
		bsonElements []types.KeyValue[any]
		want         bson.D
	}{
		{
			name: "nil params",
			want: bson.D{bson.E{Key: "$min", Value: bson.D{}}},
		},
		{
			name:         "empty params",
			bsonElements: []types.KeyValue[any]{},
			want:         bson.D{bson.E{Key: "$min", Value: bson.D{}}},
		},
		{
			name:         "normal params",
			bsonElements: []types.KeyValue[any]{converter.KeyValue[any]("stock", 100), converter.KeyValue[any]("dateExpired", time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)), converter.KeyValue[any]("score", -50)},
			want:         bson.D{bson.E{Key: "$min", Value: bson.D{bson.E{Key: "stock", Value: 100}, bson.E{Key: "dateExpired", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "score", Value: -50}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().MinKeyValues(tc.bsonElements...).Build())
		})
	}
}
