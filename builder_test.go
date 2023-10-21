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

package mongox

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestBsonBuilder(t *testing.T) {
	assert.Equal(t, bson.D{}, NewBsonBuilder().Build())

	// Id()
	assert.Equal(t, bson.D{bson.E{Key: id, Value: 123}}, NewBsonBuilder().Id(123).Build())

	// Add()
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: "v1"}, bson.E{Key: "k2", Value: "v2"}}, NewBsonBuilder().Add("k1", "v1").Add("k2", "v2").Build())

	// Set()
	assert.Equal(t, bson.D{bson.E{Key: set, Value: bson.D{bson.E{Key: "k1", Value: "v1"}}}}, NewBsonBuilder().Set("k1", "v1").Build())

	// SetForMap()
	assert.Equal(t, bson.D{}, NewBsonBuilder().SetForMap(nil).Build())
	assert.Equal(t, bson.D{}, NewBsonBuilder().SetForMap(map[string]any{}).Build())
	assert.ElementsMatch(t, bson.D{bson.E{Key: set, Value: bson.D{
		bson.E{Key: "k1", Value: "v1"},
		bson.E{Key: "k2", Value: "v2"},
		bson.E{Key: "k3", Value: "v3"},
	}}}[0].Value, NewBsonBuilder().SetForMap(map[string]any{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
	}).Build()[0].Value)

	// SetForStruct()
	assert.ElementsMatch(t, bson.D{}, NewBsonBuilder().SetForStruct(nil).Build())
	assert.ElementsMatch(t, bson.D{}, NewBsonBuilder().SetForStruct(
		func() **testData {
			data := &testData{}
			return &data
		}()).Build())
	assert.ElementsMatch(t, bson.D{bson.E{Key: set, Value: bson.D{
		bson.E{Key: "_id", Value: "123"},
		bson.E{Key: "name", Value: "cmy"},
		bson.E{Key: "age", Value: 0},
	}}}[0].Value, NewBsonBuilder().SetForStruct(testData{
		Id:      "123",
		Name:    "cmy",
		Unknown: "",
	}).Build()[0].Value)
	assert.ElementsMatch(t, bson.D{bson.E{Key: set, Value: bson.D{
		bson.E{Key: "_id", Value: "123"},
		bson.E{Key: "name", Value: "cmy"},
		bson.E{Key: "age", Value: 18},
	}}}[0].Value, NewBsonBuilder().SetForStruct(testData{
		Id:      "123",
		Name:    "cmy",
		Age:     18,
		Unknown: "",
	}).Build()[0].Value)

	// in
	ints := []any{1, 2, 3, 4}
	assert.Equal(t, bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: in, Value: ints}}}}, NewBsonBuilder().In("id", ints...).Build())

	// eq
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{eq: "v1"}}}, NewBsonBuilder().Eq("k1", "v1").Build())

	// gt
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{gt: "v1"}}}, NewBsonBuilder().Gt("k1", "v1").Build())

	// gte
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{gte: "v1"}}}, NewBsonBuilder().Gte("k1", "v1").Build())

	// lt
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{lt: "v1"}}}, NewBsonBuilder().Lt("k1", "v1").Build())

	// lte
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{lte: "v1"}}}, NewBsonBuilder().Lte("k1", "v1").Build())

	// ne
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{ne: "v1"}}}, NewBsonBuilder().Ne("k1", "v1").Build())

	// nin
	assert.Equal(t, bson.D{bson.E{Key: "id", Value: bson.M{nin: ints}}}, NewBsonBuilder().Nin("id", ints...).Build())

}

type testData struct {
	Id      string `bson:"_id"`
	Name    string `bson:"name"`
	Age     int
	Unknown string `bson:"-"`
}
