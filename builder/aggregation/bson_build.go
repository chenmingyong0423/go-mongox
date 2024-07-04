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
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func Add(key string, expression ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: AddOp, Value: expression}}}}
}

func And(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: AndOp, Value: expressions}}}}
}

func ArrayElemAt(key string, expression any, index int64) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: ArrayElemAtOp, Value: []any{expression, index}}}}}
}

func ArrayToObject(key string, expression any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: ArrayToObjectOp, Value: expression}}}}
}

func Avg(key string, expression any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: AvgOp, Value: expression}}}}
}

func Concat(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: ConcatOp, Value: expressions}}}}
}

func ConcatArrays(key string, arrays ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: ConcatArraysOp, Value: arrays}}}}
}

func Cond(key string, boolExpr, tureExpr, falseExpr any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: CondOp, Value: []any{boolExpr, tureExpr, falseExpr}}}}}
}

func Contact(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: ContactOp, Value: expressions}}}}
}

func DateToString(key string, date any, opt *DateToStringOptions) bson.D {
	d := bson.D{bson.E{Key: DateOp, Value: date}}
	if opt != nil {
		if opt.Format != "" {
			d = append(d, bson.E{Key: FormatOp, Value: opt.Format})
		}
		if opt.Timezone != "" {
			d = append(d, bson.E{Key: TimezoneOp, Value: opt.Timezone})
		}
		if opt.OnNull != nil {
			d = append(d, bson.E{Key: OnNullOp, Value: opt.OnNull})
		}
	}
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: DateToStringOp, Value: d}}}}
}

func DayOfMonth(key string, date time.Time) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: DayOfMonthOp, Value: date}}}}
}

func DayOfMonthWithTimezone(key string, date time.Time, timezone string) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: DayOfMonthOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}}}}}
}

func DayOfWeek(key string, date time.Time) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: DayOfWeekOp, Value: date}}}}
}

func DayOfWeekWithTimezone(key string, date time.Time, timezone string) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: DayOfWeekOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}}}}}
}

func DayOfYear(key string, date time.Time) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: DayOfYearOp, Value: date}}}}
}

func DayOfYearWithTimezone(key string, date time.Time, timezone string) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: DayOfYearOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}}}}}
}

func Divide(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: DivideOp, Value: expressions}}}}
}

func Eq(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: EqOp, Value: expressions}}}}
}

func Filter(key string, inputArray any, cond any, opt *FilterOptions) bson.D {
	d := bson.D{bson.E{Key: InputOp, Value: inputArray}, {Key: CondWithoutOperatorOp, Value: cond}}
	if opt != nil {
		if opt.As != "" {
			d = append(d, bson.E{Key: AsOp, Value: opt.As})
		}
		if opt.Limit != 0 {
			d = append(d, bson.E{Key: LIMIT, Value: opt.Limit})
		}
	}
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: FilterOp, Value: d}}}}
}

func First(key string, expression any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: FirstOp, Value: expression}}}}
}

func Gt(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: GtOp, Value: expressions}}}}
}

func Gte(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: GteOp, Value: expressions}}}}
}

func IfNull(key string, expr, replacement any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: IfNullOp, Value: []any{expr, replacement}}}}}
}

func Last(key string, expression any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: LastOp, Value: expression}}}}
}

func Lt(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: LtOp, Value: expressions}}}}
}

func Lte(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: LteOp, Value: expressions}}}}
}

func Map(key string, inputArray any, as string, in any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: MapOp, Value: bson.D{bson.E{Key: InputOp, Value: inputArray}, {Key: AsOp, Value: as}, {Key: InOp, Value: in}}}}}}
}

func Max(key string, expression any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: MaxOp, Value: expression}}}}
}

func Min(key string, expression any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: MinOp, Value: expression}}}}
}

func Mod(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: ModOp, Value: expressions}}}}
}

func Month(key string, date time.Time) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: MonthOp, Value: date}}}}
}

func MonthWithTimezone(key string, date time.Time, timezone string) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: MonthOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}}}}}
}

func Multiply(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: MultiplyOp, Value: expressions}}}}
}

func Ne(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: NeOp, Value: expressions}}}}
}

func Not(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: NotOp, Value: expressions}}}}
}

func Or(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: OrOp, Value: expressions}}}}
}

func Push(key string, expression any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: PushOp, Value: expression}}}}
}

func Size(key string, expression any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: SizeOp, Value: expression}}}}
}

func Slice(key string, array any, nElements int64) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: SliceOp, Value: []any{array, nElements}}}}}
}

func SliceWithPosition(key string, array any, position, nElements int64) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: SliceOp, Value: []any{array, position, nElements}}}}}
}

func SubstrBytes(key string, stringExpression string, byteIndex int64, byteCount int64) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: SubstrBytesOp, Value: []any{stringExpression, byteIndex, byteCount}}}}}
}

func Subtract(key string, expression ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: SubtractOp, Value: expression}}}}
}

func Sum(key string, expression any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: SumOp, Value: expression}}}}
}

func Switch(key string, cases []CaseThen, defaultCase any) bson.D {
	branches := bson.A{}
	for _, caseThen := range cases {
		branches = append(branches, bson.D{bson.E{Key: CaseOp, Value: caseThen.Case}, {Key: ThenOp, Value: caseThen.Then}})
	}
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: SwitchOp, Value: bson.D{bson.E{Key: BranchesOp, Value: branches}, bson.E{Key: DefaultCaseOp, Value: defaultCase}}}}}}
}

func ToLower(key string, expression any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: ToLowerOp, Value: expression}}}}
}

func ToUpper(key string, expression any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: ToUpperOp, Value: expression}}}}
}

func Week(key string, date time.Time) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: WeekOp, Value: date}}}}
}

func WeekWithTimezone(key string, date time.Time, timezone string) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: WeekOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}}}}}
}

func Year(key string, date time.Time) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: YearOp, Value: date}}}}
}

func YearWithTimezone(key string, date time.Time, timezone string) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: YearOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}}}}}
}
