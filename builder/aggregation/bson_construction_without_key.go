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

	"github.com/chenmingyong0423/go-mongox/types"
	"go.mongodb.org/mongo-driver/bson"
)

func SumWithoutKey(expression any) bson.D {
	return bson.D{{Key: types.AggregationSum, Value: expression}}
}

func PushWithoutKey(expression any) bson.D {
	return bson.D{{Key: types.AggregationPush, Value: expression}}
}

func AvgWithoutKey(expression any) bson.D {
	return bson.D{{Key: types.AggregationAvg, Value: expression}}
}

func FirstWithoutKey(expression any) bson.D {
	return bson.D{{Key: types.AggregationFirst, Value: expression}}
}

func LastWithoutKey(expression any) bson.D {
	return bson.D{{Key: types.AggregationLast, Value: expression}}
}

func MinWithoutKey(expression any) bson.D {
	return bson.D{{Key: types.AggregationMin, Value: expression}}
}

func MaxWithoutKey(expression any) bson.D {
	return bson.D{{Key: types.AggregationMax, Value: expression}}
}

func AddWithoutKey(expression ...any) bson.D {
	return bson.D{{Key: types.AggregationAdd, Value: expression}}
}

func MultiplyWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationMultiply, Value: expressions}}
}

func SubtractWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationSubtract, Value: expressions}}
}

func DivideWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationDivide, Value: expressions}}
}

func ModWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationMod, Value: expressions}}
}

func ArrayElemAtWithoutKey(expression any, index int64) bson.D {
	return bson.D{{Key: types.AggregationArrayElemAt, Value: []any{expression, index}}}
}

func ConcatArraysWithoutKey(arrays ...any) bson.D {
	return bson.D{{Key: types.AggregationConcatArrays, Value: arrays}}
}

func ArrayToObjectWithoutKey(expression any) bson.D {
	return bson.D{{Key: types.AggregationArrayToObject, Value: expression}}
}

func SizeWithoutKey(expression any) bson.D {
	return bson.D{{Key: types.AggregationSize, Value: expression}}
}

func SliceWithoutKey(array any, nElements int64) bson.D {
	return bson.D{{Key: types.AggregationSlice, Value: []any{array, nElements}}}
}

func SliceWithPositionWithoutKey(array any, position, nElements int64) bson.D {
	return bson.D{{Key: types.AggregationSlice, Value: []any{array, position, nElements}}}
}

func MapWithoutKey(inputArray any, as string, in any) bson.D {
	return bson.D{{Key: types.AggregationMap, Value: bson.D{
		{Key: types.AggregationInput, Value: inputArray},
		{Key: types.AggregationAs, Value: as},
		{Key: types.AggregationIn, Value: in},
	}}}
}

func FilterWithoutKey(inputArray any, cond any, opt *types.FilterOptions) bson.D {
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

func EqWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationEq, Value: expressions}}
}

func NeWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationNe, Value: expressions}}
}

func GtWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationGt, Value: expressions}}
}

func GteWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationGte, Value: expressions}}
}

func LtWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationLt, Value: expressions}}
}

func LteWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationLte, Value: expressions}}
}

func CondWithoutKey(boolExpr, tureExpr, falseExpr any) bson.D {
	return bson.D{{Key: types.AggregationCond, Value: []any{boolExpr, tureExpr, falseExpr}}}
}

func IfNullWithoutKey(expr, replacement any) bson.D {
	return bson.D{{Key: types.AggregationIfNull, Value: []any{expr, replacement}}}
}
func SwitchWithoutKey(cases []types.CaseThen, defaultCase any) bson.D {
	branches := bson.A{}
	for _, caseThen := range cases {
		branches = append(branches, bson.D{bson.E{Key: types.Case, Value: caseThen.Case}, {Key: types.Then, Value: caseThen.Then}})
	}
	return bson.D{bson.E{Key: types.AggregationSwitch, Value: bson.D{bson.E{Key: types.Branches, Value: branches}, bson.E{Key: types.DefaultCase, Value: defaultCase}}}}
}

func DateToStringWithoutKey(date any, opt *types.DateToStringOptions) bson.D {
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

func DayOfMonthWithoutKey(date time.Time) bson.D {
	return bson.D{{Key: types.AggregationDayOfMonth, Value: date}}
}

func DayOfMonthWithTimezoneWithoutKey(date time.Time, timezone string) bson.D {
	return bson.D{{Key: types.AggregationDayOfMonth, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}}}
}

func DayOfWeekWithoutKey(date time.Time) bson.D {
	return bson.D{{Key: types.AggregationDayOfWeek, Value: date}}
}

func DayOfWeekWithTimezoneWithoutKey(date time.Time, timezone string) bson.D {
	return bson.D{{Key: types.AggregationDayOfWeek, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}}}
}

func DayOfYearWithoutKey(date time.Time) bson.D {
	return bson.D{{Key: types.AggregationDayOfYear, Value: date}}
}

func DayOfYearWithTimezoneWithoutKey(date time.Time, timezone string) bson.D {
	return bson.D{{Key: types.AggregationDayOfYear, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}}}
}

func YearWithoutKey(date time.Time) bson.D {
	return bson.D{{Key: types.AggregationYear, Value: date}}
}

func YearWithTimezoneWithoutKey(date time.Time, timezone string) bson.D {
	return bson.D{{Key: types.AggregationYear, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}}}
}

func MonthWithoutKey(date time.Time) bson.D {
	return bson.D{{Key: types.AggregationMonth, Value: date}}
}

func MonthWithTimezoneWithoutKey(date time.Time, timezone string) bson.D {
	return bson.D{{Key: types.AggregationMonth, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}}}
}

func WeekWithoutKey(date time.Time) bson.D {
	return bson.D{{Key: types.AggregationWeek, Value: date}}
}

func WeekWithTimezoneWithoutKey(date time.Time, timezone string) bson.D {
	return bson.D{{Key: types.AggregationWeek, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}}}
}

func AndWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationAnd, Value: expressions}}
}

func OrWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationOr, Value: expressions}}
}

func NotWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationNot, Value: expressions}}
}

func ConcatWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationConcat, Value: expressions}}
}

func SubstrBytesWithoutKey(stringExpression string, byteIndex int64, byteCount int64) bson.D {
	return bson.D{{Key: types.AggregationSubstrBytes, Value: []any{stringExpression, byteIndex, byteCount}}}
}

func ToLowerWithoutKey(expression any) bson.D {
	return bson.D{{Key: types.AggregationToLower, Value: expression}}
}

func ToUpperWithoutKey(expression any) bson.D {
	return bson.D{{Key: types.AggregationToUpper, Value: expression}}
}

func ContactWithoutKey(expressions ...any) bson.D {
	return bson.D{{Key: types.AggregationContact, Value: expressions}}
}
