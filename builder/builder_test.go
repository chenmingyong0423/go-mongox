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

package builder

import (
	"testing"

	"github.com/chenmingyong0423/go-mongox/pkg"

	"go.mongodb.org/mongo-driver/bson/bsontype"

	"github.com/chenmingyong0423/go-mongox/internal/types"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestBsonBuilder(t *testing.T) {
	assert.Equal(t, bson.D{}, NewBsonBuilder().Build())

	// Id()
	assert.Equal(t, bson.D{bson.E{Key: types.Id, Value: 123}}, NewBsonBuilder().Id(123).Build())

	// Add()
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: "v1"}, bson.E{Key: "k2", Value: "v2"}}, NewBsonBuilder().Add("k1", "v1").Add("k2", "v2").Build())

	// Set()
	assert.Equal(t, bson.D{bson.E{Key: types.Set, Value: bson.D{bson.E{Key: "k1", Value: "v1"}}}}, NewBsonBuilder().Set("k1", "v1").Build())

	// SetForMap()
	assert.Equal(t, bson.D{}, NewBsonBuilder().SetForMap(nil).Build())
	assert.Equal(t, bson.D{}, NewBsonBuilder().SetForMap(map[string]any{}).Build())
	assert.ElementsMatch(t, bson.D{bson.E{Key: types.Set, Value: bson.D{
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
	assert.ElementsMatch(t, bson.D{bson.E{Key: types.Set, Value: bson.D{
		bson.E{Key: "_id", Value: "123"},
		bson.E{Key: "name", Value: "cmy"},
		bson.E{Key: "age", Value: int32(0)},
	}}}[0].Value, NewBsonBuilder().SetForStruct(testData{
		Id:      "123",
		Name:    "cmy",
		Unknown: "",
	}).Build()[0].Value)
	assert.ElementsMatch(t, bson.D{bson.E{Key: types.Set, Value: bson.D{
		bson.E{Key: "_id", Value: "123"},
		bson.E{Key: "name", Value: "cmy"},
		bson.E{Key: "age", Value: int32(18)},
	}}}[0].Value, NewBsonBuilder().SetForStruct(testData{
		Id:      "123",
		Name:    "cmy",
		Age:     18,
		Unknown: "",
	}).Build()[0].Value)

	// eq
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{types.Eq: "v1"}}}, NewBsonBuilder().Eq("k1", "v1").Build())

	// gt
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{types.Gt: "v1"}}}, NewBsonBuilder().Gt("k1", "v1").Build())

	// gte
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{types.Gte: "v1"}}}, NewBsonBuilder().Gte("k1", "v1").Build())

	// lt
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{types.Lt: "v1"}}}, NewBsonBuilder().Lt("k1", "v1").Build())

	// lte
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{types.Lte: "v1"}}}, NewBsonBuilder().Lte("k1", "v1").Build())

	// ne
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{types.Ne: "v1"}}}, NewBsonBuilder().Ne("k1", "v1").Build())

	// And
	assert.Equal(t, bson.D{bson.E{Key: types.And, Value: []bson.D{
		{bson.E{Key: "k1", Value: "v1"}},
		{bson.E{Key: "k2", Value: "v2"}},
	}}}, NewBsonBuilder().And(
		bson.D{bson.E{Key: "k1", Value: "v1"}},
		bson.D{bson.E{Key: "k2", Value: "v2"}},
	).Build())
	assert.Equal(t, bson.D{bson.E{Key: types.And, Value: []bson.D{
		{bson.E{Key: "k1", Value: "v1"}},
		{bson.E{Key: "k2", Value: "v2"}},
	}}}, NewBsonBuilder().And(NewBsonBuilder().Add("k1", "v1").Build(), NewBsonBuilder().Add("k2", "v2").Build()).Build())

	// Not 测试用例
	assert.Equal(t, bson.D{bson.E{Key: types.Not, Value: bson.D{bson.E{Key: "k1", Value: "v1"}}}}, NewBsonBuilder().Not(bson.D{bson.E{Key: "k1", Value: "v1"}}).Build())

	// Nor
	assert.Equal(t, bson.D{bson.E{Key: types.Nor, Value: []bson.D{
		{bson.E{Key: "k1", Value: "v1"}},
		{bson.E{Key: "k2", Value: "v2"}},
	}}}, NewBsonBuilder().Nor(
		bson.D{bson.E{Key: "k1", Value: "v1"}},
		bson.D{bson.E{Key: "k2", Value: "v2"}},
	).Build())
	assert.Equal(t, bson.D{bson.E{Key: types.Nor, Value: []bson.D{
		{bson.E{Key: "k1", Value: "v1"}},
		{bson.E{Key: "k2", Value: "v2"}},
	}}}, NewBsonBuilder().Nor(NewBsonBuilder().Add("k1", "v1").Build(), NewBsonBuilder().Add("k2", "v2").Build()).Build())

	// Or
	assert.Equal(t, bson.D{bson.E{Key: types.Or, Value: []bson.D{
		{bson.E{Key: "k1", Value: "v1"}},
		{bson.E{Key: "k2", Value: "v2"}},
	}}}, NewBsonBuilder().Or(
		bson.D{bson.E{Key: "k1", Value: "v1"}},
		bson.D{bson.E{Key: "k2", Value: "v2"}},
	).Build())
	assert.Equal(t, bson.D{bson.E{Key: types.Or, Value: []bson.D{
		{bson.E{Key: "k1", Value: "v1"}},
		{bson.E{Key: "k2", Value: "v2"}},
	}}}, NewBsonBuilder().Or(NewBsonBuilder().Add("k1", "v1").Build(), NewBsonBuilder().Add("k2", "v2").Build()).Build())

	// Exists
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{types.Exists: true}}}, NewBsonBuilder().Exists("k1", true).Build())

	// Type
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{types.Type: bson.TypeString}}}, NewBsonBuilder().Type("k1", bson.TypeString).Build())

	// TypeAlias
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{types.Type: "string"}}}, NewBsonBuilder().TypeAlias("k1", "string").Build())

	// ElemMatch
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{types.ElemMatch: bson.D{bson.E{Key: "k2", Value: "v2"}}}}}, NewBsonBuilder().ElemMatch("k1", bson.D{bson.E{Key: "k2", Value: "v2"}}).Build())
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{types.ElemMatch: NewBsonBuilder().Add("k2", "v2").Build()}}}, NewBsonBuilder().ElemMatch("k1", bson.D{bson.E{Key: "k2", Value: "v2"}}).Build())

	// Size
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{types.Size: 1}}}, NewBsonBuilder().Size("k1", 1).Build())
}

