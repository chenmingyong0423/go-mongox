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
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Client struct {
	client *mongo.Client
	cfg    *Config
}

func NewClient(client *mongo.Client, config *Config) *Client {
	return &Client{
		client: client,
		cfg:    config,
	}
}

// Client returns the mongo client
func (c *Client) Client() *mongo.Client {
	return c.client
}

func (c *Client) config() *Config {
	return c.cfg
}

func (c *Client) Disconnect(ctx context.Context) error {
	return c.client.Disconnect(ctx)
}

func (c *Client) NewDatabase(database string) *Database {
	return newDatabase(c, database)
}
