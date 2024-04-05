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

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_fieldUpdateBuilder_Set(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: "cmy"}}}}, BsonBuilder().Set("name", "cmy").Build())
	})
	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: "cmy"}, bson.E{Key: "age", Value: 24}}}}, BsonBuilder().Set("name", "cmy").Set("age", 24).Build())
	})
}

func Test_fieldUpdateBuilder_Unset(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "$unset", Value: bson.D{}}}, BsonBuilder().Unset().Build())
	assert.Equal(t, bson.D{{Key: "$unset", Value: bson.D{bson.E{Key: "name", Value: ""}}}}, BsonBuilder().Unset("name").Build())
	assert.Equal(t, bson.D{{Key: "$unset", Value: bson.D{bson.E{Key: "name", Value: ""}, bson.E{Key: "age", Value: ""}}}}, BsonBuilder().Unset("name", "age").Build())
}

func Test_fieldUpdateBuilder_SetOnInsert(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$setOnInsert", Value: bson.D{bson.E{Key: "name", Value: "cmy"}}}}, BsonBuilder().SetOnInsert("name", "cmy").Build())
	})

	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$setOnInsert", Value: bson.D{bson.E{Key: "name", Value: "cmy"}, bson.E{Key: "age", Value: 24}}}}, BsonBuilder().SetOnInsert("name", "cmy").SetOnInsert("age", 24).Build())
	})
}

func Test_fieldUpdateBuilder_Inc(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$inc", Value: bson.D{bson.E{Key: "orders", Value: 1}}}}, BsonBuilder().Inc("orders", 1).Build())
	})

	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$inc", Value: bson.D{bson.E{Key: "orders", Value: 1}, bson.E{Key: "ratings", Value: -1}}}}, BsonBuilder().Inc("orders", 1).Inc("ratings", -1).Build())
	})
}

func Test_fieldUpdateBuilder_Min(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$min", Value: bson.D{bson.E{Key: "stock", Value: 100}}}}, BsonBuilder().Min("stock", 100).Build())
	})
	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$min", Value: bson.D{bson.E{Key: "stock", Value: 100}, bson.E{Key: "dateExpired", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}}}}, BsonBuilder().Min("stock", 100).Min("dateExpired", time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)).Build())
	})
}

func Test_fieldUpdateBuilder_Max(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$max", Value: bson.D{bson.E{Key: "stock", Value: 100}}}}, BsonBuilder().Max("stock", 100).Build())
	})
	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$max", Value: bson.D{bson.E{Key: "stock", Value: 100}, bson.E{Key: "dateExpired", Value: time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)}}}}, BsonBuilder().Max("stock", 100).Max("dateExpired", time.Date(2023, 10, 24, 0, 0, 0, 0, time.UTC)).Build())
	})
}

func Test_fieldUpdateBuilder_Mul(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$mul", Value: bson.D{bson.E{Key: "price", Value: 1.25}}}}, BsonBuilder().Mul("price", 1.25).Build())
	})
	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$mul", Value: bson.D{bson.E{Key: "price", Value: 1.25}, bson.E{Key: "quantity", Value: 2}}}}, BsonBuilder().Mul("price", 1.25).Mul("quantity", 2).Build())
	})
}

func Test_fieldUpdateBuilder_Rename(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$rename", Value: bson.D{bson.E{Key: "name", Value: "name"}}}}, BsonBuilder().Rename("name", "name").Build())
	})
	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$rename", Value: bson.D{bson.E{Key: "name", Value: "name"}, bson.E{Key: "age", Value: "age"}}}}, BsonBuilder().Rename("name", "name").Rename("age", "age").Build())
	})
}

func Test_fieldUpdateBuilder_CurrentDate(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$currentDate", Value: bson.D{bson.E{Key: "lastModified", Value: true}}}}, BsonBuilder().CurrentDate("lastModified", true).Build())
	})

	t.Run("multiple operation", func(t *testing.T) {
		assert.Equal(t, bson.D{{Key: "$currentDate", Value: bson.D{bson.E{Key: "lastModified", Value: true}, bson.E{Key: "cancellation.date", Value: bson.D{bson.E{Key: "$type", Value: "timestamp"}}}}}}, BsonBuilder().CurrentDate("lastModified", true).CurrentDate("cancellation.date", bsonx.D("$type", "timestamp")).Build())
	})
}
