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
	"testing"

	"github.com/chenmingyong0423/go-mongox/v2/updater"

	"github.com/chenmingyong0423/go-mongox/v2/creator"

	"github.com/chenmingyong0423/go-mongox/v2/finder"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func TestCollection_New(t *testing.T) {
	result := NewCollection[any](NewClient(&mongo.Client{}, &Config{}).NewDatabase("db-test"), "collection-test")

	assert.NotNil(t, result, "Expected non-nil Collection")
}

func TestCollection_Finder(t *testing.T) {
	f := finder.NewFinder[any](&mongo.Collection{}, nil, nil)
	assert.NotNil(t, f, "Expected non-nil Finder")
}

func TestCollection_Creator(t *testing.T) {
	c := creator.NewCreator[any](&mongo.Collection{}, nil, nil)
	assert.NotNil(t, c, "Expected non-nil Creator")
}

func TestCollection_Updater(t *testing.T) {
	u := updater.NewUpdater[any](&mongo.Collection{}, nil, nil)
	assert.NotNil(t, u, "Expected non-nil Updater")
}

func TestCollection_Deleter(t *testing.T) {
	d := NewCollection[any](NewClient(&mongo.Client{}, &Config{}).NewDatabase("db-test"), "collection-test").Deleter()
	assert.NotNil(t, d, "Expected non-nil Deleter")
}

func TestCollection_Aggregator(t *testing.T) {

	a := NewCollection[any](NewClient(&mongo.Client{}, &Config{}).NewDatabase("db-test"), "collection-test").Aggregator()
	assert.NotNil(t, a, "Expected non-nil Aggregator")
}

func TestCollection_Collection(t *testing.T) {
	a := NewCollection[any](NewClient(&mongo.Client{}, &Config{}).NewDatabase("db-test"), "collection-test")
	assert.NotNil(t, a.Collection(), "Expected non-nil *mongo.Collection")
}
