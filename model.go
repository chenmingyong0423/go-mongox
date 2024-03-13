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

package mongox

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Model is a base struct which includes the following fields:
// - ID: the primary key of the document
// - CreatedAt: the time when the document was created
// - UpdatedAt: the time when the document was last updated
// It may be embedded into a struct to provide these fields.
// Example:
//
//	type User struct {
//		mongox.Model
//	}
type Model struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func (m *Model) DefaultId() {
	if m.ID.IsZero() {
		m.ID = primitive.NewObjectID()
	}
}

func (m *Model) DefaultCreatedAt() {
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now().Local()
	}
}

func (m *Model) DefaultUpdatedAt() {
	m.UpdatedAt = time.Now().Local()
}
