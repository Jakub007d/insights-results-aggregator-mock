// Copyright 2020 Red Hat, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"github.com/rs/zerolog/log"
)

// responseDataError is used as the error message when the responses functions return an error
const responseDataError = "Unexpected error during response data encoding"

// AuthenticationError happens during auth problems, for example malformed token
type AuthenticationError struct {
	errString string
}

// handleServerError handles separate server errors and sends appropriate responses
func handleServerError(err error) {
	log.Error().Err(err).Msg("handleServerError()")
}
