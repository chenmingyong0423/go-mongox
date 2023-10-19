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
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCollection_e2e_FindById(t *testing.T) {
	clientErr := errors.New("client error")

	gotMc, err := Open(context.Background(), "db-test", Config{
		ClientOpts: []*options.ClientOptions{options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
			Username:   "test",
			Password:   "test",
			AuthSource: "db-test",
		})},
		DbOpts: nil,
	})
	assert.NoError(t, err)

	coll := gotMc.Coll("testUser")

	testCases := []struct {
		name string

		before func(ctx context.Context, t *testing.T)
		after  func(ctx context.Context, t *testing.T)

		ctx       context.Context
		id        any
		expectPtr any
		opts      []*options.FindOneOptions

		wantExpect any
		wantErr    error
	}{
		{
			name: "client error",
			before: func(ctx context.Context, t *testing.T) {
				coll.err = clientErr
			},
			after: func(ctx context.Context, t *testing.T) {
				coll.err = nil
			},

			wantErr: clientErr,
		},
		{
			name: "no document",
			ctx:  context.Background(),

			before: func(ctx context.Context, t *testing.T) {
			},
			after: func(ctx context.Context, t *testing.T) {
			},

			id:        "1",
			expectPtr: &testUser{},

			wantExpect: &testUser{},
			wantErr:    mongo.ErrNoDocuments,
		},
		{
			name: "found",
			ctx:  context.Background(),

			before: func(ctx context.Context, t *testing.T) {
				_, fErr := coll.coll.InsertOne(ctx, testUser{Id: "1", Name: "cmy"})
				assert.NoError(t, fErr)
			},
			after: func(ctx context.Context, t *testing.T) {
				_, fErr := coll.coll.DeleteOne(ctx, bson.D{bson.E{Key: "_id", Value: 1}})
				assert.NoError(t, fErr)
			},

			id:        "1",
			expectPtr: &testUser{},

			wantExpect: &testUser{Id: "1", Name: "cmy"},
			wantErr:    nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(tc.ctx, t)
			tErr := coll.FindById(tc.ctx, tc.id, tc.expectPtr, tc.opts...)
			tc.after(tc.ctx, t)
			assert.Equal(t, tc.wantErr, tErr)
			assert.Equal(t, tc.wantExpect, tc.expectPtr)
		})
	}
}
