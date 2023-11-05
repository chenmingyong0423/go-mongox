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

package aggregation

import (
	"testing"

	"github.com/chenmingyong0423/go-mongox/pkg/utils"
	"github.com/chenmingyong0423/go-mongox/types"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestStageBuilder_AddFields(t *testing.T) {
	testCases := []struct {
		name      string
		keyValues []any
		want      mongo.Pipeline
	}{
		{
			name:      "nil keyValues",
			keyValues: nil,
			want:      mongo.Pipeline{},
		},
		{
			name:      "odd keyValues",
			keyValues: []any{"totalHomework", BsonBuilder().Sum("$homework").Build(), "totalQuiz"},
			want:      mongo.Pipeline{},
		},
		{
			name:      "even keyValues",
			keyValues: []any{"totalHomework", BsonBuilder().Sum("$homework").Build(), "totalQuiz", BsonBuilder().Sum("$quiz").Build()},
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$addFields", Value: bson.D{
					bson.E{Key: "totalHomework", Value: bson.D{bson.E{Key: "$sum", Value: "$homework"}}},
					bson.E{Key: "totalQuiz", Value: bson.D{bson.E{Key: "$sum", Value: "$quiz"}}},
				}}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, StageBsonBuilder().AddFields(tc.keyValues...).Build())
		})
	}
}

func TestStageBuilder_AddFieldsForMap(t *testing.T) {
	testCases := []struct {
		name      string
		keyValues map[string]any
		want      mongo.Pipeline
	}{
		{
			name:      "nil keyValues",
			keyValues: nil,
			want:      mongo.Pipeline{},
		},
		{
			name:      "empty keyValues",
			keyValues: map[string]any{},
			want:      mongo.Pipeline{bson.D{bson.E{Key: "$addFields", Value: bson.D{}}}},
		},
		{
			name: "not nil keyValues",
			keyValues: map[string]any{
				"totalHomework": BsonBuilder().Sum("$homework").Build(),
				"totalQuiz":     BsonBuilder().Sum("$quiz").Build(),
			},
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$addFields", Value: bson.D{
					bson.E{Key: "totalHomework", Value: bson.D{bson.E{Key: "$sum", Value: "$homework"}}},
					bson.E{Key: "totalQuiz", Value: bson.D{bson.E{Key: "$sum", Value: "$quiz"}}},
				}}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, utils.EqualPipelineElements(tc.want, StageBsonBuilder().AddFieldsForMap(tc.keyValues).Build()))
		})
	}
}

func TestStageBuilder_Set(t *testing.T) {
	testCases := []struct {
		name      string
		keyValues []any
		want      mongo.Pipeline
	}{
		{
			name:      "nil keyValues",
			keyValues: nil,
			want:      mongo.Pipeline{},
		},
		{
			name:      "odd keyValues",
			keyValues: []any{"totalHomework", BsonBuilder().Sum("$homework").Build(), "totalQuiz"},
			want:      mongo.Pipeline{},
		},
		{
			name:      "even keyValues",
			keyValues: []any{"totalHomework", BsonBuilder().Sum("$homework").Build(), "totalQuiz", BsonBuilder().Sum("$quiz").Build()},
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$set", Value: bson.D{
					bson.E{Key: "totalHomework", Value: bson.D{bson.E{Key: "$sum", Value: "$homework"}}},
					bson.E{Key: "totalQuiz", Value: bson.D{bson.E{Key: "$sum", Value: "$quiz"}}},
				}}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, StageBsonBuilder().Set(tc.keyValues...).Build())
		})
	}
}

func TestStageBuilder_SetForMap(t *testing.T) {
	testCases := []struct {
		name      string
		keyValues map[string]any
		want      mongo.Pipeline
	}{
		{
			name:      "nil keyValues",
			keyValues: nil,
			want:      mongo.Pipeline{},
		},
		{
			name:      "empty keyValues",
			keyValues: map[string]any{},
			want:      mongo.Pipeline{bson.D{bson.E{Key: "$set", Value: bson.D{}}}},
		},
		{
			name: "not nil keyValues",
			keyValues: map[string]any{
				"totalHomework": BsonBuilder().Sum("$homework").Build(),
				"totalQuiz":     BsonBuilder().Sum("$quiz").Build(),
			},
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$set", Value: bson.D{
					bson.E{Key: "totalHomework", Value: bson.D{bson.E{Key: "$sum", Value: "$homework"}}},
					bson.E{Key: "totalQuiz", Value: bson.D{bson.E{Key: "$sum", Value: "$quiz"}}},
				}}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, utils.EqualPipelineElements(tc.want, StageBsonBuilder().SetForMap(tc.keyValues).Build()))
		})
	}
}

