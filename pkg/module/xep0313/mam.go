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

package xep0313

import (
	"context"

	"github.com/jackal-xmpp/stravaganza/v2"
	"github.com/ortuman/jackal/pkg/hook"
	"github.com/ortuman/jackal/pkg/log"
	"github.com/ortuman/jackal/pkg/repository"
	"github.com/ortuman/jackal/pkg/router"
)

const (
	mamNamespace = "urn:xmpp:mam:2"
)

const (
	// ModuleName represents MAM module name.
	ModuleName = "mam"

	// XEPNumber represents MAM XEP number.
	XEPNumber = "0313"
)

// Mam represents a MAM (XEP-0313) module type.
type Mam struct {
	rep    repository.VCard
	router router.Router
	hk     *hook.Hooks
}

// New returns a new initialized Mam instance.
func New(
	router router.Router,
	rep repository.Repository,
	hk *hook.Hooks,
) *Mam {
	return &Mam{
		router: router,
		rep:    rep,
		hk:     hk,
	}
}

// Name returns Mam module name.
func (m *Mam) Name() string { return ModuleName }

// StreamFeature returns Mam module stream feature.
func (m *Mam) StreamFeature(_ context.Context, _ string) (stravaganza.Element, error) {
	return nil, nil
}

// ServerFeatures returns Mam server disco features.
func (m *Mam) ServerFeatures(_ context.Context) ([]string, error) {
	return nil, nil
}

// AccountFeatures returns Mam account disco features.
func (m *Mam) AccountFeatures(_ context.Context) ([]string, error) {
	return []string{mamNamespace, mamNamespace + "extended"}, nil
}

// MatchesNamespace tells whether namespace matches Mam module.
func (m *Mam) MatchesNamespace(namespace string, _ bool) bool {
	return namespace == mamNamespace
}

// ProcessIQ process a Mam iq.
func (m *Mam) ProcessIQ(ctx context.Context, iq *stravaganza.IQ) error {
	return nil
}

// Start starts module.
func (m *Mam) Start(_ context.Context) error {
	m.hk.AddHook(hook.UserDeleted, m.onUserDeleted, hook.DefaultPriority)

	log.Infow("Started module", "xep", XEPNumber)
	return nil
}

// Stop stops Mam module.
func (m *Mam) Stop(_ context.Context) error {
	m.hk.RemoveHook(hook.UserDeleted, m.onUserDeleted)

	log.Infow("Stopped module", "xep", XEPNumber)
	return nil
}

func (m *Mam) onUserDeleted(ctx context.Context, execCtx *hook.ExecutionContext) error {
	// TODO(ortuman): implement me dude!
	return nil
}
