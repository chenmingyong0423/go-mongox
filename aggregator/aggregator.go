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

package aggregator

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=aggregator.go -destination=../mock/aggregator.mock.go -package=mocks
type iAggregator[T any] interface {
	Aggregate(ctx context.Context, opts ...*options.AggregateOptions) ([]*T, error)
	AggregateWithParse(ctx context.Context, result any, opts ...*options.AggregateOptions) error
}

type Aggregator[T any] struct {
	collection *mongo.Collection
	pipeline   any
}

func NewAggregator[T any](collection *mongo.Collection) *Aggregator[T] {
	return &Aggregator[T]{
		collection: collection,
	}
}

func (a *Aggregator[T]) Pipeline(pipeline any) *Aggregator[T] {
	a.pipeline = pipeline
	return a
}

func (a *Aggregator[T]) Aggregate(ctx context.Context, opts ...*options.AggregateOptions) ([]*T, error) {
	cursor, err := a.collection.Aggregate(ctx, a.pipeline, opts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	result := make([]*T, 0)
	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AggregateWithParse is used to parse the result of the aggregation
// result must be a pointer to a slice
func (a *Aggregator[T]) AggregateWithParse(ctx context.Context, result any, opts ...*options.AggregateOptions) error {
	cursor, err := a.collection.Aggregate(ctx, a.pipeline, opts...)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, result)
	if err != nil {
		return err
	}
	return nil
}