func TestStageBuilder_Bucket(t *testing.T) {
	testCases := []struct {
		name       string
		groupBy    any
		opt        *types.BucketOptions
		boundaries []any
		want       mongo.Pipeline
	}{
		{
			name:       "defaultKey and output are nil",
			groupBy:    "$year_born",
			opt:        nil,
			boundaries: []any{1840, 1850, 1860, 1870, 1880},
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$bucket", Value: bson.D{
					bson.E{Key: "groupBy", Value: "$year_born"},
					bson.E{Key: "boundaries", Value: []any{1840, 1850, 1860, 1870, 1880}},
				}}},
			},
		},
		{
			name:    "output is nil",
			groupBy: "$year_born",
			opt: &types.BucketOptions{
				DefaultKey: "Other",
				Output:     nil,
			},
			boundaries: []any{1840, 1850, 1860, 1870, 1880},
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$bucket", Value: bson.D{
					bson.E{Key: "groupBy", Value: "$year_born"},
					bson.E{Key: "boundaries", Value: []any{1840, 1850, 1860, 1870, 1880}},
					bson.E{Key: "default", Value: "Other"},
				}}},
			},
		},
		{
			name:    "defaultKey is empty",
			groupBy: "$year_born",
			opt: &types.BucketOptions{
				DefaultKey: nil,
				Output:     BsonBuilder().AddKeyValues("count", BsonBuilder().Sum(1).Build()).Build(),
			},
			boundaries: []any{1840, 1850, 1860, 1870, 1880},
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$bucket", Value: bson.D{
					bson.E{Key: "groupBy", Value: "$year_born"},
					bson.E{Key: "boundaries", Value: []any{1840, 1850, 1860, 1870, 1880}},
					bson.E{Key: "output", Value: bson.D{bson.E{Key: "count", Value: bson.D{bson.E{Key: "$sum", Value: 1}}}}},
				}}},
			},
		},
		{
			name:    "boundaries is nil",
			groupBy: "$year_born",
			opt: &types.BucketOptions{
				DefaultKey: "Other",
				Output:     BsonBuilder().AddKeyValues("count", BsonBuilder().Sum(1).Build()).Build(),
			},
			boundaries: nil,
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$bucket", Value: bson.D{
					bson.E{Key: "groupBy", Value: "$year_born"},
					bson.E{Key: "boundaries", Value: ([]any)(nil)},
					bson.E{Key: "default", Value: "Other"},
					bson.E{Key: "output", Value: bson.D{bson.E{Key: "count", Value: bson.D{bson.E{Key: "$sum", Value: 1}}}}},
				}}},
			},
		},
		{
			name:    "all not nil",
			groupBy: "$year_born",
			opt: &types.BucketOptions{
				DefaultKey: "Other",
				Output: BsonBuilder().AddKeyValues(
					"count", BsonBuilder().Sum(1).Build(),
					"artists", BsonBuilder().Push(BsonBuilder().AddKeyValues("name", BsonBuilder().Contact("$first_name", " ", "$last_name").Build(), "year_born", "$year_born").Build()).Build(),
				).Build(),
			},
			boundaries: []any{1840, 1850, 1860, 1870, 1880},
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$bucket", Value: bson.D{
					bson.E{Key: "groupBy", Value: "$year_born"},
					bson.E{Key: "boundaries", Value: []any{1840, 1850, 1860, 1870, 1880}},
					bson.E{Key: "default", Value: "Other"},
					bson.E{Key: "output", Value: bson.D{
						bson.E{Key: "count", Value: bson.D{bson.E{Key: "$sum", Value: 1}}},
						bson.E{Key: "artists", Value: bson.D{
							bson.E{Key: "$push", Value: bson.D{
								bson.E{Key: "name", Value: bson.D{bson.E{Key: "$concat", Value: []any{"$first_name", " ", "$last_name"}}}},
								bson.E{Key: "year_born", Value: "$year_born"},
							}},
						}},
					}},
				}}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, StageBsonBuilder().Bucket(tc.groupBy, tc.boundaries, tc.opt).Build())
		})
	}
}

