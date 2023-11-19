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

	"github.com/chenmingyong0423/go-mongox/types"
	"go.mongodb.org/mongo-driver/bson"
)

func Sum(expression any) bson.D {
	return bson.D{{Key: types.AggregationSum, Value: expression}}
}

func SumMany(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationSum, Value: expressions}}
}

func Push(expression any) bson.D {
	return bson.D{{Key: types.AggregationPush, Value: expression}}
}

func Avg(expression any) bson.D {
	return bson.D{{Key: types.AggregationAvg, Value: expression}}
}

func AvgMany(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationAvg, Value: expressions}}
}

func First(expression any) bson.D {
	return bson.D{{Key: types.AggregationFirst, Value: expression}}
}

func Last(expression any) bson.D {
	return bson.D{{Key: types.AggregationLast, Value: expression}}
}

func Min(expression any) bson.D {
	return bson.D{{Key: types.AggregationMin, Value: expression}}
}

func MinMany(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationMin, Value: expressions}}
}

func Max(expression any) bson.D {
	return bson.D{{Key: types.AggregationMax, Value: expression}}
}

func MaxMany(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationMax, Value: expressions}}
}

func Add(expression any) bson.D {
	return bson.D{{Key: types.AggregationAdd, Value: expression}}
}

func Multiply(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationMultiply, Value: expressions}}
}

func Subtract(s string, start, length int64) bson.D {
	return bson.D{{Key: types.AggregationSubtract, Value: []any{s, start, length}}}
}

func Divide(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationDivide, Value: expressions}}
}

func Mod(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationMod, Value: expressions}}
}

func ArrayElemAt(expression any, index int64) bson.D {
	return bson.D{{Key: types.AggregationArrayElemAt, Value: []any{expression, index}}}
}

func ConcatArrays(arrays ...any) bson.D {
	return bson.D{{Key: types.AggregationConcatArrays, Value: arrays}}
}

func ArrayToObject(expression any) bson.D {
	return bson.D{{Key: types.AggregationArrayToObject, Value: expression}}
}

func Size(expression any) bson.D {
	return bson.D{{Key: types.AggregationSize, Value: expression}}
}

func Slice(array any, nElements int64) bson.D {
	return bson.D{{Key: types.AggregationSlice, Value: []any{array, nElements}}}
}

func SliceWithPosition(array any, position, nElements int64) bson.D {
	return bson.D{{Key: types.AggregationSlice, Value: []any{array, position, nElements}}}
}

func Map(inputArray any, as string, in any) bson.D {
	return bson.D{{Key: types.AggregationMap, Value: bson.D{
		{Key: types.AggregationInput, Value: inputArray},
		{Key: types.AggregationAs, Value: as},
		{Key: types.AggregationIn, Value: in},
	}}}
}

func Filter(inputArray any, cond any, opt *types.FilterOptions) bson.D {
	d := bson.D{{Key: types.AggregationInput, Value: inputArray}, {Key: types.AggregationCondWithoutOperator, Value: cond}}
	if opt != nil {
		if opt.As != "" {
			d = append(d, bson.E{Key: types.AggregationAs, Value: opt.As})
		}
		if opt.Limit != 0 {
			d = append(d, bson.E{Key: types.AggregationLimit, Value: opt.Limit})
		}
	}
	return bson.D{{Key: types.AggregationFilter, Value: d}}
}

func Eq(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationEq, Value: expressions}}
}

func Ne(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationNe, Value: expressions}}
}

func Gt(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationGt, Value: expressions}}
}

func Gte(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationGte, Value: expressions}}
}

func Lt(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationLt, Value: expressions}}
}

func Lte(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationLte, Value: expressions}}
}

func Cond(boolExpr, tureExpr, falseExpr any) bson.D {
	return bson.D{{Key: types.AggregationCond, Value: []any{boolExpr, tureExpr, falseExpr}}}
}

func IfNull(expr, replacement any) bson.D {
	return bson.D{{Key: types.AggregationIfNull, Value: []any{expr, replacement}}}
}

// Switch
// cases: [case, then, case, then]
func Switch(cases []any, defaultCase any) bson.D {
	if len(cases) != 0 && len(cases)%2 == 0 {
		branches := bson.A{}
		for i := 0; i < len(cases); i += 2 {
			branches = append(branches, bson.D{{Key: types.Case, Value: cases[i]}, {Key: types.Then, Value: cases[i+1]}})
		}
		return bson.D{bson.E{Key: types.AggregationSwitch, Value: bson.D{bson.E{Key: types.Branches, Value: branches}, bson.E{Key: types.DefaultCase, Value: defaultCase}}}}
	}
	return bson.D{}
}

func DateToString(date any, opt *types.DateToStringOptions) bson.D {
	d := bson.D{bson.E{Key: types.Date, Value: date}}
	if opt != nil {
		if opt.Format != "" {
			d = append(d, bson.E{Key: types.Format, Value: opt.Format})
		}
		if opt.Timezone != "" {
			d = append(d, bson.E{Key: types.Timezone, Value: opt.Timezone})
		}
		if opt.OnNull != nil {
			d = append(d, bson.E{Key: types.OnNull, Value: opt.OnNull})
		}
	}
	return bson.D{{Key: types.AggregationDateToString, Value: d}}
}

func DayOfMonth(date time.Time) bson.D {
	return bson.D{{Key: types.AggregationDayOfMonth, Value: date}}
}

func DayOfMonthWithTimezone(date time.Time, timezone string) bson.D {
	return bson.D{{Key: types.AggregationDayOfMonth, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}}}
}

func DayOfWeek(date time.Time) bson.D {
	return bson.D{{Key: types.AggregationDayOfWeek, Value: date}}
}

func DayOfWeekWithTimezone(date time.Time, timezone string) bson.D {
	return bson.D{{Key: types.AggregationDayOfWeek, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}}}
}

func DayOfYear(date time.Time) bson.D {
	return bson.D{{Key: types.AggregationDayOfYear, Value: date}}
}

func DayOfYearWithTimezone(date time.Time, timezone string) bson.D {
	return bson.D{{Key: types.AggregationDayOfYear, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}}}
}

func Year(date time.Time) bson.D {
	return bson.D{{Key: types.AggregationYear, Value: date}}
}

func YearWithTimezone(date time.Time, timezone string) bson.D {
	return bson.D{{Key: types.AggregationYear, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}}}
}

func Month(date time.Time) bson.D {
	return bson.D{{Key: types.AggregationMonth, Value: date}}
}

func MonthWithTimezone(date time.Time, timezone string) bson.D {
	return bson.D{{Key: types.AggregationMonth, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}}}
}

func Week(date time.Time) bson.D {
	return bson.D{{Key: types.AggregationWeek, Value: date}}
}

func WeekWithTimezone(date time.Time, timezone string) bson.D {
	return bson.D{{Key: types.AggregationWeek, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}}}
}

func And(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationAnd, Value: expressions}}
}

func Or(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationOr, Value: expressions}}
}

func Not(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationNot, Value: expressions}}
}

func Concat(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationConcat, Value: expressions}}
}

func SubstrBytes(stringExpression string, byteIndex int64, byteCount int64) bson.D {
	return bson.D{{Key: types.AggregationSubstrBytes, Value: []any{stringExpression, byteIndex, byteCount}}}
}

func ToLower(expression any) bson.D {
	return bson.D{{Key: types.AggregationToLower, Value: expression}}
}

func ToUpper(expression any) bson.D {
	return bson.D{{Key: types.AggregationToUpper, Value: expression}}
}

func Contact(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationContact, Value: expressions}}
}
