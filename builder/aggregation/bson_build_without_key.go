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

package aggregation

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func SumWithoutKey(expression any) bson.D {
	return bson.D{{Key: SumOp, Value: expression}}
}

func PushWithoutKey(expression any) bson.D {
	return bson.D{{Key: PushOp, Value: expression}}
}

func AvgWithoutKey(expression any) bson.D {
	return bson.D{{Key: AvgOp, Value: expression}}
}

func FirstWithoutKey(expression any) bson.D {
	return bson.D{{Key: FirstOp, Value: expression}}
}

func LastWithoutKey(expression any) bson.D {
	return bson.D{{Key: LastOp, Value: expression}}
}

func MinWithoutKey(expression any) bson.D {
	return bson.D{{Key: MinOp, Value: expression}}
}

func MaxWithoutKey(expression any) bson.D {
	return bson.D{{Key: MaxOp, Value: expression}}
}

func AddWithoutKey(expression ...any) bson.D {
	return bson.D{{Key: AddOp, Value: expression}}
}

func MultiplyWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: MultiplyOp, Value: expressions}}
}

func SubtractWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: SubtractOp, Value: expressions}}
}

func DivideWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: DivideOp, Value: expressions}}
}

func ModWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: ModOp, Value: expressions}}
}

func ArrayElemAtWithoutKey(expression any, index int64) bson.D {
	return bson.D{{Key: ArrayElemAtOp, Value: []any{expression, index}}}
}

func ConcatArraysWithoutKey(arrays ...any) bson.D {
	return bson.D{{Key: ConcatArraysOp, Value: arrays}}
}

func ArrayToObjectWithoutKey(expression any) bson.D {
	return bson.D{{Key: ArrayToObjectOp, Value: expression}}
}

func SizeWithoutKey(expression any) bson.D {
	return bson.D{{Key: SizeOp, Value: expression}}
}

func SliceWithoutKey(array any, nElements int64) bson.D {
	return bson.D{{Key: SliceOp, Value: []any{array, nElements}}}
}

func SliceWithPositionWithoutKey(array any, position, nElements int64) bson.D {
	return bson.D{{Key: SliceOp, Value: []any{array, position, nElements}}}
}

func MapWithoutKey(inputArray any, as string, in any) bson.D {
	return bson.D{{Key: MapOp, Value: bson.D{
		{Key: InputOp, Value: inputArray},
		{Key: AsOp, Value: as},
		{Key: InOp, Value: in},
	}}}
}

func FilterWithoutKey(inputArray any, cond any, opt *FilterOptions) bson.D {
	d := bson.D{{Key: InputOp, Value: inputArray}, {Key: CondWithoutOperatorOp, Value: cond}}
	if opt != nil {
		if opt.As != "" {
			d = append(d, bson.E{Key: AsOp, Value: opt.As})
		}
		if opt.Limit != 0 {
			d = append(d, bson.E{Key: LIMIT, Value: opt.Limit})
		}
	}
	return bson.D{{Key: FilterOp, Value: d}}
}

func EqWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: EqOp, Value: expressions}}
}

func NeWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: NeOp, Value: expressions}}
}

func GtWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: GtOp, Value: expressions}}
}

func GteWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: GteOp, Value: expressions}}
}

func LtWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: LtOp, Value: expressions}}
}

func LteWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: LteOp, Value: expressions}}
}

func CondWithoutKey(boolExpr, tureExpr, falseExpr any) bson.D {
	return bson.D{{Key: CondOp, Value: []any{boolExpr, tureExpr, falseExpr}}}
}

func IfNullWithoutKey(expr, replacement any) bson.D {
	return bson.D{{Key: IfNullOp, Value: []any{expr, replacement}}}
}
func SwitchWithoutKey(cases []CaseThen, defaultCase any) bson.D {
	branches := bson.A{}
	for _, caseThen := range cases {
		branches = append(branches, bson.D{bson.E{Key: CaseOp, Value: caseThen.Case}, {Key: ThenOp, Value: caseThen.Then}})
	}
	return bson.D{bson.E{Key: SwitchOp, Value: bson.D{bson.E{Key: BranchesOp, Value: branches}, bson.E{Key: DefaultCaseOp, Value: defaultCase}}}}
}

func DateToStringWithoutKey(date any, opt *DateToStringOptions) bson.D {
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
	return bson.D{{Key: DateToStringOp, Value: d}}
}

func DayOfMonthWithoutKey(date time.Time) bson.D {
	return bson.D{{Key: DayOfMonthOp, Value: date}}
}

func DayOfMonthWithTimezoneWithoutKey(date time.Time, timezone string) bson.D {
	return bson.D{{Key: DayOfMonthOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}}}
}

func DayOfWeekWithoutKey(date time.Time) bson.D {
	return bson.D{{Key: DayOfWeekOp, Value: date}}
}

func DayOfWeekWithTimezoneWithoutKey(date time.Time, timezone string) bson.D {
	return bson.D{{Key: DayOfWeekOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}}}
}

func DayOfYearWithoutKey(date time.Time) bson.D {
	return bson.D{{Key: DayOfYearOp, Value: date}}
}

func DayOfYearWithTimezoneWithoutKey(date time.Time, timezone string) bson.D {
	return bson.D{{Key: DayOfYearOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}}}
}

func YearWithoutKey(date time.Time) bson.D {
	return bson.D{{Key: YearOp, Value: date}}
}

func YearWithTimezoneWithoutKey(date time.Time, timezone string) bson.D {
	return bson.D{{Key: YearOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}}}
}

func MonthWithoutKey(date time.Time) bson.D {
	return bson.D{{Key: MonthOp, Value: date}}
}

func MonthWithTimezoneWithoutKey(date time.Time, timezone string) bson.D {
	return bson.D{{Key: MonthOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}}}
}

func WeekWithoutKey(date time.Time) bson.D {
	return bson.D{{Key: WeekOp, Value: date}}
}

func WeekWithTimezoneWithoutKey(date time.Time, timezone string) bson.D {
	return bson.D{{Key: WeekOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}}}
}

func AndWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: AndOp, Value: expressions}}
}

func OrWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: OrOp, Value: expressions}}
}

func NotWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: NotOp, Value: expressions}}
}

func ConcatWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: ConcatOp, Value: expressions}}
}

func SubstrBytesWithoutKey(stringExpression string, byteIndex int64, byteCount int64) bson.D {
	return bson.D{{Key: SubstrBytesOp, Value: []any{stringExpression, byteIndex, byteCount}}}
}

func ToLowerWithoutKey(expression any) bson.D {
	return bson.D{{Key: ToLowerOp, Value: expression}}
}

func ToUpperWithoutKey(expression any) bson.D {
	return bson.D{{Key: ToUpperOp, Value: expression}}
}

func ContactWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: ContactOp, Value: expressions}}
}
