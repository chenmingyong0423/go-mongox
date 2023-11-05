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

package aggregation

import (
	"github.com/chenmingyong0423/go-mongox/types"
	"go.mongodb.org/mongo-driver/bson"
)

type condBuilder struct {
	parent *Builder
}

func (b condBuilder) Cond(boolExpr, tureExpr, falseExpr any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationCond, Value: []any{boolExpr, tureExpr, falseExpr}})
	return b.parent
}

func (b condBuilder) IfNull(expr, replacement any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationIfNull, Value: []any{expr, replacement}})
	return b.parent
}

// Switch
// cases: [case, then, case, then]
func (b condBuilder) Switch(cases []any, defaultCase any) *Builder {
	if len(cases) != 0 && len(cases)%2 == 0 {
		branches := bson.A{}
		for i := 0; i < len(cases); i += 2 {
			branches = append(branches, bson.D{{Key: types.Case, Value: cases[i]}, {Key: types.Then, Value: cases[i+1]}})
		}
		b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationSwitch, Value: bson.D{bson.E{Key: types.Branches, Value: branches}, bson.E{Key: types.DefaultCase, Value: defaultCase}}})
	}
	return b.parent
}
