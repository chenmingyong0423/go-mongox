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

package query

import (
	"testing"

	"github.com/chenmingyong0423/go-mongox/types"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_comparisonQueryBuilder(t *testing.T) {
	// eq
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.D{bson.E{Key: types.Eq, Value: "v1"}}}}, BsonBuilder().Eq("k1", "v1").Build())

	// gt
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.D{bson.E{Key: types.Gt, Value: "v1"}}}}, BsonBuilder().Gt("k1", "v1").Build())

	// gte
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.D{bson.E{Key: types.Gte, Value: "v1"}}}}, BsonBuilder().Gte("k1", "v1").Build())

	// lt
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.D{bson.E{Key: types.Lt, Value: "v1"}}}}, BsonBuilder().Lt("k1", "v1").Build())

	// lte
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.D{bson.E{Key: types.Lte, Value: "v1"}}}}, BsonBuilder().Lte("k1", "v1").Build())

	// ne
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.D{bson.E{Key: types.Ne, Value: "v1"}}}}, BsonBuilder().Ne("k1", "v1").Build())
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
			want: bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: ([]any)(nil)}}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []any{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []any{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []any{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []any{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []any{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []any{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().In(tc.key, tc.values...).Build())
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
			want: bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: ([]uint)(nil)}}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []uint{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []uint{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []uint{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []uint{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []uint{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []uint{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().InUint(tc.key, tc.values...).Build())
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
			want: bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: ([]uint8)(nil)}}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []uint8{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []uint8{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []uint8{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []uint8{uint8(1)}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []uint8{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []uint8{uint8(1), uint8(2), uint8(3)}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().InUint8(tc.key, tc.values...).Build())
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
			want: bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: ([]uint16)(nil)}}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []uint16{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []uint16{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []uint16{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []uint16{uint16(1)}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []uint16{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []uint16{uint16(1), uint16(2), uint16(3)}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().InUint16(tc.key, tc.values...).Build())
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
			want: bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: ([]uint32)(nil)}}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []uint32{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []uint32{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []uint32{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []uint32{uint32(1)}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []uint32{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []uint32{uint32(1), uint32(2), uint32(3)}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().InUint32(tc.key, tc.values...).Build())
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
			want: bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: ([]uint64)(nil)}}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []uint64{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []uint64{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []uint64{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []uint64{uint64(1)}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []uint64{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []uint64{uint64(1), uint64(2), uint64(3)}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().InUint64(tc.key, tc.values...).Build())
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
			want: bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: ([]int)(nil)}}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []int{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []int{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []int{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []int{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []int{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []int{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().InInt(tc.key, tc.values...).Build())
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
			want: bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: ([]int8)(nil)}}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []int8{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []int8{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []int8{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []int8{int8(1)}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []int8{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []int8{int8(1), int8(2), int8(3)}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().InInt8(tc.key, tc.values...).Build())
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
			want: bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: ([]int16)(nil)}}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []int16{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []int16{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []int16{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []int16{int16(1)}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []int16{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []int16{int16(1), int16(2), int16(3)}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().InInt16(tc.key, tc.values...).Build())
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
			want: bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: ([]int32)(nil)}}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []int32{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []int32{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []int32{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []int32{int32(1)}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []int32{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []int32{int32(1), int32(2), int32(3)}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().InInt32(tc.key, tc.values...).Build())
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
			want: bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: ([]int64)(nil)}}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []int64{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []int64{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []int64{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []int64{int64(1)}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []int64{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []int64{int64(1), int64(2), int64(3)}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().InInt64(tc.key, tc.values...).Build())
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
			want: bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: ([]float32)(nil)}}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []float32{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []float32{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []float32{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []float32{float32(1)}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []float32{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []float32{float32(1), float32(2), float32(3)}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := BsonBuilder().InFloat32(tc.key, tc.values...).Build()
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
			want: bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: ([]float64)(nil)}}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []float64{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []float64{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []float64{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []float64{float64(1)}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []float64{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []float64{float64(1), float64(2), float64(3)}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := BsonBuilder().InFloat64(tc.key, tc.values...).Build()
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
			want: bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: ([]string)(nil)}}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []string{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []string{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []string{"1"},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []string{"1"}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []string{"1", "2", "3"},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.In, Value: []string{"1", "2", "3"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := BsonBuilder().InString(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
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
			want: bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: ([]any)(nil)}}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []any{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []any{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []any{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []any{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []any{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []any{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().Nin(tc.key, tc.values...).Build())
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
			want: bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: ([]uint)(nil)}}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []uint{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []uint{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []uint{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []uint{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []uint{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []uint{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := BsonBuilder().NinUint(tc.key, tc.values...).Build()
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
			want: bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: ([]uint8)(nil)}}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []uint8{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []uint8{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []uint8{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []uint8{uint8(1)}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []uint8{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []uint8{uint8(1), uint8(2), uint8(3)}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := BsonBuilder().NinUint8(tc.key, tc.values...).Build()
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
			want: bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: ([]uint16)(nil)}}}},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []uint16{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []uint16{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []uint16{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []uint16{uint16(1)}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []uint16{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []uint16{uint16(1), uint16(2), uint16(3)}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := BsonBuilder().NinUint16(tc.key, tc.values...).Build()
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
				bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: ([]uint32)(nil)}}},
			},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []uint32{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []uint32{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []uint32{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []uint32{uint32(1)}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []uint32{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []uint32{uint32(1), uint32(2), uint32(3)}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := BsonBuilder().NinUint32(tc.key, tc.values...).Build()
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
				bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: ([]uint64)(nil)}}},
			},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []uint64{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []uint64{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []uint64{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []uint64{uint64(1)}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []uint64{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []uint64{uint64(1), uint64(2), uint64(3)}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := BsonBuilder().NinUint64(tc.key, tc.values...).Build()
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
				bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: ([]int)(nil)}}},
			},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []int{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []int{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []int{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []int{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []int{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []int{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := BsonBuilder().NinInt(tc.key, tc.values...).Build()
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
				bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: ([]int8)(nil)}}},
			},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []int8{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []int8{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []int8{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []int8{int8(1)}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []int8{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []int8{int8(1), int8(2), int8(3)}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := BsonBuilder().NinInt8(tc.key, tc.values...).Build()
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
				bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: ([]int16)(nil)}}},
			},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []int16{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []int16{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []int16{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []int16{int16(1)}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []int16{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []int16{int16(1), int16(2), int16(3)}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := BsonBuilder().NinInt16(tc.key, tc.values...).Build()
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
				bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: ([]int32)(nil)}}},
			},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []int32{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []int32{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []int32{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []int32{int32(1)}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []int32{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []int32{int32(1), int32(2), int32(3)}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := BsonBuilder().NinInt32(tc.key, tc.values...).Build()
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
				bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: ([]int64)(nil)}}},
			},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []int64{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []int64{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []int64{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []int64{int64(1)}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []int64{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []int64{int64(1), int64(2), int64(3)}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := BsonBuilder().NinInt64(tc.key, tc.values...).Build()
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
				bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: ([]float32)(nil)}}},
			},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []float32{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []float32{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []float32{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []float32{float32(1)}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []float32{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []float32{float32(1), float32(2), float32(3)}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := BsonBuilder().NinFloat32(tc.key, tc.values...).Build()
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
				bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: ([]float64)(nil)}}},
			},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []float64{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []float64{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []float64{1},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []float64{float64(1)}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []float64{1, 2, 3},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []float64{float64(1), float64(2), float64(3)}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := BsonBuilder().NinFloat64(tc.key, tc.values...).Build()
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
				bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: ([]string)(nil)}}},
			},
		},

		{
			name:   "empty values",
			key:    "id",
			values: []string{},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []string{}}}}},
		},
		{
			name:   "one value",
			key:    "id",
			values: []string{"1"},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []string{"1"}}}}},
		},
		{
			name:   "multiple values",
			key:    "id",
			values: []string{"1", "2", "3"},
			want:   bson.D{bson.E{Key: "id", Value: bson.D{bson.E{Key: types.Nin, Value: []string{"1", "2", "3"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := BsonBuilder().NinString(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}
