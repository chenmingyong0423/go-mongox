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
	"time"

	mongoxErr "github.com/chenmingyong0423/go-mongox/error"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Test_e2e_Open(t *testing.T) {

	testCases := []struct {
		name string

		ctx context.Context
		db  string
		cfg Config

		hasError bool
	}{
		{
			name: "error uri",
			ctx:  context.Background(),
			db:   "dd-test",
			cfg: Config{
				ClientOpts: []*options.ClientOptions{options.Client().ApplyURI("invalid-mongo-uri")},
				DbOpts:     nil,
			},
			hasError: true,
		},
		{
			name: "error port",
			ctx:  context.Background(),
			db:   "db-test",
			cfg: Config{
				ClientOpts: []*options.ClientOptions{options.Client().ApplyURI("mongodb://localhost:27018").SetAuth(options.Credential{
					Username:   "test",
					Password:   "test",
					AuthSource: "db-test",
				})},
				DbOpts: nil,
			},
			hasError: true,
		},
		{
			name: "no error",
			ctx:  context.Background(),
			db:   "db-test",
			cfg: Config{
				ClientOpts: []*options.ClientOptions{options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
					Username:   "test",
					Password:   "test",
					AuthSource: "db-test",
				})},
				DbOpts: nil,
			},
			hasError: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotMc, err := Open(tc.ctx, tc.db, tc.cfg)
			if err == nil {
				withTimeout, cancelFunc := context.WithTimeout(tc.ctx, 5*time.Second)
				defer cancelFunc()
				err = gotMc.client.Ping(withTimeout, readpref.Primary())
				defer func() {
					assert.NoError(t, gotMc.Disconnect(tc.ctx))
				}()
			}
			if tc.hasError {
				assert.NotNil(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_e2e_Coll(t *testing.T) {
	testCases := []struct {
		name string

		model    any
		opts     []*options.CollectionOptions
		collName string

		wantErr error
	}{
		{
			name:  "nil model",
			model: nil,

			wantErr: mongoxErr.ErrModelIsNil,
		},
		{
			name:  "invalid model type",
			model: 1234,

			wantErr: mongoxErr.ErrInvalidModelType,
		},
		{
			name: "no error when the model type is multiple pointer",
			model: func() **testUser {
				res := &testUser{}
				return &res
			}(),
			collName: "",
			wantErr:  nil,
		},
		{
			name:     "no error when the model type is struct without implement CollectionName interface",
			model:    testUser{},
			collName: "testUser",
			wantErr:  nil,
		},
		{
			name:     "no error when the model type is struct pointer without implement CollectionName interface",
			model:    &testUser{},
			collName: "testUser",
			wantErr:  nil,
		},
		{
			name:     "no error when the model type is struct which implements CollectionName interface",
			model:    testComment{},
			collName: "test_comment",
			wantErr:  nil,
		},
		{
			name:     "no error when the model type is struct which implements CollectionName interface by pointer",
			model:    testPost{},
			collName: "testPost",
			wantErr:  nil,
		},
		{
			name:     "no error when the model type is struct pointer which implements CollectionName interface",
			model:    &testPost{},
			collName: "test_post",
			wantErr:  nil,
		},
		{
			name:     "no error when the model type is string",
			model:    "test_post",
			collName: "test_post",
			wantErr:  nil,
		},
		{
			name:     "no error when the model type is string with an empty value",
			model:    "",
			collName: "",
			wantErr:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			withTimeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
			defer cancelFunc()
			mongoClient, err := Open(withTimeout, "db-test", Config{
				ClientOpts: []*options.ClientOptions{options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
					Username:   "test",
					Password:   "test",
					AuthSource: "db-test",
				})},
				DbOpts: nil,
			})
			assert.NoError(t, err)
			coll := mongoClient.Coll(tc.model, tc.opts...)
			assert.Equal(t, tc.wantErr, coll.err)
			assert.Equal(t, tc.collName, coll.coll.Name())
		})
	}
}

type testUser struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

type testPost struct {
}

func (t *testPost) CollectionName() string {
	return "test_post"
}

type testComment struct {
}

func (t testComment) CollectionName() string {
	return "test_comment"
}
