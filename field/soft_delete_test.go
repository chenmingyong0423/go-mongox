// Copyright 2025 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package field

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestParseDeletedAt(t *testing.T) {
	tests := []struct {
		name      string
		doc       any
		key       string
		precision TimeType
	}{
		{"time", struct{ DeletedAt time.Time }{}, "deletedat", UnixTime},
		{"int", struct{ DeletedAt int }{}, "deletedat", UnixSecond},
		{"int64", struct {
			DeletedAt int64 `bson:"deleted_at,omitempty"`
		}{}, "deleted_at", UnixSecond},
		{"tagged time", struct {
			Removed time.Time `bson:"removed" mongox:"autoDeleteTime"`
		}{}, "removed", UnixTime},
		{"second", struct {
			Removed int64 `mongox:"autoDeleteTime:second"`
		}{}, "removed", UnixSecond},
		{"milli", struct {
			DeletedAt int64 `mongox:"autoDeleteTime:milli"`
		}{}, "deletedat", UnixMillisecond},
		{"nano", struct {
			Removed int `mongox:"autoDeleteTime:nano"`
		}{}, "removed", UnixNanosecond},
		{"unsupported", struct{ DeletedAt string }{}, "deletedat", 0},
		{"invalid precision", struct {
			DeletedAt time.Time `mongox:"autoDeleteTime:milli"`
		}{}, "deletedat", 0},
		{"unknown tag", struct {
			DeletedAt int64 `mongox:"autoDeleteTime:unknown"`
		}{}, "deletedat", 0},
		{"disabled", struct {
			DeletedAt time.Time `mongox:"-"`
		}{}, "deletedat", 0},
		{"ignored", struct {
			DeletedAt time.Time `bson:"-"`
		}{}, "-", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := ParseFields(tt.doc)
			require.Len(t, fields, 1)
			require.Equal(t, tt.key, fields[0].MongoField)
			require.Equal(t, tt.precision, fields[0].AutoDeleteTime)
			if tt.precision != 0 {
				encoded, err := bson.Marshal(tt.doc)
				require.NoError(t, err)
				// Verify the default name against the real driver codec.
				if tt.key != "deleted_at" {
					require.NotZero(t, bson.Raw(encoded).Lookup(tt.key).Type)
				}
			}
		})
	}
}

func TestParseInlineDeletedAt(t *testing.T) {
	type Audit struct {
		DeletedAt int64 `bson:"removed_at" mongox:"autoDeleteTime:milli"`
	}
	type Model struct {
		Audit `bson:",inline"`
	}
	type Doc struct {
		Model `bson:",inline"`
	}
	fields := ParseFields(new(Doc))
	deleted := fields[0].InlinedFields[0].InlinedFields[0]
	require.Equal(t, "removed_at", deleted.MongoField)
	require.Equal(t, UnixMillisecond, deleted.AutoDeleteTime)
}