type testData struct {
	Id      string `bson:"_id"`
	Name    string `bson:"name"`
	Age     int
	Unknown string `bson:"-"`
}

func TestBsonBuilder_In(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []any
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{bson.E{Key: "id", Value: bson.M{types.In: ([]any)(nil)}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []any{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []any{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{1}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []any{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{1, 2, 3}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewBsonBuilder().In(tc.key, tc.values...).Build())
		})
	}
}

func TestBsonBuilder_InUint(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []uint
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{bson.E{Key: "id", Value: bson.M{types.In: ([]any)(nil)}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []uint{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []uint{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{uint(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []uint{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{uint(1), uint(2), uint(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewBsonBuilder().InUint(tc.key, tc.values...).Build())
		})
	}
}

func TestBsonBuilder_InUint8(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []uint8
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{bson.E{Key: "id", Value: bson.M{types.In: ([]any)(nil)}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []uint8{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []uint8{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{uint8(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []uint8{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{uint8(1), uint8(2), uint8(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewBsonBuilder().InUint8(tc.key, tc.values...).Build())
		})
	}
}

func TestBsonBuilder_InUint16(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []uint16
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{bson.E{Key: "id", Value: bson.M{types.In: ([]any)(nil)}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []uint16{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []uint16{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{uint16(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []uint16{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{uint16(1), uint16(2), uint16(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewBsonBuilder().InUint16(tc.key, tc.values...).Build())
		})
	}
}

func TestBsonBuilder_InUint32(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []uint32
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{bson.E{Key: "id", Value: bson.M{types.In: ([]any)(nil)}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []uint32{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []uint32{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{uint32(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []uint32{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{uint32(1), uint32(2), uint32(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewBsonBuilder().InUint32(tc.key, tc.values...).Build())
		})
	}
}

func TestBsonBuilder_InUint64(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []uint64
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{bson.E{Key: "id", Value: bson.M{types.In: ([]any)(nil)}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []uint64{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []uint64{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{uint64(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []uint64{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{uint64(1), uint64(2), uint64(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewBsonBuilder().InUint64(tc.key, tc.values...).Build())
		})
	}
}

func TestBsonBuilder_InInt(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []int
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{bson.E{Key: "id", Value: bson.M{types.In: ([]any)(nil)}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []int{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []int{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{1}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []int{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{1, 2, 3}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewBsonBuilder().InInt(tc.key, tc.values...).Build())
		})
	}
}

func TestBsonBuilder_InInt8(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []int8
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{bson.E{Key: "id", Value: bson.M{types.In: ([]any)(nil)}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []int8{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []int8{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{int8(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []int8{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{int8(1), int8(2), int8(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewBsonBuilder().InInt8(tc.key, tc.values...).Build())
		})
	}
}

func TestBsonBuilder_InInt16(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []int16
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{bson.E{Key: "id", Value: bson.M{types.In: ([]any)(nil)}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []int16{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []int16{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{int16(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []int16{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{int16(1), int16(2), int16(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewBsonBuilder().InInt16(tc.key, tc.values...).Build())
		})
	}
}

func TestBsonBuilder_InInt32(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []int32
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{bson.E{Key: "id", Value: bson.M{types.In: ([]any)(nil)}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []int32{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []int32{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{int32(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []int32{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{int32(1), int32(2), int32(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewBsonBuilder().InInt32(tc.key, tc.values...).Build())
		})
	}
}

func TestBsonBuilder_InInt64(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []int64
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{bson.E{Key: "id", Value: bson.M{types.In: ([]any)(nil)}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []int64{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []int64{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{int64(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []int64{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{int64(1), int64(2), int64(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewBsonBuilder().InInt64(tc.key, tc.values...).Build())
		})
	}
}

func TestBsonBuilder_InFloat32(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []float32
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{bson.E{Key: "id", Value: bson.M{types.In: ([]any)(nil)}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []float32{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []float32{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{float32(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []float32{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{float32(1), float32(2), float32(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewBsonBuilder().InFloat32(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_InFloat64(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []float64
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{bson.E{Key: "id", Value: bson.M{types.In: ([]any)(nil)}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []float64{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{}}}}},
		{
			name:   "one value",
			key:    "id",
			values: []float64{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{float64(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []float64{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{float64(1), float64(2), float64(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewBsonBuilder().InFloat64(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_InString(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []string
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{bson.E{Key: "id", Value: bson.M{types.In: ([]any)(nil)}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []string{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []string{"1"},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{"1"}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []string{"1", "2", "3"},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.In: []any{"1", "2", "3"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewBsonBuilder().InString(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_toAnySlice(t *testing.T) {

	testCases := []struct {
		name   string
		values []int64
		want   []any
	}{
		{
			name: "nil values",
			want: []any(nil),
		},
		{
			name:   "empty values",
			values: []int64{},
			want:   []any{},
		},
		{
			name:   "one value",
			values: []int64{1},
			want:   []any{int64(1)},
		},
		{
			name:   "multiple values",
			values: []int64{1, 2, 3},
			want:   []any{int64(1), int64(2), int64(3)},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, pkg.ToAnySlice(tc.values...))
		})
	}
}

func TestBsonBuilder_Nin(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []any
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: ([]any)(nil)}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []any{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []any{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{1}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []any{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{1, 2, 3}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewBsonBuilder().Nin(tc.key, tc.values...).Build())
		})
	}
}

func TestBsonBuilder_NinUint(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []uint
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: ([]any)(nil)}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []uint{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []uint{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{uint(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []uint{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{uint(1), uint(2), uint(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewBsonBuilder().NinUint(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_NinUint8(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []uint8
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{
				bson.E{Key: "id", Value: bson.M{types.Nin: ([]any)(nil)}},
			},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []uint8{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []uint8{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{uint8(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []uint8{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{uint8(1), uint8(2), uint8(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewBsonBuilder().NinUint8(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_NinUint16(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []uint16
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{
				bson.E{Key: "id", Value: bson.M{types.Nin: ([]any)(nil)}},
			},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []uint16{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []uint16{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{uint16(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []uint16{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{uint16(1), uint16(2), uint16(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewBsonBuilder().NinUint16(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_NinUint32(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []uint32
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{
				bson.E{Key: "id", Value: bson.M{types.Nin: ([]any)(nil)}},
			},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []uint32{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []uint32{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{uint32(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []uint32{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{uint32(1), uint32(2), uint32(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewBsonBuilder().NinUint32(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_NinUint64(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []uint64
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{
				bson.E{Key: "id", Value: bson.M{types.Nin: ([]any)(nil)}},
			},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []uint64{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []uint64{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{uint64(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []uint64{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{uint64(1), uint64(2), uint64(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := NewBsonBuilder().NinUint64(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_NinInt(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []int
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{
				bson.E{Key: "id", Value: bson.M{types.Nin: ([]any)(nil)}},
			},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []int{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []int{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{1}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []int{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{1, 2, 3}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := NewBsonBuilder().NinInt(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_NinInt8(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []int8
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{
				bson.E{Key: "id", Value: bson.M{types.Nin: ([]any)(nil)}},
			},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []int8{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []int8{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{int8(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []int8{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{int8(1), int8(2), int8(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := NewBsonBuilder().NinInt8(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_NinInt16(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []int16
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{
				bson.E{Key: "id", Value: bson.M{types.Nin: ([]any)(nil)}},
			},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []int16{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []int16{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{int16(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []int16{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{int16(1), int16(2), int16(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := NewBsonBuilder().NinInt16(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_NinInt32(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []int32
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{
				bson.E{Key: "id", Value: bson.M{types.Nin: ([]any)(nil)}},
			},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []int32{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []int32{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{int32(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []int32{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{int32(1), int32(2), int32(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := NewBsonBuilder().NinInt32(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_NinInt64(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []int64
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{
				bson.E{Key: "id", Value: bson.M{types.Nin: ([]any)(nil)}},
			},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []int64{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []int64{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{int64(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []int64{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{int64(1), int64(2), int64(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := NewBsonBuilder().NinInt64(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_NinFloat32(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []float32
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{
				bson.E{Key: "id", Value: bson.M{types.Nin: ([]any)(nil)}},
			},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []float32{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []float32{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{float32(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []float32{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{float32(1), float32(2), float32(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := NewBsonBuilder().NinFloat32(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_NinFloat64(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []float64
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{
				bson.E{Key: "id", Value: bson.M{types.Nin: ([]any)(nil)}},
			},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []float64{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []float64{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{float64(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []float64{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{float64(1), float64(2), float64(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := NewBsonBuilder().NinFloat64(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_NinString(t *testing.T) {
	testCases := []struct {
		name   string
		key    string
		values []string
		want   bson.D
	}{
		{
			name: "nil values",
			key:  "id",
			want: bson.D{
				bson.E{Key: "id", Value: bson.M{types.Nin: ([]any)(nil)}},
			},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []string{},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []string{"1"},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{"1"}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []string{"1", "2", "3"},
			want:   bson.D{bson.E{Key: "id", Value: bson.M{types.Nin: []any{"1", "2", "3"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := NewBsonBuilder().NinString(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_TypeArray(t *testing.T) {

	testCases := []struct {
		name string
		key  string
		ts   []bsontype.Type

		want bson.D
	}{
		{
			name: "nil values",
			key:  "name",
			want: bson.D{
				bson.E{Key: "name", Value: bson.M{types.Type: ([]bsontype.Type)(nil)}},
			},
		},
		{
			name: "empty values",
			key:  "name",
			ts:   []bsontype.Type{},
			want: bson.D{
				bson.E{Key: "name", Value: bson.M{types.Type: []bsontype.Type{}}},
			},
		},
		{
			name: "one value",
			key:  "name",
			ts:   []bsontype.Type{bson.TypeString},
			want: bson.D{
				bson.E{Key: "name", Value: bson.M{types.Type: []bsontype.Type{bson.TypeString}}},
			},
		},
		{
			name: "multiple values",
			key:  "name",
			ts:   []bsontype.Type{bson.TypeString, bson.TypeInt32},
			want: bson.D{
				bson.E{Key: "name", Value: bson.M{types.Type: []bsontype.Type{bson.TypeString, bson.TypeInt32}}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewBsonBuilder().TypeArray(tc.key, tc.ts...).Build())
		})
	}
}

func TestBsonBuilder_TypeArrayAlias(t *testing.T) {

	testCases := []struct {
		name string
		key  string
		ts   []string

		want bson.D
	}{
		{
			name: "nil values",
			key:  "name",
			want: bson.D{
				bson.E{Key: "name", Value: bson.M{types.Type: ([]string)(nil)}},
			},
		},
		{
			name: "empty values",
			key:  "name",
			ts:   []string{},
			want: bson.D{
				bson.E{Key: "name", Value: bson.M{types.Type: []string{}}},
			},
		},
		{
			name: "one value",
			key:  "name",
			ts:   []string{"string"},
			want: bson.D{
				bson.E{Key: "name", Value: bson.M{types.Type: []string{"string"}}},
			},
		},
		{
			name: "multiple values",
			key:  "name",
			ts:   []string{"string", "int32"},
			want: bson.D{
				bson.E{Key: "name", Value: bson.M{types.Type: []string{"string", "int32"}}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewBsonBuilder().TypeArrayAlias(tc.key, tc.ts...).Build())
		})
	}
}

func TestBsonBuilder_All(t *testing.T) {

	testCases := []struct {
		name   string
		key    string
		values []any

		want bson.D
	}{
		{
			name: "nil values",
			key:  "name",
			want: bson.D{
				bson.E{Key: "name", Value: bson.M{types.All: ([]any)(nil)}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []any{},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []any{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{1}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []any{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{1, 2, 3}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, NewBsonBuilder().All(tc.key, tc.values...).Build())
		})
	}
}

func TestBsonBuilder_AllUint(t *testing.T) {

	testCases := []struct {
		name   string
		key    string
		values []uint

		want bson.D
	}{
		{
			name: "nil values",
			key:  "name",
			want: bson.D{
				bson.E{Key: "name", Value: bson.M{types.All: ([]any)(nil)}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []uint{},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []uint{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{uint(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []uint{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{uint(1), uint(2), uint(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewBsonBuilder().AllUint(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllUint8(t *testing.T) {

	testCases := []struct {
		name   string
		key    string
		values []uint8

		want bson.D
	}{
		{
			name: "nil values",
			key:  "name",
			want: bson.D{
				bson.E{Key: "name", Value: bson.M{types.All: ([]any)(nil)}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []uint8{},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []uint8{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{uint8(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []uint8{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{uint8(1), uint8(2), uint8(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewBsonBuilder().AllUint8(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllUint16(t *testing.T) {

	testCases := []struct {
		name   string
		key    string
		values []uint16

		want bson.D
	}{
		{
			name: "nil values",
			key:  "name",
			want: bson.D{
				bson.E{Key: "name", Value: bson.M{types.All: ([]any)(nil)}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []uint16{},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []uint16{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{uint16(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []uint16{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{uint16(1), uint16(2), uint16(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewBsonBuilder().AllUint16(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllUint32(t *testing.T) {

	testCases := []struct {
		name   string
		key    string
		values []uint32

		want bson.D
	}{
		{
			name: "nil values",
			key:  "name",
			want: bson.D{
				bson.E{Key: "name", Value: bson.M{types.All: ([]any)(nil)}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []uint32{},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []uint32{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{uint32(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []uint32{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{uint32(1), uint32(2), uint32(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewBsonBuilder().AllUint32(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllUint64(t *testing.T) {

	testCases := []struct {
		name   string
		key    string
		values []uint64

		want bson.D
	}{
		{
			name: "nil values",
			key:  "name",
			want: bson.D{
				bson.E{Key: "name", Value: bson.M{types.All: ([]any)(nil)}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []uint64{},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []uint64{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{uint64(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []uint64{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{uint64(1), uint64(2), uint64(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewBsonBuilder().AllUint64(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllInt(t *testing.T) {

	testCases := []struct {
		name   string
		key    string
		values []int

		want bson.D
	}{
		{
			name: "nil values",
			key:  "name",
			want: bson.D{
				bson.E{Key: "name", Value: bson.M{types.All: ([]any)(nil)}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []int{},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []int{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{1}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []int{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{1, 2, 3}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewBsonBuilder().AllInt(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllInt8(t *testing.T) {

	testCases := []struct {
		name   string
		key    string
		values []int8

		want bson.D
	}{
		{
			name: "nil values",
			key:  "name",
			want: bson.D{
				bson.E{Key: "name", Value: bson.M{types.All: ([]any)(nil)}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []int8{},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []int8{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{int8(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []int8{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{int8(1), int8(2), int8(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := NewBsonBuilder().AllInt8(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllInt16(t *testing.T) {

	testCases := []struct {
		name   string
		key    string
		values []int16

		want bson.D
	}{
		{
			name: "nil values",
			key:  "name",
			want: bson.D{
				bson.E{Key: "name", Value: bson.M{types.All: ([]any)(nil)}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []int16{},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []int16{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{int16(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []int16{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{int16(1), int16(2), int16(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := NewBsonBuilder().AllInt16(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllInt32(t *testing.T) {

	testCases := []struct {
		name   string
		key    string
		values []int32

		want bson.D
	}{
		{
			name: "nil values",
			key:  "name",
			want: bson.D{
				bson.E{Key: "name", Value: bson.M{types.All: ([]any)(nil)}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []int32{},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []int32{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{int32(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []int32{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{int32(1), int32(2), int32(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := NewBsonBuilder().AllInt32(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllInt64(t *testing.T) {

	testCases := []struct {
		name   string
		key    string
		values []int64

		want bson.D
	}{
		{
			name: "nil values",
			key:  "name",
			want: bson.D{
				bson.E{Key: "name", Value: bson.M{types.All: ([]any)(nil)}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []int64{},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []int64{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{int64(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []int64{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{int64(1), int64(2), int64(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := NewBsonBuilder().AllInt64(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllFloat32(t *testing.T) {

	testCases := []struct {
		name   string
		key    string
		values []float32

		want bson.D
	}{
		{
			name: "nil values",
			key:  "name",
			want: bson.D{
				bson.E{Key: "name", Value: bson.M{types.All: ([]any)(nil)}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []float32{},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []float32{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{float32(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []float32{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{float32(1), float32(2), float32(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewBsonBuilder().AllFloat32(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllFloat64(t *testing.T) {

	testCases := []struct {
		name   string
		key    string
		values []float64

		want bson.D
	}{
		{
			name: "nil values",
			key:  "name",
			want: bson.D{
				bson.E{Key: "name", Value: bson.M{types.All: ([]any)(nil)}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []float64{},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []float64{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{float64(1)}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []float64{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{float64(1), float64(2), float64(3)}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewBsonBuilder().AllFloat64(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllString(t *testing.T) {

	testCases := []struct {
		name   string
		key    string
		values []string

		want bson.D
	}{
		{
			name: "nil values",
			key:  "name",
			want: bson.D{
				bson.E{Key: "name", Value: bson.M{types.All: ([]any)(nil)}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []string{},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []string{"1"},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{"1"}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []string{"1", "2", "3"},
			want:   bson.D{bson.E{Key: "name", Value: bson.M{types.All: []any{"1", "2", "3"}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewBsonBuilder().AllString(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}
