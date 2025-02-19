// Copyright 2025 chenmingyong0423

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
	"github.com/chenmingyong0423/go-mongox/v2/callback"
	"github.com/chenmingyong0423/go-mongox/v2/operation"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Database struct {
	client *Client
	db     *mongo.Database
	// callbacks for database
	callbacks *callback.Callback
}

func newDatabase(c *Client, database string) *Database {
	return &Database{
		client:    c,
		db:        c.client.Database(database),
		callbacks: callback.InitializeCallbacks(),
	}
}

func (d *Database) Database() *mongo.Database {
	return d.db
}

func (d *Database) RegisterPlugin(name string, cb callback.CbFn, opType operation.OpType) {
	d.callbacks.Register(opType, name, cb)
}

func (d *Database) RemovePlugin(name string, opType operation.OpType) {
	d.callbacks.Remove(opType, name)
}