func TestStageBuilder_BucketAuto(t *testing.T) {
	testCases := []struct {
		name    string
		groupBy any
		opt     *types.BucketAutoOptions
		buckets int
		want    mongo.Pipeline
	}{
		{
			name:    "output and granularity are nil",
			groupBy: "$price",
			opt:     nil,
			buckets: 4,
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$bucketAuto", Value: bson.D{
					bson.E{Key: "groupBy", Value: "$price"},
					bson.E{Key: "buckets", Value: 4},
				}}},
			},
		},
		{
			name:    "output is nil",
			groupBy: "$price",
			opt: &types.BucketAutoOptions{
				Granularity: "R5",
			},
			buckets: 4,
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$bucketAuto", Value: bson.D{
					bson.E{Key: "groupBy", Value: "$price"},
					bson.E{Key: "buckets", Value: 4},
					bson.E{Key: "granularity", Value: "R5"},
				}}},
			},
		},
		{
			name:    "granularity is empty",
			groupBy: "$price",
			opt: &types.BucketAutoOptions{
				Output: BsonBuilder().AddKeyValues("count", BsonBuilder().Sum(1).Build()).Build(),
			},
			buckets: 4,
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$bucketAuto", Value: bson.D{
					bson.E{Key: "groupBy", Value: "$price"},
					bson.E{Key: "buckets", Value: 4},
					bson.E{Key: "output", Value: bson.D{bson.E{Key: "count", Value: bson.D{bson.E{Key: "$sum", Value: 1}}}}},
				}}},
			},
		},
		{
			name:    "normal",
			groupBy: "$price",
			opt: &types.BucketAutoOptions{
				Output:      BsonBuilder().AddKeyValues("count", BsonBuilder().Sum(1).Build()).Build(),
				Granularity: "R5",
			},
			buckets: 4,
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$bucketAuto", Value: bson.D{
					bson.E{Key: "groupBy", Value: "$price"},
					bson.E{Key: "buckets", Value: 4},
					bson.E{Key: "output", Value: bson.D{bson.E{Key: "count", Value: bson.D{bson.E{Key: "$sum", Value: 1}}}}},
					bson.E{Key: "granularity", Value: "R5"},
				}}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, StageBsonBuilder().BucketAuto(tc.groupBy, tc.buckets, tc.opt).Build())
		})
	}
}

