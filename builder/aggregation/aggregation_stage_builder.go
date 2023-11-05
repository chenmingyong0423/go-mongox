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
	"github.com/chenmingyong0423/go-mongox/converter"
	"github.com/chenmingyong0423/go-mongox/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StageBuilder struct {
	pipeline mongo.Pipeline
}

func StageBsonBuilder() *StageBuilder {
	return &StageBuilder{pipeline: make([]bson.D, 0, 4)}
}

func (b *StageBuilder) AddFields(keyValues ...any) *StageBuilder {
	if keyValues != nil && len(keyValues)%2 == 0 {
		d := bson.D{}
		for i := 0; i < len(keyValues); i += 2 {
			k, ok := keyValues[i].(string)
			if !ok {
				continue
			}
			d = append(d, bson.E{Key: k, Value: keyValues[i+1]})
		}
		b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageAddFields, Value: d}})
	}
	return b
}

func (b *StageBuilder) AddFieldsForMap(keyValues map[string]any) *StageBuilder {
	if keyValues != nil {
		d := bson.D{}
		for k, v := range keyValues {
			d = append(d, bson.E{Key: k, Value: v})
		}
		b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageAddFields, Value: d}})
	}
	return b
}

func (b *StageBuilder) Set(keyValues ...any) *StageBuilder {
	if keyValues != nil && len(keyValues)%2 == 0 {
		d := bson.D{}
		for i := 0; i < len(keyValues); i += 2 {
			k, ok := keyValues[i].(string)
			if !ok {
				continue
			}
			d = append(d, bson.E{Key: k, Value: keyValues[i+1]})
		}
		b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageSet, Value: d}})
	}
	return b
}

func (b *StageBuilder) SetForMap(keyValues map[string]any) *StageBuilder {
	if keyValues != nil {
		d := bson.D{}
		for k, v := range keyValues {
			d = append(d, bson.E{Key: k, Value: v})
		}
		b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageSet, Value: d}})
	}
	return b
}

func (b *StageBuilder) Bucket(groupBy any, boundaries []any, opt *types.BucketOptions) *StageBuilder {
	d := bson.D{
		bson.E{Key: types.GroupBy, Value: groupBy},
		bson.E{Key: types.Boundaries, Value: boundaries},
	}
	if opt != nil {
		if opt.DefaultKey != nil {
			d = append(d, bson.E{Key: types.Default, Value: opt.DefaultKey})
		}
		if opt.Output != nil {
			d = append(d, bson.E{Key: types.Output, Value: opt.Output})
		}
	}
	b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageBucket, Value: d}})
	return b
}

func (b *StageBuilder) BucketAuto(groupBy any, buckets int, opt *types.BucketAutoOptions) *StageBuilder {
	d := bson.D{
		bson.E{Key: types.GroupBy, Value: groupBy},
		bson.E{Key: types.Buckets, Value: buckets},
	}
	if opt != nil {
		if opt.Output != nil {
			d = append(d, bson.E{Key: types.Output, Value: opt.Output})
		}
		if opt.Granularity != "" {
			d = append(d, bson.E{Key: types.Granularity, Value: opt.Granularity})
		}
	}
	b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageBucketAuto, Value: d}})
	return b
}

func (b *StageBuilder) Match(expression any) *StageBuilder {
	b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageMatch, Value: expression}})
	return b
}

func (b *StageBuilder) Group(id any, accumulators ...any) *StageBuilder {
	d := bson.D{{Key: "_id", Value: id}}
	if accumulators != nil && len(accumulators)%2 == 0 {
		for i := 0; i < len(accumulators); i += 2 {
			k, ok := accumulators[i].(string)
			if !ok {
				continue
			}
			d = append(d, bson.E{Key: k, Value: accumulators[i+1]})
		}
	}
	b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageGroup, Value: d}})
	return b
}

func (b *StageBuilder) GroupMap(id any, accumulators map[string]map[string]any) *StageBuilder {
	d := bson.D{{Key: "_id", Value: id}}
	bsonAccumulators := converter.MapToBson(accumulators)
	d = append(d, bsonAccumulators...)
	b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageGroup, Value: d}})
	return b
}

