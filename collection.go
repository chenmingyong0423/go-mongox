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

package mongox

import (
	"github.com/chenmingyong0423/go-mongox/aggregator"
	"github.com/chenmingyong0423/go-mongox/creator"
	"github.com/chenmingyong0423/go-mongox/deleter"
	"github.com/chenmingyong0423/go-mongox/finder"
	"github.com/chenmingyong0423/go-mongox/updater"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewCollection[T any](collection *mongo.Collection) *Collection[T] {
	return &Collection[T]{collection: collection}
}

type Collection[T any] struct {
	collection *mongo.Collection
}

func (c *Collection[T]) Finder() *finder.Finder[T] {
	return finder.NewFinder[T](c.collection)
}

func (c *Collection[T]) Creator() *creator.Creator[T] {
	return creator.NewCreator[T](c.collection)
}

func (c *Collection[T]) Updater() *updater.Updater[T] {
	return updater.NewUpdater[T](c.collection)
}

func (c *Collection[T]) Deleter() *deleter.Deleter[T] {
	return deleter.NewDeleter[T](c.collection)
}
func (c *Collection[T]) Aggregator() *aggregator.Aggregator[T] {
	return aggregator.NewAggregator[T](c.collection)
}

func (c *Collection[T]) Collection() *mongo.Collection {
	return c.collection
}