func TestStageBuilder_Match(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		want       mongo.Pipeline
	}{
		{
			name:       "expression is nil",
			expression: nil,
			want:       mongo.Pipeline{bson.D{bson.E{Key: "$match", Value: nil}}},
		},
		{
			name:       "expression is not nil",
			expression: BsonBuilder().AddKeyValues("author", "dave").Build(),
			want:       mongo.Pipeline{bson.D{bson.E{Key: "$match", Value: bson.D{bson.E{Key: "author", Value: "dave"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, StageBsonBuilder().Match(tc.expression).Build())
		})
	}
}

func TestStageBuilder_Group(t *testing.T) {
	testCases := []struct {
		name         string
		id           any
		accumulators []any
		want         mongo.Pipeline
	}{
		{
			name:         "id and accumulators are nil",
			id:           nil,
			accumulators: nil,
			want:         mongo.Pipeline{bson.D{bson.E{Key: "$group", Value: bson.D{bson.E{Key: "_id", Value: nil}}}}},
		},
		{
			name: "id is nil",
			id:   nil,
			accumulators: []any{
				"totalSaleAmount", BsonBuilder().Sum(BsonBuilder().Multiply("$price", "$quantity").Build()).Build(),
				"averageQuantity", BsonBuilder().Avg("$quantity").Build(),
				"count", BsonBuilder().Sum(1).Build(),
			},
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$group", Value: bson.D{
					bson.E{Key: "_id", Value: nil},
					bson.E{Key: "totalSaleAmount", Value: bson.D{bson.E{Key: "$sum", Value: bson.D{bson.E{Key: "$multiply", Value: []any{"$price", "$quantity"}}}}}},
					bson.E{Key: "averageQuantity", Value: bson.D{bson.E{Key: "$avg", Value: "$quantity"}}},
					bson.E{Key: "count", Value: bson.D{bson.E{Key: "$sum", Value: 1}}},
				}}},
			},
		},
		{
			name: "accumulators is nil",
			id: BsonBuilder().DateToString("$date", &types.DateToStringOptions{
				Format:   "%Y-%m-%d",
				Timezone: "",
				OnNull:   nil,
			}).Build(),
			accumulators: nil,
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$group", Value: bson.D{
					bson.E{Key: "_id", Value: bson.D{bson.E{Key: "$dateToString", Value: bson.D{bson.E{Key: "date", Value: "$date"}, bson.E{Key: "format", Value: "%Y-%m-%d"}}}}},
				}}},
			},
		},
		{
			name: "id and accumulators are not nil",
			id: BsonBuilder().DateToString("$date", &types.DateToStringOptions{
				Format:   "%Y-%m-%d",
				Timezone: "",
				OnNull:   nil,
			}).Build(),
			accumulators: []any{
				"totalSaleAmount", BsonBuilder().Sum(BsonBuilder().Multiply("$price", "$quantity").Build()).Build(),
				"averageQuantity", BsonBuilder().Avg("$quantity").Build(),
				"count", BsonBuilder().Sum(1).Build(),
			},
			want: mongo.Pipeline{
				bson.D{bson.E{Key: "$group", Value: bson.D{
					bson.E{Key: "_id", Value: bson.D{bson.E{Key: "$dateToString", Value: bson.D{bson.E{Key: "date", Value: "$date"}, bson.E{Key: "format", Value: "%Y-%m-%d"}}}}},
					bson.E{Key: "totalSaleAmount", Value: bson.D{bson.E{Key: "$sum", Value: bson.D{bson.E{Key: "$multiply", Value: []any{"$price", "$quantity"}}}}}},
					bson.E{Key: "averageQuantity", Value: bson.D{bson.E{Key: "$avg", Value: "$quantity"}}},
					bson.E{Key: "count", Value: bson.D{bson.E{Key: "$sum", Value: 1}}},
				}}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, StageBsonBuilder().Group(tc.id, tc.accumulators...).Build())
		})
	}
}

