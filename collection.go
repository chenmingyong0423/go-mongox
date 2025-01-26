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
	"github.com/chenmingyong0423/go-mongox/v2/aggregator"
	"github.com/chenmingyong0423/go-mongox/v2/callback"
	"github.com/chenmingyong0423/go-mongox/v2/creator"
	"github.com/chenmingyong0423/go-mongox/v2/deleter"
	"github.com/chenmingyong0423/go-mongox/v2/finder"
	"github.com/chenmingyong0423/go-mongox/v2/updater"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewCollection[T any](db *Database, collection string) *Collection[T] {
	return &Collection[T]{
		db:         db,
		collection: db.database().Collection(collection),
		callbacks:  db.callbacks,
	}
}

type Collection[T any] struct {
	db         *Database
	collection *mongo.Collection
	// callbacks inherited from database
	callbacks *callback.Callback
}

func (c *Collection[T]) Finder() *finder.Finder[T] {
	return finder.NewFinder[T](c.collection, c.callbacks)
}

func (c *Collection[T]) Creator() *creator.Creator[T] {
	return creator.NewCreator[T](c.collection, c.callbacks)
}

func (c *Collection[T]) Updater() *updater.Updater[T] {
	return updater.NewUpdater[T](c.collection, c.callbacks)
}

func (c *Collection[T]) Deleter() *deleter.Deleter[T] {
	return deleter.NewDeleter[T](c.collection, c.callbacks)
}
func (c *Collection[T]) Aggregator() *aggregator.Aggregator[T] {
	return aggregator.NewAggregator[T](c.collection)
}

func (c *Collection[T]) Collection() *mongo.Collection {
	return c.collection
}
