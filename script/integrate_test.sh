# Copyright 2023 chenmingyong0423

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -e
docker-compose -f script/integration_test_compose.yml down -v
docker-compose -f script/integration_test_compose.yml up -d

#go test ./... -race -cover -failfast -count=1 -parallel=1 -tags=e2e
go list ./... | xargs -I {} sh -c 'go test -race -cover -failfast -count=1 -parallel=1 -tags=e2e {} || exit 255'
docker-compose -f script/integration_test_compose.yml down -v