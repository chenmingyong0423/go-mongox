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

func (b *StageBuilder) AddFields(value any) *StageBuilder {
	b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageAddFields, Value: value}})
	return b
}

func (b *StageBuilder) Set(value any) *StageBuilder {
	b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageSet, Value: value}})
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

func (b *StageBuilder) Group(id any, accumulators ...bson.E) *StageBuilder {
	d := bson.D{{Key: "_id", Value: id}}
	d = append(d, accumulators...)
	b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageGroup, Value: d}})
	return b
}

func (b *StageBuilder) Sort(value any) *StageBuilder {
	b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageSort, Value: value}})
	return b
}

func (b *StageBuilder) Project(value any) *StageBuilder {
	b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageProject, Value: value}})
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

func (b *StageBuilder) Facet(value any) *StageBuilder {
	b.pipeline = append(b.pipeline, bson.D{bson.E{Key: types.AggregationStageFacet, Value: value}})
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

func (b *StageBuilder) Lookup(from, as string, opt *LookUpOptions) *StageBuilder {
	d := bson.D{bson.E{Key: "from", Value: from}}
	if opt.LocalField != "" && opt.ForeignField != "" {
		d = append(d, bson.E{Key: "localField", Value: opt.LocalField})
		d = append(d, bson.E{Key: "foreignField", Value: opt.ForeignField})
	}
	if len(opt.Let) > 0 {
		d = append(d, bson.E{Key: "let", Value: opt.Let})
	}
	if len(opt.Pipeline) > 0 {
		d = append(d, bson.E{Key: "pipeline", Value: opt.Pipeline})
	}
	d = append(d, bson.E{Key: "as", Value: as})
	b.pipeline = append(b.pipeline, bson.D{bson.E{Key: StageLookUp, Value: d}})
	return b
}

func (b *StageBuilder) Build() mongo.Pipeline {
	return b.pipeline
}
