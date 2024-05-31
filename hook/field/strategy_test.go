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

package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBeforeInsert(t *testing.T) {
	testCases := []struct {
		name string
		doc  any

		wantErr      error
		validateFunc func(*testing.T, *model)
	}{
		{
			name:    "nil document",
			doc:     nil,
			wantErr: nil,
		},
		{
			name:    "the type not implement DefaultModelHook",
			doc:     struct{}{},
			wantErr: nil,
		},
		{
			name:    "default id and created at",
			doc:     &model{},
			wantErr: nil,
			validateFunc: func(t *testing.T, m *model) {
				assert.NotZero(t, m.ID)
				assert.NotZero(t, m.CreatedAt)
				assert.Zero(t, m.UpdatedAt)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := BeforeInsert(tc.doc)
			assert.Equal(t, tc.wantErr, err)
			if tc.validateFunc != nil {
				tc.validateFunc(t, tc.doc.(*model))
			}
		})
	}
}

func TestBeforeUpdate(t *testing.T) {
	testCases := []struct {
		name string
		doc  any

		wantErr      error
		validateFunc func(*testing.T, *model)
	}{
		{
			name:    "nil document",
			doc:     nil,
			wantErr: nil,
		},
		{
			name:    "the type not implement DefaultModelHook",
			doc:     struct{}{},
			wantErr: nil,
		},
		{
			name:    "default updated at",
			doc:     &model{},
			wantErr: nil,
			validateFunc: func(t *testing.T, m *model) {
				assert.Zero(t, m.ID)
				assert.Zero(t, m.CreatedAt)
				assert.NotZero(t, m.UpdatedAt)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := BeforeUpdate(tc.doc)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func TestBeforeUpsert(t *testing.T) {
	testCases := []struct {
		name string
		doc  any

		wantErr      error
		validateFunc func(*testing.T, *model)
	}{
		{
			name:    "nil document",
			doc:     nil,
			wantErr: nil,
		},
		{
			name:    "the type not implement DefaultModelHook",
			doc:     struct{}{},
			wantErr: nil,
		},
		{
			name:    "default id, created at and updated at",
			doc:     &model{},
			wantErr: nil,
			validateFunc: func(t *testing.T, m *model) {
				assert.NotZero(t, m.ID)
				assert.NotZero(t, m.CreatedAt)
				assert.NotZero(t, m.UpdatedAt)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := BeforeUpsert(tc.doc)
			assert.Equal(t, tc.wantErr, err)
			if tc.validateFunc != nil {
				tc.validateFunc(t, tc.doc.(*model))
			}
		})
	}
}