func (b *StageBuilder) Sort(keyValues ...any) *StageBuilder {
	if keyValues != nil && len(keyValues)%2 == 0 {
		d := bson.D{}
		for i := 0; i < len(keyValues); i += 2 {
			k, ok := keyValues[i].(string)
			if !ok {
				continue
			}
			d = append(d, bson.E{Key: k, Value: keyValues[i+1]})
		}
		b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageSort, Value: d}})
	}
	return b
}

func (b *StageBuilder) SortMap(keyValues map[string]any) *StageBuilder {
	if keyValues != nil {
		d := bson.D{}
		for k, v := range keyValues {
			d = append(d, bson.E{Key: k, Value: v})
		}
		b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageSort, Value: d}})
	}
	return b
}

func (b *StageBuilder) Project(keyValues ...any) *StageBuilder {
	if keyValues != nil && len(keyValues)%2 == 0 {
		d := bson.D{}
		for i := 0; i < len(keyValues); i += 2 {
			k, ok := keyValues[i].(string)
			if !ok {
				continue
			}
			d = append(d, bson.E{Key: k, Value: keyValues[i+1]})
		}
		b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageProject, Value: d}})
	}
	return b
}

func (b *StageBuilder) ProjectMap(keyValues map[string]any) *StageBuilder {
	if keyValues != nil {
		d := bson.D{}
		for k, v := range keyValues {
			d = append(d, bson.E{Key: k, Value: v})
		}
		b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageProject, Value: d}})
	}
	return b
}

func (b *StageBuilder) Limit(limit int64) *StageBuilder {
	b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageLimit, Value: limit}})
	return b
}

func (b *StageBuilder) Skip(skip int64) *StageBuilder {
	b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageSkip, Value: skip}})
	return b
}

func (b *StageBuilder) Unwind(path string, opt *types.UnWindOptions) *StageBuilder {
	if opt == nil {
		b.pipeline = append(b.pipeline, bson.D{{Key: types.AggregationStageUnwind, Value: path}})
	} else {
		d := bson.D{{Key: "path", Value: path}}
		if opt.IncludeArrayIndex != "" {
			d = append(d, bson.E{Key: "includeArrayIndex", Value: opt.IncludeArrayIndex})
		}
		if opt.PreserveNullAndEmptyArrays {
			d = append(d, bson.E{Key: "preserveNullAndEmptyArrays", Value: opt.PreserveNullAndEmptyArrays})
		}
		b.pipeline = append(b.pipeline, bson.D{{Key: types.AggregationStageUnwind, Value: d}})
	}
	return b
}

func (b *StageBuilder) ReplaceWith(replacementDocument any) *StageBuilder {
	b.pipeline = append(b.pipeline, bson.D{{Key: types.AggregationStageReplaceWith, Value: replacementDocument}})
	return b
}

func (b *StageBuilder) Facet(keyValues ...any) *StageBuilder {
	if keyValues != nil && len(keyValues)%2 == 0 {
		d := bson.D{}
		for i := 0; i < len(keyValues); i += 2 {
			k, ok := keyValues[i].(string)
			if !ok {
				continue
			}
			d = append(d, bson.E{Key: k, Value: keyValues[i+1]})
		}
		b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageFacet, Value: d}})
	}
	return b
}

func (b *StageBuilder) FacetMap(keyValues map[string]any) *StageBuilder {
	if keyValues != nil {
		d := bson.D{}
		for k, v := range keyValues {
			d = append(d, bson.E{Key: k, Value: v})
		}
		b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageFacet, Value: d}})
	}
	return b
}

func (b *StageBuilder) SortByCount(expression any) *StageBuilder {
	b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageSortByCount, Value: expression}})
	return b
}

func (b *StageBuilder) Count(countName string) *StageBuilder {
	b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageCount, Value: countName}})
	return b
}

func (b *StageBuilder) Build() mongo.Pipeline {
	return b.pipeline
}
