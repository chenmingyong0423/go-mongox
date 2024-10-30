// Copyright 2024 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validator

import (
	"context"
	"errors"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/go-playground/validator/v10"

	"github.com/stretchr/testify/require"

	"github.com/chenmingyong0423/go-mongox/v2/operation"
)

type User struct {
	Name string `bson:"name"`
	Age  int    `bson:"age" validate:"gte=0,lte=150"`
}

func TestExecute(t *testing.T) {

	testCases := []struct {
		name string

		ctx    context.Context
		doc    *operation.OpContext
		opType operation.OpType

		errFunc require.ErrorAssertionFunc
	}{
		{
			name:    "unaccepted operation type",
			ctx:     context.Background(),
			doc:     operation.NewOpContext(nil, operation.WithDoc(&User{Name: "Mingyong Chen", Age: 18})),
			opType:  operation.OpTypeAfterInsert,
			errFunc: require.NoError,
		},
		{
			name:    "nil value",
			ctx:     context.Background(),
			doc:     operation.NewOpContext(nil),
			opType:  operation.OpTypeBeforeInsert,
			errFunc: require.NoError,
		},
		{
			name:    "unsupported type",
			ctx:     context.Background(),
			doc:     operation.NewOpContext(nil, operation.WithDoc(6)),
			opType:  operation.OpTypeBeforeInsert,
			errFunc: require.NoError,
		},
		{
			name: "unsupported point type",
			ctx:  context.Background(),
			doc: operation.NewOpContext(nil, operation.WithDoc(func() *int {
				i := 6
				return &i
			}())),
			opType:  operation.OpTypeBeforeInsert,
			errFunc: require.NoError,
		},
		{
			name:    "special unsupported type - time.Time{}",
			ctx:     context.Background(),
			doc:     operation.NewOpContext(nil, operation.WithDoc(&time.Time{})),
			opType:  operation.OpTypeBeforeInsert,
			errFunc: require.NoError,
		},
		{
			name:    "*User(nil)",
			ctx:     context.Background(),
			doc:     operation.NewOpContext(nil, operation.WithDoc((*User)(nil))),
			opType:  operation.OpTypeBeforeInsert,
			errFunc: require.NoError,
		},
		{
			name:   "fails to validate struct in case of BeforeInsert",
			ctx:    context.Background(),
			doc:    operation.NewOpContext(nil, operation.WithDoc(&User{Age: -1})),
			opType: operation.OpTypeBeforeInsert,
			errFunc: func(t require.TestingT, err error, i ...interface{}) {
				var e validator.ValidationErrors
				if !errors.As(err, &e) {
					log.Fatal(err)
				}
			},
		},
		{
			name:   "fails to validate struct in case of BeforeUpsert",
			ctx:    context.Background(),
			doc:    operation.NewOpContext(nil, operation.WithReplacement(&User{Age: -1})),
			opType: operation.OpTypeBeforeUpsert,
			errFunc: func(t require.TestingT, err error, i ...interface{}) {
				var e validator.ValidationErrors
				if !errors.As(err, &e) {
					log.Fatal(err)
				}
			},
		},
		{
			name: "fails to validate slice in case of BeforeInsert",
			ctx:  context.Background(),
			doc: operation.NewOpContext(nil, operation.WithDoc([]*User{
				{Age: -1},
				{Age: 18},
			})),
			opType: operation.OpTypeBeforeInsert,
			errFunc: func(t require.TestingT, err error, i ...interface{}) {
				var e validator.ValidationErrors
				if !errors.As(err, &e) {
					log.Fatal(err)
				}
			},
		},
		{
			name:    "validate struct successfully in case of BeforeInsert",
			ctx:     context.Background(),
			doc:     operation.NewOpContext(nil, operation.WithDoc(&User{Age: 18})),
			opType:  operation.OpTypeBeforeInsert,
			errFunc: require.NoError,
		},
		{
			name: "validate slice successfully in case of BeforeInsert",
			ctx:  context.Background(),
			doc: operation.NewOpContext(nil, operation.WithDoc([]*User{
				{Age: 18},
				{Age: 20},
			})),
			opType:  operation.OpTypeBeforeInsert,
			errFunc: require.NoError,
		},
		{
			name:    "validate struct successfully in case of BeforeUpsert",
			ctx:     context.Background(),
			doc:     operation.NewOpContext(nil, operation.WithReplacement(&User{Age: 18})),
			opType:  operation.OpTypeBeforeUpsert,
			errFunc: require.NoError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := Execute(tc.ctx, tc.doc, tc.opType)
			tc.errFunc(t, err)
		})
	}
}

func Test_getPayload(t *testing.T) {
	testCases := []struct {
		name   string
		opCtx  *operation.OpContext
		opType operation.OpType

		want any
	}{
		{
			name:   "nil opCtx",
			opCtx:  nil,
			opType: operation.OpTypeBeforeInsert,

			want: nil,
		},
		{
			name:   "empty opCtx",
			opCtx:  &operation.OpContext{},
			opType: operation.OpTypeBeforeInsert,

			want: nil,
		},
		{
			name:   "unexpected opType",
			opCtx:  &operation.OpContext{Doc: &User{}},
			opType: "unexpected",

			want: nil,
		},
		{
			name:   "BeforeInsert",
			opCtx:  &operation.OpContext{Doc: &User{}},
			opType: operation.OpTypeBeforeInsert,

			want: &User{},
		},
		{
			name:   "BeforeUpsert",
			opCtx:  &operation.OpContext{Replacement: &User{}},
			opType: operation.OpTypeBeforeUpsert,

			want: &User{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := getPayload(tc.opCtx, tc.opType)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSetValidate(t *testing.T) {
	v := validate
	SetValidate(nil)
	assert.Equal(t, v, validate)

	v = validator.New()
	SetValidate(v)
	assert.Equal(t, v, validate)
}
