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
	assert.Equal(t, bson.D{{Key: "age", Value: bson.D{bson.E{Key: "$elemMatch", Value: bson.D{bson.E{Key: "$gt", Value: 1}}}}}}, BsonBuilder().ElemMatch("age", BsonBuilder().Add("$gt", 1).Build()).Build())
}

func TestBsonBuilder_All(t *testing.T) {

	testCases := []struct {
		name   string
		key    string
		values []any

		builder *Builder

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: BsonBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: ([]any)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []any{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []any{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: BsonBuilder(),
			values:  []any{1},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []any{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []any{1, 2, 3},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []any{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: BsonBuilder().Gt("age", 18),
			values:  []any{18, 19, 20},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.Gt, Value: 18}, bson.E{Key: types.All, Value: []any{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.builder.All(tc.key, tc.values...).Build())
		})
	}
}

func TestBsonBuilder_AllUint(t *testing.T) {

	testCases := []struct {
		name   string
		key    string
		values []uint

		builder *Builder

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: BsonBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: ([]uint)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []uint{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []uint{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: BsonBuilder(),
			values:  []uint{1},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []uint{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []uint{1, 2, 3},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []uint{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: BsonBuilder().Gt("age", 18),
			values:  []uint{18, 19, 20},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.Gt, Value: 18}, bson.E{Key: types.All, Value: []uint{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.builder.AllUint(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllUint8(t *testing.T) {

	testCases := []struct {
		name    string
		key     string
		values  []uint8
		builder *Builder

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: BsonBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: ([]uint8)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []uint8{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []uint8{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: BsonBuilder(),
			values:  []uint8{1},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []uint8{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []uint8{1, 2, 3},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []uint8{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: BsonBuilder().Gt("age", 18),
			values:  []uint8{18, 19, 20},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.Gt, Value: 18}, bson.E{Key: types.All, Value: []uint8{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.builder.AllUint8(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllUint16(t *testing.T) {

	testCases := []struct {
		name    string
		key     string
		values  []uint16
		builder *Builder

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: BsonBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: ([]uint16)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []uint16{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []uint16{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: BsonBuilder(),
			values:  []uint16{1},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []uint16{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []uint16{1, 2, 3},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []uint16{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: BsonBuilder().Gt("age", 18),
			values:  []uint16{18, 19, 20},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.Gt, Value: 18}, bson.E{Key: types.All, Value: []uint16{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.builder.AllUint16(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllUint32(t *testing.T) {

	testCases := []struct {
		name    string
		key     string
		builder *Builder
		values  []uint32

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: BsonBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: ([]uint32)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []uint32{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []uint32{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: BsonBuilder(),
			values:  []uint32{1},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []uint32{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []uint32{1, 2, 3},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []uint32{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: BsonBuilder().Gt("age", 18),
			values:  []uint32{18, 19, 20},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.Gt, Value: 18}, bson.E{Key: types.All, Value: []uint32{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.builder.AllUint32(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllUint64(t *testing.T) {

	testCases := []struct {
		name    string
		key     string
		values  []uint64
		builder *Builder

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: BsonBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: ([]uint64)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []uint64{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []uint64{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: BsonBuilder(),
			values:  []uint64{1},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []uint64{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []uint64{1, 2, 3},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []uint64{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: BsonBuilder().Gt("age", 18),
			values:  []uint64{18, 19, 20},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.Gt, Value: 18}, bson.E{Key: types.All, Value: []uint64{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.builder.AllUint64(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllInt(t *testing.T) {

	testCases := []struct {
		name    string
		key     string
		values  []int
		builder *Builder

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: BsonBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: ([]int)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []int{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []int{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: BsonBuilder(),
			values:  []int{1},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []int{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []int{1, 2, 3},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []int{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: BsonBuilder().Gt("age", 18),
			values:  []int{18, 19, 20},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.Gt, Value: 18}, bson.E{Key: types.All, Value: []int{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.builder.AllInt(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllInt8(t *testing.T) {

	testCases := []struct {
		name    string
		key     string
		values  []int8
		builder *Builder

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: BsonBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: ([]int8)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []int8{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []int8{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: BsonBuilder(),
			values:  []int8{1},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []int8{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []int8{1, 2, 3},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []int8{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: BsonBuilder().Gt("age", 18),
			values:  []int8{18, 19, 20},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.Gt, Value: 18}, bson.E{Key: types.All, Value: []int8{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := tc.builder.AllInt8(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllInt16(t *testing.T) {
	testCases := []struct {
		name    string
		key     string
		values  []int16
		builder *Builder

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: BsonBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: ([]int16)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []int16{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []int16{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: BsonBuilder(),
			values:  []int16{1},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []int16{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []int16{1, 2, 3},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []int16{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: BsonBuilder().Gt("age", 18),
			values:  []int16{18, 19, 20},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.Gt, Value: 18}, bson.E{Key: types.All, Value: []int16{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.builder.AllInt16(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllInt32(t *testing.T) {

	testCases := []struct {
		name    string
		key     string
		values  []int32
		builder *Builder
		want    bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: BsonBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: ([]int32)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			values:  []int32{},
			builder: BsonBuilder(),
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []int32{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			values:  []int32{1},
			builder: BsonBuilder(),
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []int32{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			values:  []int32{1, 2, 3},
			builder: BsonBuilder(),
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []int32{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			values:  []int32{18, 19, 20},
			builder: BsonBuilder().Gt("age", 18),
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.Gt, Value: 18}, bson.E{Key: types.All, Value: []int32{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := tc.builder.AllInt32(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllInt64(t *testing.T) {

	testCases := []struct {
		name    string
		key     string
		values  []int64
		builder *Builder

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: BsonBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: ([]int64)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			values:  []int64{},
			builder: BsonBuilder(),
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []int64{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			values:  []int64{1},
			builder: BsonBuilder(),
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []int64{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			values:  []int64{1, 2, 3},
			builder: BsonBuilder(),
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []int64{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			values:  []int64{18, 19, 20},
			builder: BsonBuilder().Gt("age", 18),
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.Gt, Value: 18}, bson.E{Key: types.All, Value: []int64{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.builder.AllInt64(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllFloat32(t *testing.T) {

	testCases := []struct {
		name    string
		key     string
		values  []float32
		builder *Builder

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: BsonBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: ([]float32)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []float32{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []float32{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: BsonBuilder(),
			values:  []float32{1},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []float32{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []float32{1, 2, 3},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []float32{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: BsonBuilder().Gt("age", 18),
			values:  []float32{18, 19, 20},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.Gt, Value: 18}, bson.E{Key: types.All, Value: []float32{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.builder.AllFloat32(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllFloat64(t *testing.T) {

	testCases := []struct {
		name    string
		key     string
		values  []float64
		builder *Builder

		want bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: BsonBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: ([]float64)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []float64{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []float64{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: BsonBuilder(),
			values:  []float64{1},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []float64{1}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []float64{1, 2, 3},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []float64{1, 2, 3}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: BsonBuilder().Gt("age", 18),
			values:  []float64{18, 19, 20},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.Gt, Value: 18}, bson.E{Key: types.All, Value: []float64{18, 19, 20}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.builder.AllFloat64(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestBsonBuilder_AllString(t *testing.T) {

	testCases := []struct {
		name    string
		key     string
		values  []string
		builder *Builder
		want    bson.D
	}{
		{
			name:    "nil values",
			key:     "age",
			builder: BsonBuilder(),
			want: bson.D{
				bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: ([]string)(nil)}}},
			},
		},
		{
			name:    "empty values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []string{},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []string{}}}}},
		},
		{
			name:    "one value",
			key:     "age",
			builder: BsonBuilder(),
			values:  []string{"1"},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []string{"1"}}}}},
		},
		{
			name:    "multiple values",
			key:     "age",
			builder: BsonBuilder(),
			values:  []string{"1", "2", "3"},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.All, Value: []string{"1", "2", "3"}}}}},
		},
		{
			name:    "merge value",
			key:     "age",
			builder: BsonBuilder().Gt("age", 18),
			values:  []string{"18", "19", "20"},
			want:    bson.D{bson.E{Key: "age", Value: bson.D{bson.E{Key: types.Gt, Value: 18}, bson.E{Key: types.All, Value: []string{"18", "19", "20"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.builder.AllString(tc.key, tc.values...).Build()
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_arrayQueryBuilder_Size(t *testing.T) {
	assert.Equal(t, bson.D{{Key: "age", Value: bson.D{bson.E{Key: "$size", Value: 1}}}}, BsonBuilder().Size("age", 1).Build())

	assert.Equal(t, bson.D{{Key: "age", Value: bson.D{bson.E{Key: "$gt", Value: 18}, bson.E{Key: "$size", Value: 1}}}}, BsonBuilder().Gt("age", 18).Size("age", 1).Build())
}
