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

	"github.com/chenmingyong0423/go-mongox/v2/bsonx"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Test_fieldUpdateBuilder_Set(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: "cmy"}}}}, NewBuilder().Set("name", "cmy").Build())
	})
	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: "cmy"}, bson.E{Key: "age", Value: 24}}}}, NewBuilder().Set("name", "cmy").Set("age", 24).Build())
	})
}

func Test_fieldUpdateBuilder_SetFields(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "nil value",
			value: nil,
			want:  bson.D{bson.E{Key: "$set", Value: nil}},
		},
		{
			name:  "string value",
			value: "Alice",
			want:  bson.D{bson.E{Key: "$set", Value: "Alice"}},
		},
		{
			name:  "map value",
			value: map[string]any{"name": "Alice"},
			want:  bson.D{bson.E{Key: "$set", Value: map[string]any{"name": "Alice"}}},
		},
		{
			name:  "bson value",
			value: bson.D{bson.E{Key: "name", Value: "Alice"}},
			want:  bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: "Alice"}}}},
		},
		{
			name:  "pointer struct value",
			value: &bson.D{bson.E{Key: "name", Value: "Alice"}},
			want:  bson.D{bson.E{Key: "$set", Value: &bson.D{bson.E{Key: "name", Value: "Alice"}}}},
		},
		{
			name:  "struct value",
			value: struct{ Name string }{Name: "Alice"},
			want:  bson.D{bson.E{Key: "$set", Value: struct{ Name string }{Name: "Alice"}}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewBuilder().SetFields(tc.value).Build())
		})
	}

	t.Run("multiple set_fields", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$set", Value: struct{ Name string }{Name: "cmy"}}, {Key: "$set", Value: struct{ Age int64 }{Age: 24}}}, NewBuilder().SetFields(struct{ Name string }{Name: "cmy"}).SetFields(struct{ Age int64 }{Age: 24}).Build())
	})

	t.Run("set_fields with struct value and set single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$set", Value: struct{ Name string }{Name: "cmy"}}, {Key: "$set", Value: bson.D{bson.E{Key: "name", Value: "cmy"}}}}, NewBuilder().SetFields(struct{ Name string }{Name: "cmy"}).Set("name", "cmy").Build())
	})

	t.Run("set_fields with pointer struct value and set multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$set", Value: struct{ Name string }{Name: "cmy"}}, {Key: "$set", Value: bson.D{bson.E{Key: "name", Value: "cmy"}, bson.E{Key: "age", Value: 24}}}}, NewBuilder().SetFields(struct{ Name string }{Name: "cmy"}).Set("name", "cmy").Set("age", 24).Build())
	})
}

func Test_fieldUpdateBuilder_Unset(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "$unset", Value: bson.D{}}}, NewBuilder().Unset().Build())
	assert.Equal(t, bson.D{{Key: "$unset", Value: bson.D{bson.E{Key: "name", Value: ""}}}}, NewBuilder().Unset("name").Build())
	assert.Equal(t, bson.D{{Key: "$unset", Value: bson.D{bson.E{Key: "name", Value: ""}, bson.E{Key: "age", Value: ""}}}}, NewBuilder().Unset("name", "age").Build())
}

func Test_fieldUpdateBuilder_SetOnInsert(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$setOnInsert", Value: bson.D{bson.E{Key: "name", Value: "cmy"}}}}, NewBuilder().SetOnInsert("name", "cmy").Build())
	})

	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$setOnInsert", Value: bson.D{bson.E{Key: "name", Value: "cmy"}, bson.E{Key: "age", Value: 24}}}}, NewBuilder().SetOnInsert("name", "cmy").SetOnInsert("age", 24).Build())
	})
}

func Test_fieldUpdateBuilder_Inc(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$inc", Value: bson.D{bson.E{Key: "orders", Value: 1}}}}, NewBuilder().Inc("orders", 1).Build())
	})

	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$inc", Value: bson.D{bson.E{Key: "orders", Value: 1}, bson.E{Key: "ratings", Value: -1}}}}, NewBuilder().Inc("orders", 1).Inc("ratings", -1).Build())
	})
}

func Test_fieldUpdateBuilder_Min(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$min", Value: bson.D{bson.E{Key: "stock", Value: 100}}}}, NewBuilder().Min("stock", 100).Build())
	})
	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$min", Value: bson.D{bson.E{Key: "stock", Value: 100}, bson.E{Key: "dateExpired", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}}}}, NewBuilder().Min("stock", 100).Min("dateExpired", time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)).Build())
	})
}

func Test_fieldUpdateBuilder_Max(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$max", Value: bson.D{bson.E{Key: "stock", Value: 100}}}}, NewBuilder().Max("stock", 100).Build())
	})
	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$max", Value: bson.D{bson.E{Key: "stock", Value: 100}, bson.E{Key: "dateExpired", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}}}}, NewBuilder().Max("stock", 100).Max("dateExpired", time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)).Build())
	})
}

func Test_fieldUpdateBuilder_Mul(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$mul", Value: bson.D{bson.E{Key: "price", Value: 1.25}}}}, NewBuilder().Mul("price", 1.25).Build())
	})
	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$mul", Value: bson.D{bson.E{Key: "price", Value: 1.25}, bson.E{Key: "quantity", Value: 2}}}}, NewBuilder().Mul("price", 1.25).Mul("quantity", 2).Build())
	})
}

func Test_fieldUpdateBuilder_Rename(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$rename", Value: bson.D{bson.E{Key: "name", Value: "name"}}}}, NewBuilder().Rename("name", "name").Build())
	})
	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$rename", Value: bson.D{bson.E{Key: "name", Value: "name"}, bson.E{Key: "age", Value: "age"}}}}, NewBuilder().Rename("name", "name").Rename("age", "age").Build())
	})
}

func Test_fieldUpdateBuilder_CurrentDate(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$currentDate", Value: bson.D{bson.E{Key: "lastModified", Value: true}}}}, NewBuilder().CurrentDate("lastModified", true).Build())
	})

	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$currentDate", Value: bson.D{bson.E{Key: "lastModified", Value: true}, bson.E{Key: "cancellation.date", Value: bson.D{bson.E{Key: "$type", Value: "timestamp"}}}}}}, NewBuilder().CurrentDate("lastModified", true).CurrentDate("cancellation.date", bsonx.D("$type", "timestamp")).Build())
	})
}
