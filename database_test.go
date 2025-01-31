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
	"context"
	"testing"

	"github.com/chenmingyong0423/go-mongox/v2/operation"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func Test_newDatabase(t *testing.T) {
	db := newDatabase(NewClient(&mongo.Client{}, &Config{}), "db-test")
	require.Equal(t, db.Database().Name(), "db-test")

	db.RegisterPlugin("global before find", func(ctx context.Context, opCtx *operation.OpContext, opts ...any) error {
		return nil
	}, operation.OpTypeBeforeFind)

	db.RemovePlugin("global before find", operation.OpTypeBeforeFind)
}
