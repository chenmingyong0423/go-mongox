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

func Test_arrayQueryBuilder_ElemMatch(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "name", Value: bson.D{bson.E{Key: "$elemMatch", Value: bson.D{bson.E{Key: "$gt", Value: 1}}}}}}, BsonBuilder().ElemMatch("name", BsonBuilder().Add(types.KV("$gt", 1)).Build()).Build())
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
				bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: ([]any)(nil)}}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []any{},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []any{}}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []any{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []any{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []any{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []any{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, BsonBuilder().All(tc.key, tc.values...).Build())
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
				bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: ([]uint)(nil)}}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []uint{},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []uint{}}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []uint{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []uint{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []uint{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []uint{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := BsonBuilder().AllUint(tc.key, tc.values...).Build()
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
				bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: ([]uint8)(nil)}}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []uint8{},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []uint8{}}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []uint8{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []uint8{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []uint8{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []uint8{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := BsonBuilder().AllUint8(tc.key, tc.values...).Build()
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
				bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: ([]uint16)(nil)}}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []uint16{},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []uint16{}}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []uint16{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []uint16{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []uint16{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []uint16{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := BsonBuilder().AllUint16(tc.key, tc.values...).Build()
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
				bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: ([]uint32)(nil)}}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []uint32{},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []uint32{}}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []uint32{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []uint32{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []uint32{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []uint32{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := BsonBuilder().AllUint32(tc.key, tc.values...).Build()
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
				bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: ([]uint64)(nil)}}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []uint64{},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []uint64{}}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []uint64{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []uint64{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []uint64{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []uint64{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := BsonBuilder().AllUint64(tc.key, tc.values...).Build()
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
				bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: ([]int)(nil)}}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []int{},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []int{}}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []int{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []int{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []int{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []int{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := BsonBuilder().AllInt(tc.key, tc.values...).Build()
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
				bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: ([]int8)(nil)}}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []int8{},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []int8{}}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []int8{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []int8{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []int8{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []int8{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := BsonBuilder().AllInt8(tc.key, tc.values...).Build()
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
				bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: ([]int16)(nil)}}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []int16{},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []int16{}}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []int16{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []int16{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []int16{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []int16{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := BsonBuilder().AllInt16(tc.key, tc.values...).Build()
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
				bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: ([]int32)(nil)}}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []int32{},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []int32{}}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []int32{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []int32{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []int32{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []int32{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := BsonBuilder().AllInt32(tc.key, tc.values...).Build()
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
				bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: ([]int64)(nil)}}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []int64{},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []int64{}}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []int64{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []int64{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []int64{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []int64{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := BsonBuilder().AllInt64(tc.key, tc.values...).Build()
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
				bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: ([]float32)(nil)}}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []float32{},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []float32{}}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []float32{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []float32{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []float32{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []float32{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := BsonBuilder().AllFloat32(tc.key, tc.values...).Build()
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
				bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: ([]float64)(nil)}}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []float64{},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []float64{}}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []float64{1},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []float64{1}}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []float64{1, 2, 3},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []float64{1, 2, 3}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := BsonBuilder().AllFloat64(tc.key, tc.values...).Build()
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
				bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: ([]string)(nil)}}},
			},
		},
		{
			name:   "empty values",
			key:    "name",
			values: []string{},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []string{}}}}},
		},
		{
			name:   "one value",
			key:    "name",
			values: []string{"1"},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []string{"1"}}}}},
		},
		{
			name:   "multiple values",
			key:    "name",
			values: []string{"1", "2", "3"},
			want:   bson.D{bson.E{Key: "name", Value: bson.D{bson.E{Key: types.All, Value: []string{"1", "2", "3"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := BsonBuilder().AllString(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_arrayQueryBuilder_Size(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "name", Value: bson.D{bson.E{Key: "$size", Value: 1}}}}, BsonBuilder().Size("name", 1).Build())
}
