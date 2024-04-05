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

	"github.com/chenmingyong0423/go-mongox/pkg/utils"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_arrayUpdateBuilder_AddToSet(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(
			t,
			bson.D{bson.E{Key: "$addToSet", Value: bson.D{bson.E{Key: "colors", Value: "mauve"}}}},
			BsonBuilder().AddToSet("colors", "mauve").Build(),
		)
	})

	t.Run("multiple operations", func(t *testing.T) {
		assert.Equal(
			t,
			bson.D{bson.E{Key: "$addToSet", Value: bson.D{bson.E{Key: "colors", Value: "mauve"}, bson.E{Key: "letters", Value: []string{"a", "b", "c"}}}}},
			BsonBuilder().AddToSet("colors", "mauve").AddToSet("letters", []string{"a", "b", "c"}).Build(),
		)
	})
}

func Test_arrayUpdateBuilder_Pop(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$pop", Value: bson.D{bson.E{Key: "scores", Value: 1}}}}, BsonBuilder().Pop("scores", 1).Build())
	})

	t.Run("multiple operations", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$pop", Value: bson.D{bson.E{Key: "scores", Value: 1}, bson.E{Key: "letters", Value: -1}}}}, BsonBuilder().Pop("scores", 1).Pop("letters", -1).Build())
	})
}

func Test_arrayUpdateBuilder_Pull(t *testing.T) {
	// { $pull: { fruits: { $in: [ "apples", "oranges" ] }, votes: { $gte: 6 }, vegetables: "carrots" } }
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(
			t,
			bson.D{bson.E{Key: "$pull", Value: bson.D{bson.E{Key: "fruits", Value: bson.D{bson.E{Key: "$in", Value: []string{"apples", "oranges"}}}}}}},
			BsonBuilder().Pull("fruits", bsonx.D("$in", []string{"apples", "oranges"})).Build(),
		)
	})

}

func Test_arrayUpdateBuilder_Push(t *testing.T) {
	t.Run("single operation", func(t *testing.T) {
		assert.Equal(t,
			bson.D{bson.E{Key: "$push", Value: bson.D{bson.E{Key: "scores", Value: 89}}}},
			BsonBuilder().Push("scores", 89).Build(),
		)
	})
	t.Run("multiple operations", func(t *testing.T) {
		assert.Equal(t,
			bson.D{bson.E{Key: "$push", Value: bson.D{bson.E{Key: "scores", Value: 89}, bson.E{Key: "letters", Value: "a"}}}},
			BsonBuilder().Push("scores", 89).Push("letters", "a").Build(),
		)
	})
}