func TestStageBuilder_GroupMap(t *testing.T) {
	testCases := []struct {
		name         string
		id           any
		accumulators map[string]map[string]any
		want         mongo.Pipeline
	}{
		{
			name:         "id and accumulators are nil",
			id:           nil,
			accumulators: nil,
			want:         mongo.Pipeline{bson.D{bson.E{Key: "$group", Value: bson.D{bson.E{Key: "_id", Value: nil}}}}},
		},
		{
			name:         "accumulators are nil",
			id:           "$author",
			accumulators: nil,
			want:         mongo.Pipeline{bson.D{bson.E{Key: "$group", Value: bson.D{bson.E{Key: "_id", Value: "$author"}}}}},
		},
		{
			name: "string of id",
			id:   "$author",
			accumulators: map[string]map[string]any{
				"totalSaleAmount": {types.AggregationSum: BsonBuilder().Multiply("$price", "$quantity").Build()},
			},
			want: mongo.Pipeline{bson.D{bson.E{Key: "$group", Value: bson.D{
				bson.E{Key: "_id", Value: "$author"},
				bson.E{Key: "totalSaleAmount", Value: bson.D{bson.E{Key: "$sum", Value: bson.D{bson.E{Key: "$multiply", Value: []any{"$price", "$quantity"}}}}}},
			}}}},
		},
		{
			name: "bsonD of id",
			id:   BsonBuilder().AddKeyValues("x", 1, "y", 1).Build(),
			accumulators: map[string]map[string]any{
				"totalSaleAmount": {types.AggregationSum: BsonBuilder().Multiply("$price", "$quantity").Build()},
			},
			want: mongo.Pipeline{bson.D{bson.E{Key: "$group", Value: bson.D{bson.E{Key: "_id", Value: bson.D{
				bson.E{Key: "x", Value: 1},
				bson.E{Key: "y", Value: 1},
			}},
				bson.E{Key: "totalSaleAmount", Value: bson.D{bson.E{Key: "$sum", Value: bson.D{bson.E{Key: "$multiply", Value: []any{"$price", "$quantity"}}}}}},
			}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, StageBsonBuilder().GroupMap(tc.id, tc.accumulators).Build())
		})
	}
}

func TestStageBuilder_Sort(t *testing.T) {
	testCases := []struct {
		name      string
		keyValues []any
		want      mongo.Pipeline
	}{
		{
			name:      "empty",
			keyValues: nil,
			want:      mongo.Pipeline{},
		},
		{
			name:      "odd",
			keyValues: []any{"name", 1, "age"},
			want:      mongo.Pipeline{},
		},
		{
			// { $sort : { name : 1, age: -1 } }
			name:      "even",
			keyValues: []any{"name", 1, "age", -1},
			want:      mongo.Pipeline{bson.D{bson.E{Key: "$sort", Value: bson.D{bson.E{Key: "name", Value: 1}, bson.E{Key: "age", Value: -1}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, StageBsonBuilder().Sort(tc.keyValues...).Build())
		})
	}
}

func TestStageBuilder_SortMap(t *testing.T) {
	testCases := []struct {
		name      string
		keyValues map[string]any
		want      mongo.Pipeline
	}{
		{
			name:      "empty",
			keyValues: nil,
			want:      mongo.Pipeline{},
		},
		{
			name:      "empty",
			keyValues: map[string]any{},
			want:      mongo.Pipeline{bson.D{bson.E{Key: "$sort", Value: bson.D{}}}},
		},
		{
			name: "not empty",
			keyValues: map[string]any{
				"name": 1,
				"age":  -1,
			},
			want: mongo.Pipeline{bson.D{bson.E{Key: "$sort", Value: bson.D{bson.E{Key: "name", Value: 1}, bson.E{Key: "age", Value: -1}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, StageBsonBuilder().SortMap(tc.keyValues).Build())
		})
	}
}

func TestStageBuilder_Project(t *testing.T) {
	testCases := []struct {
		name      string
		keyValues []any
		want      mongo.Pipeline
	}{
		{
			name:      "nil keyValues",
			keyValues: nil,
			want:      mongo.Pipeline{},
		},
		{
			name:      "odd keyValues",
			keyValues: []any{"_id", 0, "title", 1, "author"},
			want:      mongo.Pipeline{},
		},
		{
			name:      "even keyValues",
			keyValues: []any{"_id", 0, "title", 1, "author", 1},
			want: mongo.Pipeline{bson.D{bson.E{Key: "$project", Value: bson.D{
				bson.E{Key: "_id", Value: 0},
				bson.E{Key: "title", Value: 1},
				bson.E{Key: "author", Value: 1},
			}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, StageBsonBuilder().Project(tc.keyValues...).Build())
		})
	}
}

func TestStageBuilder_ProjectMap(t *testing.T) {
	testCases := []struct {
		name      string
		keyValues map[string]any
		want      mongo.Pipeline
	}{
		{
			name:      "nil keyValues",
			keyValues: nil,
			want:      mongo.Pipeline{},
		},
		{
			name:      "empty keyValues",
			keyValues: map[string]any{},
			want:      mongo.Pipeline{bson.D{bson.E{Key: "$project", Value: bson.D{}}}},
		},
		{
			name: "not nil keyValues",
			keyValues: map[string]any{
				"_id":    0,
				"title":  1,
				"author": 1,
			},
			want: mongo.Pipeline{bson.D{bson.E{Key: "$project", Value: bson.D{
				bson.E{Key: "_id", Value: 0},
				bson.E{Key: "title", Value: 1},
				bson.E{Key: "author", Value: 1},
			}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.True(t, utils.EqualPipelineElements(tc.want, StageBsonBuilder().ProjectMap(tc.keyValues).Build()))
		})
	}
}

func TestStageBuilder_Limit(t *testing.T) {
	assert.Equal(t, mongo.Pipeline{bson.D{bson.E{Key: "$limit", Value: int64(10)}}}, StageBsonBuilder().Limit(10).Build())
}

func TestStageBuilder_Skip(t *testing.T) {
	assert.Equal(t, mongo.Pipeline{bson.D{bson.E{Key: "$skip", Value: int64(10)}}}, StageBsonBuilder().Skip(10).Build())
}

func TestStageBuilder_Unwind(t *testing.T) {
	testCases := []struct {
		name string
		path string
		opt  *types.UnWindOptions
		want mongo.Pipeline
	}{
		{
			name: "opt is nil",
			path: "$sizes",
			opt:  nil,
			want: mongo.Pipeline{bson.D{bson.E{Key: "$unwind", Value: "$sizes"}}},
		},
		{
			name: "opt is not nil and includeArrayIndex is not empty",
			path: "$sizes",
			opt: &types.UnWindOptions{
				IncludeArrayIndex: "arrayIndex",
			},
			want: mongo.Pipeline{bson.D{bson.E{Key: "$unwind", Value: bson.D{
				bson.E{Key: "path", Value: "$sizes"},
				bson.E{Key: "includeArrayIndex", Value: "arrayIndex"},
			},
			}}},
		},
		{
			name: "opt is not nil and preserveNullAndEmptyArrays is true",
			path: "$sizes",
			opt: &types.UnWindOptions{
				PreserveNullAndEmptyArrays: true,
			},
			want: mongo.Pipeline{bson.D{bson.E{Key: "$unwind", Value: bson.D{
				bson.E{Key: "path", Value: "$sizes"},
				bson.E{Key: "preserveNullAndEmptyArrays", Value: true},
			}}}},
		},
		{
			name: "opt is not nil and includeArrayIndex is not empty and preserveNullAndEmptyArrays is true",
			path: "$sizes",
			opt: &types.UnWindOptions{
				IncludeArrayIndex:          "arrayIndex",
				PreserveNullAndEmptyArrays: true,
			},
			want: mongo.Pipeline{bson.D{bson.E{Key: "$unwind", Value: bson.D{
				bson.E{Key: "path", Value: "$sizes"},
				bson.E{Key: "includeArrayIndex", Value: "arrayIndex"},
				bson.E{Key: "preserveNullAndEmptyArrays", Value: true},
			}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, StageBsonBuilder().Unwind(tc.path, tc.opt).Build())
		})
	}
}

func TestStageBuilder_ReplaceWith(t *testing.T) {
	testCases := []struct {
		name                string
		replacementDocument any
		want                mongo.Pipeline
	}{
		{
			name:                "nil replacementDocument",
			replacementDocument: nil,
			want:                mongo.Pipeline{bson.D{bson.E{Key: "$replaceWith", Value: nil}}},
		},
		{
			name:                "replacementDocument of string",
			replacementDocument: "$name",
			want:                mongo.Pipeline{bson.D{bson.E{Key: "$replaceWith", Value: "$name"}}},
		},
		{
			name:                "replacementDocument of bson.D",
			replacementDocument: BsonBuilder().ArrayToObject("$items").Build(),
			want:                mongo.Pipeline{bson.D{bson.E{Key: "$replaceWith", Value: bson.D{bson.E{Key: "$arrayToObject", Value: "$items"}}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, StageBsonBuilder().ReplaceWith(tc.replacementDocument).Build())
		})
	}
}

func TestStageBuilder_Facet(t *testing.T) {
	testCases := []struct {
		name   string
		facets []any
		want   mongo.Pipeline
	}{
		{
			name:   "nil facets",
			facets: nil,
			want:   mongo.Pipeline{},
		},
		{
			name:   "empty facets",
			facets: []any{},
			want:   mongo.Pipeline{bson.D{bson.E{Key: "$facet", Value: bson.D{}}}},
		},
		{
			// [
			//  {
			//    $facet: {
			//      "categorizedByTags": [
			//        { $unwind: "$tags" },
			//        { $sortByCount: "$tags" }
			//      ],
			//      "categorizedByPrice": [
			//        // Filter out documents without a price e.g., _id: 7
			//        { $match: { price: { $exists: 1 } } },
			//        {
			//          $bucket: {
			//            groupBy: "$price",
			//            boundaries: [  0, 150, 200, 300, 400 ],
			//            default: "Other",
			//            output: {
			//              "count": { $sum: 1 },
			//              "titles": { $push: "$title" }
			//            }
			//          }
			//        }
			//      ],
			//      "categorizedByYears(Auto)": [
			//        {
			//          $bucketAuto: {
			//            groupBy: "$year",
			//            buckets: 4
			//          }
			//        }
			//      ]
			//    }
			//  }
			//]
			name: "replacementDocument of bson.D",
			facets: []any{
				"categorizedByTags", StageBsonBuilder().Unwind("$tags", nil).SortByCount("$tags").Build(),

				"categorizedByPrice", StageBsonBuilder().Match(BsonBuilder().AddKeyValues("price", BsonBuilder().AddKeyValues("$exists", 1).Build()).Build()).Bucket("$price", []any{0, 150, 200, 300, 400}, &types.BucketOptions{
					DefaultKey: "Other",
					Output:     BsonBuilder().AddKeyValues("count", BsonBuilder().Sum(1).Build(), "titles", BsonBuilder().Push("$title").Build()).Build(),
				}).Build(),

				"categorizedByYears(Auto)", StageBsonBuilder().BucketAuto("$year", 4, nil).Build(),
			},
			want: mongo.Pipeline{bson.D{bson.E{Key: "$facet", Value: bson.D{
				bson.E{Key: "categorizedByTags", Value: mongo.Pipeline{
					bson.D{bson.E{Key: "$unwind", Value: "$tags"}},
					bson.D{bson.E{Key: "$sortByCount", Value: "$tags"}},
				}},
				bson.E{Key: "categorizedByPrice", Value: mongo.Pipeline{
					bson.D{bson.E{Key: "$match", Value: bson.D{bson.E{Key: "price", Value: bson.D{bson.E{Key: "$exists", Value: 1}}}}}},
					bson.D{bson.E{Key: "$bucket", Value: bson.D{
						bson.E{Key: "groupBy", Value: "$price"},
						bson.E{Key: "boundaries", Value: []any{0, 150, 200, 300, 400}},
						bson.E{Key: "default", Value: "Other"},
						bson.E{Key: "output", Value: bson.D{
							bson.E{Key: "count", Value: bson.D{bson.E{Key: "$sum", Value: 1}}},
							bson.E{Key: "titles", Value: bson.D{bson.E{Key: "$push", Value: "$title"}}},
						}},
					},
					}},
				}},
				bson.E{Key: "categorizedByYears(Auto)", Value: mongo.Pipeline{
					bson.D{bson.E{Key: "$bucketAuto", Value: bson.D{
						bson.E{Key: "groupBy", Value: "$year"},
						bson.E{Key: "buckets", Value: 4},
					}}},
				}},
			}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, StageBsonBuilder().Facet(tc.facets...).Build())
		})
	}
}

func TestStageBuilder_SortByCount(t *testing.T) {
	testCases := []struct {
		name       string
		expression any
		want       mongo.Pipeline
	}{
		{
			name:       "nil expression",
			expression: nil,
			want:       mongo.Pipeline{bson.D{bson.E{Key: "$sortByCount", Value: nil}}},
		},
		{
			name:       "expression of string",
			expression: "$tags",
			want:       mongo.Pipeline{bson.D{bson.E{Key: "$sortByCount", Value: "$tags"}}},
		},
		// { $sortByCount: { lname: "$employee.last", fname: "$employee.first" } }
		{
			name: "expression of bson.D",
			expression: BsonBuilder().AddKeyValues(
				"lname", "$employee.last",
				"fname", "$employee.first",
			).Build(),
			want: mongo.Pipeline{bson.D{bson.E{Key: "$sortByCount", Value: bson.D{
				bson.E{Key: "lname", Value: "$employee.last"},
				bson.E{Key: "fname", Value: "$employee.first"},
			}}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, StageBsonBuilder().SortByCount(tc.expression).Build())
		})
	}
}

func TestStageBuilder_Count(t *testing.T) {
	assert.Equal(t, mongo.Pipeline{bson.D{bson.E{Key: "$count", Value: "passing_scores"}}}, StageBsonBuilder().Count("passing_scores").Build())
}
