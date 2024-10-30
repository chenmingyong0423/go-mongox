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
	"go.mongodb.org/mongo-driver/v2/bson"
)

type arithmeticBuilder struct {
	parent *Builder
}

func (b *arithmeticBuilder) Abs(key string, expression any) *Builder {
	e := bson.E{Key: AbsOp, Value: expression}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arithmeticBuilder) AbsWithoutKey(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: AbsOp, Value: expression})
	return b.parent
}

func (b *arithmeticBuilder) Add(key string, expressions ...any) *Builder {
	e := bson.E{Key: AddOp, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arithmeticBuilder) AddWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: AddOp, Value: expressions})
	return b.parent
}

func (b *arithmeticBuilder) Ceil(key string, expression any) *Builder {
	e := bson.E{Key: CeilOp, Value: expression}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arithmeticBuilder) CeilWithoutKey(expression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: CeilOp, Value: expression})
	return b.parent
}

func (b *arithmeticBuilder) Divide(key string, expressions ...any) *Builder {
	e := bson.E{Key: DivideOp, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arithmeticBuilder) DivideWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: DivideOp, Value: expressions})
	return b.parent
}

func (b *arithmeticBuilder) Exp(key string, exponent any) *Builder {
	e := bson.E{Key: ExpOp, Value: exponent}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arithmeticBuilder) ExpWithoutKey(exponent any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: ExpOp, Value: exponent})
	return b.parent
}

func (b *arithmeticBuilder) Floor(key string, numberExpression any) *Builder {
	e := bson.E{Key: FloorOp, Value: numberExpression}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arithmeticBuilder) FloorWithoutKey(numberExpression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: FloorOp, Value: numberExpression})
	return b.parent
}

func (b *arithmeticBuilder) Ln(key string, numberExpression any) *Builder {
	e := bson.E{Key: LnOp, Value: numberExpression}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arithmeticBuilder) LnWithoutKey(numberExpression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: LnOp, Value: numberExpression})
	return b.parent
}

func (b *arithmeticBuilder) Log(key string, numberExpression, baseNumberExpression any) *Builder {
	e := bson.E{Key: LogOp, Value: bson.A{numberExpression, baseNumberExpression}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arithmeticBuilder) LogWithoutKey(numberExpression, baseNumberExpression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: LogOp, Value: bson.A{numberExpression, baseNumberExpression}})
	return b.parent
}

func (b *arithmeticBuilder) Log10(key string, numberExpression any) *Builder {
	e := bson.E{Key: Log10Op, Value: bson.A{numberExpression, 10}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arithmeticBuilder) Log10WithoutKey(numberExpression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: Log10Op, Value: bson.A{numberExpression, 10}})
	return b.parent
}

func (b *arithmeticBuilder) Mod(key string, expressions ...any) *Builder {
	e := bson.E{Key: ModOp, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arithmeticBuilder) ModWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: ModOp, Value: expressions})
	return b.parent
}

func (b *arithmeticBuilder) Multiply(key string, expressions ...any) *Builder {
	e := bson.E{Key: MultiplyOp, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arithmeticBuilder) MultiplyWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: MultiplyOp, Value: expressions})
	return b.parent
}

func (b *arithmeticBuilder) Pow(key string, numberExpression, exponentExpression any) *Builder {
	e := bson.E{Key: PowOp, Value: bson.A{numberExpression, exponentExpression}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arithmeticBuilder) PowWithoutKey(numberExpression, exponentExpression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: PowOp, Value: bson.A{numberExpression, exponentExpression}})
	return b.parent
}

func (b *arithmeticBuilder) Round(key string, numberExpression, placeExpression any) *Builder {
	e := bson.E{Key: RoundOp, Value: bson.A{numberExpression, placeExpression}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arithmeticBuilder) RoundWithoutKey(numberExpression, placeExpression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: RoundOp, Value: bson.A{numberExpression, placeExpression}})
	return b.parent
}

func (b *arithmeticBuilder) Sqrt(key string, numberExpression any) *Builder {
	e := bson.E{Key: SqrtOp, Value: numberExpression}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arithmeticBuilder) SqrtWithoutKey(numberExpression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: SqrtOp, Value: numberExpression})
	return b.parent
}

func (b *arithmeticBuilder) Subtract(key string, expressions ...any) *Builder {
	e := bson.E{Key: SubtractOp, Value: expressions}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arithmeticBuilder) SubtractWithoutKey(expressions ...any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: SubtractOp, Value: expressions})
	return b.parent
}

func (b *arithmeticBuilder) Trunc(key string, numberExpression, placeExpression any) *Builder {
	e := bson.E{Key: TruncOp, Value: bson.A{numberExpression, placeExpression}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *arithmeticBuilder) TruncWithoutKey(numberExpression, placeExpression any) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: TruncOp, Value: bson.A{numberExpression, placeExpression}})
	return b.parent
}
