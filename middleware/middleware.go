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

package middleware

import (
	"context"

	"github.com/chenmingyong0423/go-mongox/operation"
)

type callback func(ctx context.Context, doc any, opType operation.OpType, opts ...any) error

var middlewares = []callback{
	nil,
}

func Register(cb callback) {
	middlewares = append(middlewares, cb)
}

func Execute(ctx context.Context, doc any, opType operation.OpType, opts ...any) error {
	for _, middleware := range middlewares {
		if err := middleware(ctx, doc, opType, opts...); err != nil {
			return err
		}
	}
	return nil
}
