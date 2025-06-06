// Copyright © 2025 Attestant Limited.
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

package api

import "github.com/attestantio/go-eth2-client/spec/phase0"

// ValidatorLivenessOpts are the options for obtaining validator liveness information.
type ValidatorLivenessOpts struct {
	Common CommonOpts

	// Epoch is the epoch for which the data is obtained.
	Epoch phase0.Epoch
	// Indices is a list of validators for which to obtain the duties.
	Indices []phase0.ValidatorIndex
}
