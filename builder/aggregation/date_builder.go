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

type dateBuilder struct {
	parent *Builder
}

func (b *dateBuilder) DateToString(date any, opt *types.DateToStringOptions) *Builder {
	d := bson.D{bson.E{Key: types.Date, Value: date}}
	if opt.Format != "" {
		d = append(d, bson.E{Key: types.Format, Value: opt.Format})
	}
	if opt.Timezone != "" {
		d = append(d, bson.E{Key: types.Timezone, Value: opt.Timezone})
	}
	if opt.OnNull != nil {
		d = append(d, bson.E{Key: types.OnNull, Value: opt.OnNull})
	}
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationDateToString, Value: d})
	return b.parent
}

func (b *dateBuilder) DayOfMonth(date time.Time) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationDayOfMonth, Value: date})
	return b.parent
}

func (b *dateBuilder) DayOfMonthWithTimezone(date time.Time, timezone string) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationDayOfMonth, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}})
	return b.parent
}

func (b *dateBuilder) DayOfWeek(date time.Time) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationDayOfWeek, Value: date})
	return b.parent
}

func (b *dateBuilder) DayOfWeekWithTimezone(date time.Time, timezone string) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationDayOfWeek, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}})
	return b.parent
}

func (b *dateBuilder) DayOfYear(date time.Time) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationDayOfYear, Value: date})
	return b.parent
}

func (b *dateBuilder) DayOfYearWithTimezone(date time.Time, timezone string) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationDayOfYear, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}})
	return b.parent
}

func (b *dateBuilder) Year(date time.Time) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationYear, Value: date})
	return b.parent
}

func (b *dateBuilder) YearWithTimezone(date time.Time, timezone string) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationYear, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}})
	return b.parent
}

func (b *dateBuilder) Month(date time.Time) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationMonth, Value: date})
	return b.parent
}

func (b *dateBuilder) MonthWithTimezone(date time.Time, timezone string) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationMonth, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}})
	return b.parent
}

func (b *dateBuilder) Week(date time.Time) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationWeek, Value: date})
	return b.parent
}

func (b *dateBuilder) WeekWithTimezone(date time.Time, timezone string) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: types.AggregationWeek, Value: bson.D{bson.E{Key: types.Date, Value: date}, bson.E{Key: types.Timezone, Value: timezone}}})
	return b.parent
}
