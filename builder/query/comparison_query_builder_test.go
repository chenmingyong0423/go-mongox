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

	"github.com/chenmingyong0423/go-mongox/types"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_comparisonQueryBuilder(t *testing.T) {
	// eq
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{types.Eq: "v1"}}}, Query().Eq("k1", "v1").Build())

	// gt
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{types.Gt: "v1"}}}, Query().Gt("k1", "v1").Build())

	// gte
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{types.Gte: "v1"}}}, Query().Gte("k1", "v1").Build())

	// lt
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{types.Lt: "v1"}}}, Query().Lt("k1", "v1").Build())

	// lte
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{types.Lte: "v1"}}}, Query().Lte("k1", "v1").Build())

	// ne
	assert.Equal(t, bson.D{bson.E{Key: "k1", Value: bson.M{types.Ne: "v1"}}}, Query().Ne("k1", "v1").Build())
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
			assert.Equal(t, tc.want, Query().In(tc.key, tc.values...).Build())
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
			assert.Equal(t, tc.want, Query().InUint(tc.key, tc.values...).Build())
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
			assert.Equal(t, tc.want, Query().InUint8(tc.key, tc.values...).Build())
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
			assert.Equal(t, tc.want, Query().InUint16(tc.key, tc.values...).Build())
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
			assert.Equal(t, tc.want, Query().InUint32(tc.key, tc.values...).Build())
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
			assert.Equal(t, tc.want, Query().InUint64(tc.key, tc.values...).Build())
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
			assert.Equal(t, tc.want, Query().InInt(tc.key, tc.values...).Build())
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
			assert.Equal(t, tc.want, Query().InInt8(tc.key, tc.values...).Build())
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
			assert.Equal(t, tc.want, Query().InInt16(tc.key, tc.values...).Build())
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
			assert.Equal(t, tc.want, Query().InInt32(tc.key, tc.values...).Build())
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
			assert.Equal(t, tc.want, Query().InInt64(tc.key, tc.values...).Build())
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
			got := Query().InFloat32(tc.key, tc.values...).Build()
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
			got := Query().InFloat64(tc.key, tc.values...).Build()
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
			got := Query().InString(tc.key, tc.values...).Build()
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
			assert.Equal(t, tc.want, Query().Nin(tc.key, tc.values...).Build())
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
			got := Query().NinUint(tc.key, tc.values...).Build()
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
			got := Query().NinUint8(tc.key, tc.values...).Build()
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
			got := Query().NinUint16(tc.key, tc.values...).Build()
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
			got := Query().NinUint32(tc.key, tc.values...).Build()
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

			got := Query().NinUint64(tc.key, tc.values...).Build()
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

			got := Query().NinInt(tc.key, tc.values...).Build()
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

			got := Query().NinInt8(tc.key, tc.values...).Build()
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

			got := Query().NinInt16(tc.key, tc.values...).Build()
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

			got := Query().NinInt32(tc.key, tc.values...).Build()
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

			got := Query().NinInt64(tc.key, tc.values...).Build()
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

			got := Query().NinFloat32(tc.key, tc.values...).Build()
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

			got := Query().NinFloat64(tc.key, tc.values...).Build()
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

			got := Query().NinString(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}
