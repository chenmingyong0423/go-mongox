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

type dateBuilder struct {
	parent *Builder
}

func (b *dateBuilder) DateToString(key string, date any, opt *DateToStringOptions) *Builder {
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
	e := bson.E{Key: DateToStringOp, Value: d}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *dateBuilder) DateToStringWithoutKey(date any, opt *DateToStringOptions) *Builder {
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
	b.parent.d = append(b.parent.d, bson.E{Key: DateToStringOp, Value: d})
	return b.parent
}

func (b *dateBuilder) DayOfMonth(key string, date time.Time) *Builder {
	e := bson.E{Key: DayOfMonthOp, Value: date}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *dateBuilder) DayOfMonthWithoutKey(date time.Time) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: DayOfMonthOp, Value: date})
	return b.parent
}

func (b *dateBuilder) DayOfMonthWithTimezone(key string, date time.Time, timezone string) *Builder {
	e := bson.E{Key: DayOfMonthOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *dateBuilder) DayOfMonthWithTimezoneWithoutKey(date time.Time, timezone string) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: DayOfMonthOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}})
	return b.parent
}

func (b *dateBuilder) DayOfWeek(key string, date time.Time) *Builder {
	e := bson.E{Key: DayOfWeekOp, Value: date}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *dateBuilder) DayOfWeekWithoutKey(date time.Time) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: DayOfWeekOp, Value: date})
	return b.parent
}

func (b *dateBuilder) DayOfWeekWithTimezone(key string, date time.Time, timezone string) *Builder {
	e := bson.E{Key: DayOfWeekOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *dateBuilder) DayOfWeekWithTimezoneWithoutKey(date time.Time, timezone string) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: DayOfWeekOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}})
	return b.parent
}

func (b *dateBuilder) DayOfYear(key string, date time.Time) *Builder {
	e := bson.E{Key: DayOfYearOp, Value: date}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *dateBuilder) DayOfYearWithoutKey(date time.Time) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: DayOfYearOp, Value: date})
	return b.parent
}

func (b *dateBuilder) DayOfYearWithTimezone(key string, date time.Time, timezone string) *Builder {
	e := bson.E{Key: DayOfYearOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *dateBuilder) DayOfYearWithTimezoneWithoutKey(date time.Time, timezone string) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: DayOfYearOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}})
	return b.parent
}

func (b *dateBuilder) Year(key string, date time.Time) *Builder {
	e := bson.E{Key: YearOp, Value: date}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *dateBuilder) YearWithoutKey(date time.Time) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: YearOp, Value: date})
	return b.parent
}

func (b *dateBuilder) YearWithTimezone(key string, date time.Time, timezone string) *Builder {
	e := bson.E{Key: YearOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *dateBuilder) YearWithTimezoneWithoutKey(date time.Time, timezone string) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: YearOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}})
	return b.parent
}

func (b *dateBuilder) Month(key string, date time.Time) *Builder {
	e := bson.E{Key: MonthOp, Value: date}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *dateBuilder) MonthWithoutKey(date time.Time) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: MonthOp, Value: date})
	return b.parent
}

func (b *dateBuilder) MonthWithTimezone(key string, date time.Time, timezone string) *Builder {
	e := bson.E{Key: MonthOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *dateBuilder) MonthWithTimezoneWithoutKey(date time.Time, timezone string) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: MonthOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}})
	return b.parent
}

func (b *dateBuilder) Week(key string, date time.Time) *Builder {
	e := bson.E{Key: WeekOp, Value: date}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *dateBuilder) WeekWithoutKey(date time.Time) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: WeekOp, Value: date})
	return b.parent
}

func (b *dateBuilder) WeekWithTimezone(key string, date time.Time, timezone string) *Builder {
	e := bson.E{Key: WeekOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}}
	if !b.parent.tryMergeValue(key, e) {
		b.parent.d = append(b.parent.d, bson.E{Key: key, Value: bson.D{e}})
	}
	return b.parent
}

func (b *dateBuilder) WeekWithTimezoneWithoutKey(date time.Time, timezone string) *Builder {
	b.parent.d = append(b.parent.d, bson.E{Key: WeekOp, Value: bson.D{bson.E{Key: DateOp, Value: date}, bson.E{Key: TimezoneOp, Value: timezone}}})
	return b.parent
}
