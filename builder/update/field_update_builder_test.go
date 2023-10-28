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

	"github.com/chenmingyong0423/go-mongox/pkg"
	"github.com/chenmingyong0423/go-mongox/types"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_fieldUpdateBuilder_Set(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: "cmy"}}}}, BsonBuilder().Set("name", "cmy").Build())
}

func Test_fieldUpdateBuilder_SetForMap(t *testing.T) {

	testCases := []struct {
		name string
		data map[string]any

		want bson.D
	}{
		{
			name: "nil values",
			want: bson.D{},
		},
		{
			name: "empty values",
			data: map[string]any{},
			want: bson.D{},
		},
		{
			name: "normal values",
			data: map[string]any{"name": "cmy"},
			want: bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: "cmy"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().SetForMap(tc.data).Build())
		})
	}
}

func Test_fieldUpdateBuilder_SetForStruct(t *testing.T) {
	testCases := []struct {
		name string
		data any

		want bson.D
	}{
		{
			name: "nil values",
			want: bson.D{},
		},
		{
			name: "empty values",
			data: types.UpdatedUser{},
			want: bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: ""}, bson.E{Key: "age", Value: int64(0)}}}},
		},
		{
			name: "normal values",
			data: types.UpdatedUser{Name: "cmy", Age: 24},
			want: bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: "cmy"}, bson.E{Key: "age", Value: int64(24)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().SetForStruct(tc.data).Build())
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
		name      string
		keyValues []any
		want      bson.D
	}{
		{
			name: "zero params",
			want: bson.D{bson.E{Key: "$setOnInsert", Value: bson.D{}}},
		},
		{
			name:      "odd params",
			keyValues: []any{"name", "cmy", "age"},
			want:      bson.D{bson.E{Key: "$setOnInsert", Value: bson.D{}}},
		},
		{
			name:      "keys contain non-string",
			keyValues: []any{24, "age", "name", "cmy"},
			want:      bson.D{bson.E{Key: "$setOnInsert", Value: bson.D{bson.E{Key: "name", Value: "cmy"}}}},
		},
		{
			name:      "normal params",
			keyValues: []any{"name", "cmy", "age", 24},
			want:      bson.D{bson.E{Key: "$setOnInsert", Value: bson.D{bson.E{Key: "name", Value: "cmy"}, bson.E{Key: "age", Value: 24}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().SetOnInsert(tc.keyValues...).Build())
		})
	}
}

