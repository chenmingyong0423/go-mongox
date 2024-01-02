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

	"github.com/chenmingyong0423/go-mongox/bsonx"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestAddToSet(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "test AddToSet",
			value: bsonx.M("colors", "mauve"),
			want:  bson.D{bson.E{Key: "$addToSet", Value: bsonx.M("colors", "mauve")}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, AddToSet(tc.value))
		})
	}
}

func TestPop(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "test Pop",
			value: bsonx.M("scores", 1),
			want:  bson.D{bson.E{Key: "$pop", Value: bsonx.M("scores", 1)}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Pop(tc.value))
		})
	}
}

func TestPull(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "test Pull",
			value: bsonx.M("fruit", "apples"),
			want:  bson.D{bson.E{Key: "$pull", Value: bsonx.M("fruit", "apples")}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Pull(tc.value))
		})
	}
}

func TestPush(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "test Push",
			value: bsonx.M("scores", 89),
			want:  bson.D{bson.E{Key: "$push", Value: bsonx.M("scores", 89)}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Push(tc.value))
		})
	}
}

func TestPullAll(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []any
		want   bson.D
	}{
		{
			name:   "test PullAll",
			key:    "letters",
			values: []any{"b", "c"},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "letters", Value: []any{"b", "c"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, PullAll(tc.key, tc.values...))
		})
	}
}

func TestEach(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []any
		want   bson.D
	}{
		{
			name:   "test Each",
			key:    "scores",
			values: []any{3, 4, 5},
			want:   bson.D{bson.E{Key: "scores", Value: bson.D{bson.E{Key: "$each", Value: []any{3, 4, 5}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Each(tc.key, tc.values...))
		})
	}
}

func TestPosition(t *testing.T) {
	testCases := []struct {
		name  string
		key   string
		value any
		want  bson.D
	}{
		{
			name:  "test Position",
			key:   "scores",
			value: 3,
			want:  bson.D{bson.E{Key: "scores", Value: bson.D{bson.E{Key: "$position", Value: 3}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Position(tc.key, tc.value))
		})
	}
}

func TestSlice(t *testing.T) {
	testCases := []struct {
		name string
		key  string
		num  int
		want bson.D
	}{
		{
			name: "test Slice",
			key:  "scores",
			num:  3,
			want: bson.D{bson.E{Key: "scores", Value: bson.D{bson.E{Key: "$slice", Value: 3}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Slice(tc.key, tc.num))
		})
	}
}

func TestSort(t *testing.T) {
	testCases := []struct {
		name  string
		key   string
		value any
		want  bson.D
	}{
		{
			name:  "test Sort",
			key:   "scores",
			value: 1,
			want:  bson.D{bson.E{Key: "scores", Value: bson.D{bson.E{Key: "$sort", Value: 1}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Sort(tc.key, tc.value))
		})
	}
}

func TestSet(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "test Set",
			value: bsonx.M("name", "Alice"),
			want:  bson.D{bson.E{Key: "$set", Value: bson.M{"name": "Alice"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Set(tc.value))
		})
	}
}

func TestUnset(t *testing.T) {
	testCases := []struct {
		name  string
		value []string
		want  bson.D
	}{
		{
			name:  "test Unset",
			value: []string{"name", "age"},
			want:  bson.D{bson.E{Key: "$unset", Value: bson.D{bson.E{Key: "name", Value: ""}, bson.E{Key: "age", Value: ""}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, Unset(tc.value...))
		})
	}
}

func TestSetOnInsert(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "test SetOnInsert",
			value: bsonx.M("name", "Alice"),
			want:  bson.D{bson.E{Key: "$setOnInsert", Value: bson.M{"name": "Alice"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, SetOnInsert(tc.value))
		})
	}
}

func TestCurrentDate(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "test CurrentDate",
			value: bsonx.NewD().Add("lastModified", true).Add("cancellation.date", bsonx.M("$type", "timestamp")).Build(),
			want:  bson.D{bson.E{Key: "$currentDate", Value: bson.D{bson.E{Key: "lastModified", Value: true}, bson.E{Key: "cancellation.date", Value: bson.M{"$type": "timestamp"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, CurrentDate(tc.value))
		})
	}
}

func TestInc(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "test Inc",
			value: bsonx.M("age", 1),
			want:  bson.D{bson.E{Key: "$inc", Value: bson.M{"age": 1}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, bson.D{bson.E{Key: "$inc", Value: bson.M{"age": 1}}}, Inc(tc.value))
		})
	}
}

func TestMin(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "test Min",
			value: bsonx.M("age", 18),
			want:  bson.D{bson.E{Key: "$min", Value: bson.M{"age": 18}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, bson.D{bson.E{Key: "$min", Value: bson.M{"age": 18}}}, Min(tc.value))
		})
	}
}

func TestMax(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "test Max",
			value: bsonx.M("age", 18),
			want:  bson.D{bson.E{Key: "$max", Value: bson.M{"age": 18}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, bson.D{bson.E{Key: "$max", Value: bson.M{"age": 18}}}, Max(tc.value))
		})
	}
}

func TestMul(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "test Mul",
			value: bsonx.M("age", 2),
			want:  bson.D{bson.E{Key: "$mul", Value: bson.M{"age": 2}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, bson.D{bson.E{Key: "$mul", Value: bson.M{"age": 2}}}, Mul(tc.value))
		})
	}
}

func TestRename(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "test Rename",
			value: bsonx.M("name", "nickname"),
			want:  bson.D{bson.E{Key: "$rename", Value: bson.M{"name": "nickname"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, bson.D{bson.E{Key: "$rename", Value: bson.M{"name": "nickname"}}}, Rename(tc.value))
		})
	}
}

func TestSetSimple(t *testing.T) {
	testCases := []struct {
		name  string
		key   string
		value any
		want  bson.D
	}{
		{
			name:  "test Set",
			key:   "name",
			value: "chenmingyong",
			want:  bson.D{bson.E{Key: "$set", Value: bson.E{Key: "name", Value: "chenmingyong"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, SetSimple(tc.key, tc.value))
		})
	}
}