func Test_arrayUpdateBuilder_PullAll(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []any
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]any)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []any{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []any{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []any{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []any{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []any{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []any{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().PullAll(tc.key, tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_PullAllInt(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []int
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]int)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []int{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []int{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []int{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().PullAllInt(tc.key, tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_PullAllInt8(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []int8
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]int8)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []int8{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int8{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []int8{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int8{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []int8{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int8{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, utils.EqualBSONDElements(tc.want, BsonBuilder().PullAllInt8(tc.key, tc.values...).Build()))
		})
	}
}

func Test_arrayUpdateBuilder_PullAllInt16(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []int16
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]int16)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []int16{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int16{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []int16{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int16{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []int16{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int16{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, utils.EqualBSONDElements(tc.want, BsonBuilder().PullAllInt16(tc.key, tc.values...).Build()))
		})
	}
}

func Test_arrayUpdateBuilder_PullAllInt32(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []int32
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]int32)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []int32{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int32{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []int32{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int32{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []int32{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int32{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, utils.EqualBSONDElements(tc.want, BsonBuilder().PullAllInt32(tc.key, tc.values...).Build()))
		})
	}
}

func Test_arrayUpdateBuilder_PullAllInt64(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []int64
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]int64)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []int64{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int64{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []int64{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int64{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []int64{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []int64{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, utils.EqualBSONDElements(tc.want, BsonBuilder().PullAllInt64(tc.key, tc.values...).Build()))
		})
	}
}

func Test_arrayUpdateBuilder_PullAllString(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []string
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]string)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []string{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []string{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []string{"1"},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []string{"1"}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []string{"1", "2", "3"},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []string{"1", "2", "3"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, utils.EqualBSONDElements(tc.want, BsonBuilder().PullAllString(tc.key, tc.values...).Build()))
		})
	}
}

func Test_arrayUpdateBuilder_PullAllFloat32(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []float32
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]float32)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []float32{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []float32{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []float32{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []float32{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []float32{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []float32{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, utils.EqualBSONDElements(tc.want, BsonBuilder().PullAllFloat32(tc.key, tc.values...).Build()))
		})
	}
}

func Test_arrayUpdateBuilder_PullAllFloat64(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []float64
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]float64)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []float64{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []float64{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []float64{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []float64{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []float64{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []float64{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, utils.EqualBSONDElements(tc.want, BsonBuilder().PullAllFloat64(tc.key, tc.values...).Build()))
		})
	}
}

func Test_arrayUpdateBuilder_PullAllUint(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []uint
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]uint)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []uint{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []uint{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []uint{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, utils.EqualBSONDElements(tc.want, BsonBuilder().PullAllUint(tc.key, tc.values...).Build()))
		})
	}
}

func Test_arrayUpdateBuilder_PullAllUint8(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []uint8
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]uint8)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []uint8{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint8{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []uint8{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint8{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []uint8{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint8{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().PullAllUint8(tc.key, tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_PullAllUint16(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []uint16
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]uint16)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []uint16{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint16{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []uint16{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint16{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []uint16{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint16{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().PullAllUint16(tc.key, tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_PullAllUint32(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []uint32
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]uint32)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []uint32{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint32{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []uint32{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint32{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []uint32{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint32{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().PullAllUint32(tc.key, tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_PullAllUint64(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []uint64
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "scores",
			want: bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: ([]uint64)(nil)}}}},
		},
		{
			name:   "empty values",
			key:    "scores",
			values: []uint64{},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint64{}}}}},
		},
		{
			name:   "single values",
			key:    "scores",
			values: []uint64{1},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint64{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "scores",
			values: []uint64{1, 2, 3},
			want:   bson.D{bson.E{Key: "$pullAll", Value: bson.D{bson.E{Key: "scores", Value: []uint64{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().PullAllUint64(tc.key, tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_Each(t *testing.T) {
	testCases := []struct {
		name   string
		values []any
		want   bson.D
	}{
		{
			name:   "empty values",
			values: []any{},
			want:   bson.D{bson.E{Key: "$each", Value: []any{}}},
		},
		{
			name:   "single values",
			values: []any{"99"},
			want:   bson.D{bson.E{Key: "$each", Value: []any{"99"}}},
		},
		{
			name:   "multiple values",
			values: []any{"99", "98", "97"},
			want:   bson.D{bson.E{Key: "$each", Value: []any{"99", "98", "97"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().Each(tc.values...).Build())
		})
	}
}

func Test_arrayUpdateBuilder_EachInt(t *testing.T) {
	t.Run("test EachInt", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$each", Value: []int{99, 98, 97}}}, BsonBuilder().EachInt(99, 98, 97).Build())
	})
}

func Test_arrayUpdateBuilder_EachInt8(t *testing.T) {
	t.Run("test EachInt8", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$each", Value: []int8{99, 98, 97}}}, BsonBuilder().EachInt8(99, 98, 97).Build())
	})
}

func Test_arrayUpdateBuilder_EachInt16(t *testing.T) {
	t.Run("test EachInt16", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$each", Value: []int16{99, 98, 97}}}, BsonBuilder().EachInt16(99, 98, 97).Build())
	})
}

func Test_arrayUpdateBuilder_EachInt32(t *testing.T) {
	t.Run("test EachInt32", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$each", Value: []int32{99, 98, 97}}}, BsonBuilder().EachInt32(99, 98, 97).Build())
	})
}

func Test_arrayUpdateBuilder_EachInt64(t *testing.T) {
	t.Run("test EachInt64", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$each", Value: []int64{99, 98, 97}}}, BsonBuilder().EachInt64(99, 98, 97).Build())
	})
}

func Test_arrayUpdateBuilder_EachString(t *testing.T) {
	t.Run("test EachString", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$each", Value: []string{"99", "98", "97"}}}, BsonBuilder().EachString("99", "98", "97").Build())
	})
}

func Test_arrayUpdateBuilder_EachFloat32(t *testing.T) {
	t.Run("test EachFloat32", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$each", Value: []float32{99, 98, 97}}}, BsonBuilder().EachFloat32(99, 98, 97).Build())
	})
}

func Test_arrayUpdateBuilder_EachFloat64(t *testing.T) {
	t.Run("test EachFloat64", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$each", Value: []float64{99, 98, 97}}}, BsonBuilder().EachFloat64(99, 98, 97).Build())
	})
}

func Test_arrayUpdateBuilder_EachUint(t *testing.T) {
	t.Run("test EachUint", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$each", Value: []uint{99, 98, 97}}}, BsonBuilder().EachUint(99, 98, 97).Build())
	})
}

func Test_arrayUpdateBuilder_EachUint8(t *testing.T) {
	t.Run("test EachUint8", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$each", Value: []uint8{99, 98, 97}}}, BsonBuilder().EachUint8(99, 98, 97).Build())
	})
}

func Test_arrayUpdateBuilder_EachUint16(t *testing.T) {
	t.Run("test EachUint16", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$each", Value: []uint16{99, 98, 97}}}, BsonBuilder().EachUint16(99, 98, 97).Build())
	})
}

func Test_arrayUpdateBuilder_EachUint32(t *testing.T) {
	t.Run("test EachUint32", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$each", Value: []uint32{99, 98, 97}}}, BsonBuilder().EachUint32(99, 98, 97).Build())
	})
}

func Test_arrayUpdateBuilder_EachUint64(t *testing.T) {
	t.Run("test EachUint64", func(t *testing.T) {
		assert.Equal(t, bson.D{bson.E{Key: "$each", Value: []uint64{99, 98, 97}}}, BsonBuilder().EachUint64(99, 98, 97).Build())
	})
}

func Test_arrayUpdateBuilder_Position(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "$position", Value: 1}}, BsonBuilder().Position(1).Build())
}

func Test_arrayUpdateBuilder_Slice(t *testing.T) {
	assert.Equal(t, bson.D{bson.E{Key: "$slice", Value: 1}}, BsonBuilder().Slice(1).Build())
}

func Test_arrayUpdateBuilder_Sort(t *testing.T) {
	testCases := []struct {
		name  string
		value any
		want  bson.D
	}{
		{
			name:  "number",
			value: 1,
			want:  bson.D{bson.E{Key: "$sort", Value: 1}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().Sort(tc.value).Build())
		})
	}
}
