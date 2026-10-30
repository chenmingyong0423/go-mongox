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

package operation

// SoftDeleteScope controls which documents a finder operation can match.
type SoftDeleteScope uint8

const (
	// SoftDeleteScopeNotDeleted is the default scope.
	SoftDeleteScopeNotDeleted SoftDeleteScope = iota
	// SoftDeleteScopeWithDeleted includes both live and deleted documents.
	SoftDeleteScopeWithDeleted
	// SoftDeleteScopeOnlyDeleted includes only deleted documents.
	SoftDeleteScopeOnlyDeleted
)
