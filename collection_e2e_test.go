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

//go:build e2e

package mongox

import (
	"context"
	"testing"

	"github.com/chenmingyong0423/go-mongox/creator"

	"github.com/chenmingyong0423/go-mongox/finder"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func TestCollection_e2e_Deleter(t *testing.T) {
	collection := getCollection[any](t)

	d := collection.Deleter()
	assert.NotNil(t, d, "Expected non-nil Deleter")
}

func TestCollection_e2e_Updater(t *testing.T) {
	collection := getCollection[any](t)

	u := collection.Updater()
	assert.NotNil(t, u, "Expected non-nil Updater")
}

func TestCollection_e2e_Finder(t *testing.T) {
	collection := getCollection[any](t)

	f := finder.NewFinder[any](collection.collection)
	assert.NotNil(t, f, "Expected non-nil Finder")
}

func TestCollection_e2e_Creator(t *testing.T) {
	collection := getCollection[any](t)

	c := creator.NewCreator[any](collection.collection)
	assert.NotNil(t, c, "Expected non-nil Creator")
}

func TestCollection_e2e_New(t *testing.T) {
	assert.NotNil(t, getCollection[any](t))
}

func getCollection[T any](t *testing.T) *Collection[T] {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		Username:   "test",
		Password:   "test",
		AuthSource: "db-test",
	}))
	assert.NoError(t, err)
	assert.NoError(t, client.Ping(context.Background(), readpref.Primary()))

	collection := NewCollection[T](client.Database("db-test").Collection("test_user"))
	return collection
}