func Test_fieldUpdateBuilder_SetOnInsertForMap(t *testing.T) {
	testCases := []struct {
		name string
		data map[string]any

		want bson.D
	}{
		{
			name: "nil values",
			want: bson.D{},
		},
		{
			name: "empty values",
			data: map[string]any{},
			want: bson.D{bson.E{Key: "$setOnInsert", Value: bson.D{}}},
		},
		{
			name: "normal values",
			data: map[string]any{"name": "cmy", "age": 24},
			want: bson.D{bson.E{Key: "$setOnInsert", Value: bson.D{bson.E{Key: "name", Value: "cmy"}, bson.E{Key: "age", Value: 24}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, pkg.EqualBSONDElements(tc.want, BsonBuilder().SetOnInsertForMap(tc.data).Build()))
		})
	}
}

func Test_fieldUpdateBuilder_CurrentDate(t *testing.T) {
	testCases := []struct {
		name string
		data []any

		want bson.D
	}{
		{
			name: "zero params",
			want: bson.D{bson.E{Key: "$currentDate", Value: bson.D{}}},
		},
		{
			name: "odd params",
			data: []any{"lastModified", true, "cancellation.date"},
			want: bson.D{bson.E{Key: "$currentDate", Value: bson.D{}}},
		},
		{
			name: "keys contain non-string",
			data: []any{true, "lastModified", "cancellation.date", "timestamp"},
			want: bson.D{bson.E{Key: "$currentDate", Value: bson.D{bson.E{Key: "cancellation.date", Value: bson.M{"$type": "timestamp"}}}}},
		},
		{
			name: "normal params",
			data: []any{"lastModified", true, "cancellation.date", "timestamp"},
			want: bson.D{bson.E{Key: "$currentDate", Value: bson.D{bson.E{Key: "lastModified", Value: true}, bson.E{Key: "cancellation.date", Value: bson.M{"$type": "timestamp"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().CurrentDate(tc.data...).Build())
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
			want: bson.D{},
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
			assert.True(t, pkg.EqualBSONDElements(tc.want, BsonBuilder().CurrentDateForMap(tc.data).Build()))
		})
	}
}

func Test_fieldUpdateBuilder_Inc(t *testing.T) {
	testCases := []struct {
		name string
		data []any

		want bson.D
	}{
		{
			name: "zero params",
			want: bson.D{bson.E{Key: "$inc", Value: bson.D{}}},
		},
		{
			name: "odd params",
			data: []any{"read", 1, "likes"},
			want: bson.D{bson.E{Key: "$inc", Value: bson.D{}}},
		},
		{
			name: "keys contain non-string",
			data: []any{1, "read", "likes", 1},
			want: bson.D{bson.E{Key: "$inc", Value: bson.D{bson.E{Key: "likes", Value: 1}}}},
		},
		{
			name: "values contain non-int",
			data: []any{"read", 1, "likes", 1.1, "comments", "1"},
			want: bson.D{bson.E{Key: "$inc", Value: bson.D{bson.E{Key: "read", Value: 1}}}},
		},
		{
			name: "normal params",
			data: []any{"read", 1, "likes", 1, "comments", 1, "score", -1},
			want: bson.D{bson.E{Key: "$inc", Value: bson.D{bson.E{Key: "read", Value: 1}, bson.E{Key: "likes", Value: 1}, bson.E{Key: "comments", Value: 1}, bson.E{Key: "score", Value: -1}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().Inc(tc.data...).Build())
		})
	}
}

func Test_fieldUpdateBuilder_IncForMap(t *testing.T) {
	testCases := []struct {
		name string
		data map[string]int

		want bson.D
	}{
		{
			name: "nil values",
			want: bson.D{},
		},
		{
			name: "empty values",
			data: map[string]int{},
			want: bson.D{bson.E{Key: "$inc", Value: bson.D{}}},
		},
		{
			name: "normal values",
			data: map[string]int{"read": 1, "likes": 1, "comments": 1, "score": -1},
			want: bson.D{bson.E{Key: "$inc", Value: bson.D{bson.E{Key: "read", Value: 1}, bson.E{Key: "likes", Value: 1}, bson.E{Key: "comments", Value: 1}, bson.E{Key: "score", Value: -1}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, pkg.EqualBSONDElements(tc.want, BsonBuilder().IncForMap(tc.data).Build()))
		})
	}
}

func Test_fieldUpdateBuilder_Min(t *testing.T) {
	testCases := []struct {
		name string
		data []any

		want bson.D
	}{
		{
			name: "zero params",
			want: bson.D{bson.E{Key: "$min", Value: bson.D{}}},
		},
		{
			name: "odd params",
			data: []any{"stock", 50, "ordered"},
			want: bson.D{bson.E{Key: "$min", Value: bson.D{}}},
		},
		{
			name: "keys contain non-string",
			data: []any{50, "stock", "ordered", 100},
			want: bson.D{bson.E{Key: "$min", Value: bson.D{bson.E{Key: "ordered", Value: 100}}}},
		},
		{
			name: "normal params",
			data: []any{"stock", 50, "score", -50, "dateExpired", time.Date(2023, 10, 29, 0, 0, 0, 0, time.UTC)},
			want: bson.D{bson.E{Key: "$min", Value: bson.D{bson.E{Key: "stock", Value: 50}, bson.E{Key: "score", Value: -50}, bson.E{Key: "dateExpired", Value: time.Date(2023, 10, 29, 0, 0, 0, 0, time.UTC)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().Min(tc.data...).Build())
		})
	}
}

func Test_fieldUpdateBuilder_MinForMap(t *testing.T) {
	testCases := []struct {
		name string
		data map[string]any

		want bson.D
	}{
		{
			name: "nil values",
			want: bson.D{},
		},
		{
			name: "empty values",
			data: map[string]any{},
			want: bson.D{bson.E{Key: "$min", Value: bson.D{}}},
		},
		{
			name: "normal values",
			data: map[string]any{"stock": 50, "score": -50, "dateExpired": time.Date(2023, 10, 29, 0, 0, 0, 0, time.UTC)},
			want: bson.D{bson.E{Key: "$min", Value: bson.D{bson.E{Key: "stock", Value: 50}, bson.E{Key: "score", Value: -50}, bson.E{Key: "dateExpired", Value: time.Date(2023, 10, 29, 0, 0, 0, 0, time.UTC)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, pkg.EqualBSONDElements(tc.want, BsonBuilder().MinForMap(tc.data).Build()))
		})
	}
}

func Test_fieldUpdateBuilder_Max(t *testing.T) {
	testCases := []struct {
		name string
		data []any

		want bson.D
	}{
		{
			name: "zero params",
			want: bson.D{bson.E{Key: "$max", Value: bson.D{}}},
		},
		{
			name: "odd params",
			data: []any{"stock", 100, "ordered"},
			want: bson.D{bson.E{Key: "$max", Value: bson.D{}}},
		},
		{
			name: "keys contain non-string",
			data: []any{100, "stock", "ordered", 50},
			want: bson.D{bson.E{Key: "$max", Value: bson.D{bson.E{Key: "ordered", Value: 50}}}},
		},
		{
			name: "normal params",
			data: []any{"stock", 100, "dateExpired", time.Date(2023, 10, 29, 0, 0, 0, 0, time.UTC), "score", -50},
			want: bson.D{bson.E{Key: "$max", Value: bson.D{bson.E{Key: "stock", Value: 100}, bson.E{Key: "dateExpired", Value: time.Date(2023, 10, 29, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "score", Value: -50}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().Max(tc.data...).Build())
		})
	}
}

func Test_fieldUpdateBuilder_MaxForMap(t *testing.T) {
	testCases := []struct {
		name string
		data map[string]any

		want bson.D
	}{
		{
			name: "nil values",
			want: bson.D{},
		},
		{
			name: "empty values",
			data: map[string]any{},
			want: bson.D{bson.E{Key: "$max", Value: bson.D{}}},
		},
		{
			name: "normal values",
			data: map[string]any{"stock": 100, "dateExpired": time.Date(2023, 10, 29, 0, 0, 0, 0, time.UTC), "score": -50},
			want: bson.D{bson.E{Key: "$max", Value: bson.D{bson.E{Key: "stock", Value: 100}, bson.E{Key: "dateExpired", Value: time.Date(2023, 10, 29, 0, 0, 0, 0, time.UTC)}, bson.E{Key: "score", Value: -50}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, pkg.EqualBSONDElements(tc.want, BsonBuilder().MaxForMap(tc.data).Build()))
		})
	}
}

func Test_fieldUpdateBuilder_Mul(t *testing.T) {
	testCases := []struct {
		name string
		data []any

		want bson.D
	}{
		{
			name: "zero params",
			want: bson.D{bson.E{Key: "$mul", Value: bson.D{}}},
		},
		{
			name: "odd params",
			data: []any{"price", 1.25, "qty"},
			want: bson.D{bson.E{Key: "$mul", Value: bson.D{}}},
		},
		{
			name: "keys contain non-string",
			data: []any{1.25, "price", "qty", 2},
			want: bson.D{bson.E{Key: "$mul", Value: bson.D{bson.E{Key: "qty", Value: 2}}}},
		},
		{
			name: "values contain non-number",
			data: []any{"price", 1.25, "qty", "2.5"},
			want: bson.D{bson.E{Key: "$mul", Value: bson.D{bson.E{Key: "price", Value: 1.25}}}},
		},
		{
			name: "normal params",
			data: []any{"price", 1.25, "qty", 2, "score", -1, "n", -1.1},
			want: bson.D{bson.E{Key: "$mul", Value: bson.D{bson.E{Key: "price", Value: 1.25}, bson.E{Key: "qty", Value: 2}, bson.E{Key: "score", Value: -1}, bson.E{Key: "n", Value: -1.1}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().Mul(tc.data...).Build())
		})
	}
}

func Test_fieldUpdateBuilder_MulForMap(t *testing.T) {
	testCases := []struct {
		name string
		data map[string]any

		want bson.D
	}{
		{
			name: "nil values",
			want: bson.D{},
		},
		{
			name: "empty values",
			data: map[string]any{},
			want: bson.D{bson.E{Key: "$mul", Value: bson.D{}}},
		},
		{
			name: "values contain non-number",
			data: map[string]any{"price": 1.25, "qty": "2.5"},
			want: bson.D{bson.E{Key: "$mul", Value: bson.D{bson.E{Key: "price", Value: 1.25}}}},
		},
		{
			name: "normal values",
			data: map[string]any{"price": 1.25, "qty": 2, "score": -1, "n": -1.1},
			want: bson.D{bson.E{Key: "$mul", Value: bson.D{bson.E{Key: "price", Value: 1.25}, bson.E{Key: "qty", Value: 2}, bson.E{Key: "score", Value: -1}, bson.E{Key: "n", Value: -1.1}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, pkg.EqualBSONDElements(tc.want, BsonBuilder().MulForMap(tc.data).Build()))
		})
	}
}

func Test_fieldUpdateBuilder_Rename(t *testing.T) {
	testCases := []struct {
		name string
		data []string

		want bson.D
	}{
		{
			name: "zero params",
			want: bson.D{bson.E{Key: "$rename", Value: bson.D{}}},
		},
		{
			name: "odd params",
			data: []string{"name", "cmy", "age"},
			want: bson.D{bson.E{Key: "$rename", Value: bson.D{}}},
		},
		{
			name: "normal params",
			data: []string{"nmae", "name", "name.first", "name.last"},
			want: bson.D{bson.E{Key: "$rename", Value: bson.D{bson.E{Key: "nmae", Value: "name"}, bson.E{Key: "name.first", Value: "name.last"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().Rename(tc.data...).Build())
		})
	}
}

func Test_fieldUpdateBuilder_RenameForMap(t *testing.T) {
	testCases := []struct {
		name string
		data map[string]string

		want bson.D
	}{
		{
			name: "nil values",
			want: bson.D{},
		},
		{
			name: "empty values",
			data: map[string]string{},
			want: bson.D{bson.E{Key: "$rename", Value: bson.D{}}},
		},
		{
			name: "normal values",
			data: map[string]string{"nmae": "name", "name.first": "name.last"},
			want: bson.D{bson.E{Key: "$rename", Value: bson.D{bson.E{Key: "nmae", Value: "name"}, bson.E{Key: "name.first", Value: "name.last"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, pkg.EqualBSONDElements(tc.want, BsonBuilder().RenameForMap(tc.data).Build()))
		})
	}
}
