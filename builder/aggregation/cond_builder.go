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

func (b condBuilder) Cond(key string, boolExpr, tureExpr, falseExpr any) *Builder {
	e := bson.E{Key: types.AggregationCond, Value: []any{boolExpr, tureExpr, falseExpr}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b condBuilder) CondWithoutKey(boolExpr, tureExpr, falseExpr any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationCond, Value: []any{boolExpr, tureExpr, falseExpr}})
	return b.parent
}

func (b condBuilder) IfNull(key string, expr, replacement any) *Builder {
	e := bson.E{Key: types.AggregationIfNull, Value: []any{expr, replacement}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b condBuilder) IfNullWithoutKey(expr, replacement any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationIfNull, Value: []any{expr, replacement}})
	return b.parent
}

func (b condBuilder) Switch(key string, cases []types.CaseThen, defaultCase any) *Builder {
	branches := bson.A{}
	for _, caseThen := range cases {
		branches = append(branches, bson.D{{Key: types.Case, Value: caseThen.Case}, {Key: types.Then, Value: caseThen.Then}})
	}
	e := bson.E{Key: types.AggregationSwitch, Value: bson.D{bson.E{Key: types.Branches, Value: branches}, bson.E{Key: types.DefaultCase, Value: defaultCase}}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b condBuilder) SwitchWithoutKey(cases []types.CaseThen, defaultCase any) *Builder {
	branches := bson.A{}
	for _, caseThen := range cases {
		branches = append(branches, bson.D{bson.E{Key: types.Case, Value: caseThen.Case}, {Key: types.Then, Value: caseThen.Then}})
	}
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationSwitch, Value: bson.D{bson.E{Key: types.Branches, Value: branches}, bson.E{Key: types.DefaultCase, Value: defaultCase}}})
	return b.parent
}
