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

package hook

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DefaultModel interface {
	// DefaultId set and get
	DefaultId() primitive.ObjectID
	// DefaultCreatedAt set and get
	DefaultCreatedAt() time.Time
	// DefaultUpdatedAt set and get
	DefaultUpdatedAt() time.Time
}

type CustomModel interface {
	// CustomID set and get field name and value
	CustomID() (string, any)
	// CustomCreatedAt set and get field name and value
	CustomCreatedAt() (string, any)
	// CustomUpdatedAt set and get field name and value
	CustomUpdatedAt() (string, any)
}
