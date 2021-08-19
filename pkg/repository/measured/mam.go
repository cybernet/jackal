// Copyright 2021 The jackal Authors
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

package measuredrepository

import (
	"context"
	"github.com/ortuman/jackal/pkg/repository"
	"time"
)

type measuredMamRep struct {
	rep repository.Mam
}

func (m *measuredMamRep) DeleteArchive(ctx context.Context, archiveID string) error {
	t0 := time.Now()
	err := m.rep.DeleteArchive(ctx, archiveID)
	reportOpMetric(deleteOp, time.Since(t0).Seconds(), err == nil)
	return err
}
