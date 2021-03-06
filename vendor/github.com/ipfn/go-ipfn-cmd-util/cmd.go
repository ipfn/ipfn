// Copyright © 2017-2018 Łukasz Kurowski. All Rights Reserved.
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

// Package cmdutil implements command line utilities.
package cmdutil

import (
	"github.com/spf13/cobra"

	"github.com/ipfn/go-ipfn-cmd-util/logger"
)

// WrapCommand - Wraps cobra command function with error handler.
func WrapCommand(fn func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if err := fn(cmd, args); err != nil {
			logger.Error(err)
		}
	}
}
