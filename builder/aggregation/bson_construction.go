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

func Sum(key string, expression any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationSum, Value: expression}}}}
}

func Push(key string, expression any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationPush, Value: expression}}}}
}

func Avg(key string, expression any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationAvg, Value: expression}}}}
}

func First(key string, expression any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationFirst, Value: expression}}}}
}

func Last(key string, expression any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationLast, Value: expression}}}}
}

func Min(key string, expression any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationMin, Value: expression}}}}
}

func Max(key string, expression any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationMax, Value: expression}}}}
}

func Add(key string, expression ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationAdd, Value: expression}}}}
}

func Multiply(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationMultiply, Value: expressions}}}}
}

func Subtract(key string, expression ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationSubtract, Value: expression}}}}
}

func Divide(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationDivide, Value: expressions}}}}
}

func Mod(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationMod, Value: expressions}}}}
}

func ArrayElemAt(key string, expression any, index int64) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationArrayElemAt, Value: []any{expression, index}}}}}
}

func ConcatArrays(key string, arrays ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationConcatArrays, Value: arrays}}}}
}

func ArrayToObject(key string, expression any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationArrayToObject, Value: expression}}}}
}

func Size(key string, expression any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationSize, Value: expression}}}}
}

func Slice(key string, array any, nElements int64) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationSlice, Value: []any{array, nElements}}}}}
}

func SliceWithPosition(key string, array any, position, nElements int64) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationSlice, Value: []any{array, position, nElements}}}}}
}

func Map(key string, inputArray any, as string, in any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationMap, Value: bson.D{bson.E{Key: types.AggregationInput, Value: inputArray}, {Key: types.AggregationAs, Value: as}, {Key: types.AggregationIn, Value: in}}}}}}
}

func Filter(key string, inputArray any, cond any, opt *types.FilterOptions) bson.D {
	d := bson.D{bson.E{Key: types.AggregationInput, Value: inputArray}, {Key: types.AggregationCondWithoutOperator, Value: cond}}
	if opt != nil {
		if opt.As != "" {
			d = append(d, bson.E{Key: types.AggregationAs, Value: opt.As})
		}
		if opt.Limit != 0 {
			d = append(d, bson.E{Key: types.AggregationLimit, Value: opt.Limit})
		}
	}
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationFilter, Value: d}}}}
}

func Eq(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationEq, Value: expressions}}}}
}

func Ne(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationNe, Value: expressions}}}}
}

func Gt(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationGt, Value: expressions}}}}
}

func Gte(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationGte, Value: expressions}}}}
}

func Lt(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationLt, Value: expressions}}}}
}

func Lte(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationLte, Value: expressions}}}}
}

func Cond(key string, boolExpr, tureExpr, falseExpr any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationCond, Value: []any{boolExpr, tureExpr, falseExpr}}}}}
}

func IfNull(key string, expr, replacement any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationIfNull, Value: []any{expr, replacement}}}}}
}

func Switch(key string, cases []types.CaseThen, defaultCase any) bson.D {
	branches := bson.A{}
	for _, caseThen := range cases {
		branches = append(branches, bson.D{bson.E{Key: types.Case, Value: caseThen.Case}, {Key: types.Then, Value: caseThen.Then}})
	}
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationSwitch, Value: bson.D{bson.E{Key: types.Branches, Value: branches}, bson.E{Key: types.DefaultCase, Value: defaultCase}}}}}}
}

func DateToString(key string, date any, opt *types.DateToStringOptions) bson.D {
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
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationDateToString, Value: d}}}}
}

func DayOfMonth(key string, date time.Time) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationDayOfMonth, Value: date}}}}
}

func DayOfMonthWithTimezone(key string, date time.Time, timezone string) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationDayOfMonth, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}}}}}
}

func DayOfWeek(key string, date time.Time) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationDayOfWeek, Value: date}}}}
}

func DayOfWeekWithTimezone(key string, date time.Time, timezone string) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationDayOfWeek, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}}}}}
}

func DayOfYear(key string, date time.Time) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationDayOfYear, Value: date}}}}
}

func DayOfYearWithTimezone(key string, date time.Time, timezone string) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationDayOfYear, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}}}}}
}

func Year(key string, date time.Time) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationYear, Value: date}}}}
}

func YearWithTimezone(key string, date time.Time, timezone string) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationYear, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}}}}}
}

func Month(key string, date time.Time) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationMonth, Value: date}}}}
}

func MonthWithTimezone(key string, date time.Time, timezone string) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationMonth, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}}}}}
}

func Week(key string, date time.Time) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationWeek, Value: date}}}}
}

func WeekWithTimezone(key string, date time.Time, timezone string) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationWeek, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}}}}}
}

func And(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationAnd, Value: expressions}}}}
}

func Or(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationOr, Value: expressions}}}}
}

func Not(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationNot, Value: expressions}}}}
}

func Concat(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationConcat, Value: expressions}}}}
}

func SubstrBytes(key string, stringExpression string, byteIndex int64, byteCount int64) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationSubstrBytes, Value: []any{stringExpression, byteIndex, byteCount}}}}}
}

func ToLower(key string, expression any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationToLower, Value: expression}}}}
}

func ToUpper(key string, expression any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationToUpper, Value: expression}}}}
}

func Contact(key string, expressions ...any) bson.D {
	return bson.D{bson.E{Key: key, Value: bson.D{bson.E{Key: types.AggregationContact, Value: expressions}}}}
}
