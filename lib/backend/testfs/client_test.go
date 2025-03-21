// Copyright (c) 2016-2019 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package testfs

import (
	"testing"

	"github.com/uber-go/tally"
	"github.com/uber/kraken/lib/backend/namepath"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestClientFactory(t *testing.T) {
	require := require.New(t)

	config := Config{
		Addr:     "test",
		NamePath: namepath.Identity,
	}
	f := factory{}
	_, err := f.Create(config, nil, tally.NoopScope, zap.NewNop().Sugar())
	require.NoError(err)
}
