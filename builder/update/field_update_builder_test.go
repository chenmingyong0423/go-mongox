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

	"github.com/chenmingyong0423/go-mongox/bsonx"

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
			value: bsonx.D(bsonx.E("name", "cmy")),
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
			value: bsonx.D(bsonx.E("name", "cmy")),
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

func Test_fieldUpdateBuilder_Inc(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "bson",
			value: bsonx.NewD().Add("comments", 1).Add("score", 1).Build(),
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
			value: bsonx.NewD().Add("stock", 100).Add("dateExpired", time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)).Add("score", -50).Build(),
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
			value: bsonx.NewD().Add("stock", 100).Add("dateExpired", time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)).Add("score", -50).Build(),
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

func Test_fieldUpdateBuilder_Mul(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "bson",
			value: bsonx.NewD().Add("price", 1.25).Add("qty", 2).Add("score", -1).Add("n", -1.1).Build(),
			want:  bson.D{bson.E{Key: "$mul", Value: bson.D{bson.E{Key: "price", Value: 1.25}, bson.E{Key: "qty", Value: 2}, bson.E{Key: "score", Value: -1}, bson.E{Key: "n", Value: -1.1}}}},
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

func Test_fieldUpdateBuilder_Rename(t *testing.T) {
	testCases := []struct {
		name  string
		value any

		want bson.D
	}{
		{
			name:  "bson",
			value: bsonx.NewD().Add("nmae", "name").Add("name.first", "name.last").Build(),
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

func Test_fieldUpdateBuilder_CurrentDate(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "bson",
			value: bsonx.NewD().Add("lastModified", true).Add("cancellation.date", bsonx.M("$type", "timestamp")).Build(),
			want:  bson.D{{Key: "$currentDate", Value: bson.D{{Key: "lastModified", Value: true}, {Key: "cancellation.date", Value: bson.M{"$type": "timestamp"}}}}},
		},
		{
			name:  "map",
			value: map[string]any{"lastModified": true, "cancellation.date": bsonx.M("$type", "timestamp")},
			want:  bson.D{{Key: "$currentDate", Value: map[string]any{"lastModified": true, "cancellation.date": bson.M{"$type": "timestamp"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().CurrentDate(tc.value).Build())
		})
	}
}
