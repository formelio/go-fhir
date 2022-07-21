// Copyright 2019 - 2021 The Samply Community
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

package cmd

import (
	"regexp"

	"github.com/spf13/cobra"
)

var namePattern = regexp.MustCompile("^[A-Z]([A-Za-z0-9_]){0,254}$")

var context = newContext()

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates Go structs from FHIR resource structure definitions.",
	Run: func(cmd *cobra.Command, args []string) {

		// Walk the directory and process all JSON FHIR definitions into context.Resources.
		readSourcesToContext(args[0])

		// Generate All Resources and find the elements that are required.
		generateResources()

		// Generate Elements that were added to context.RequiredElements during the generation of Resources
		generateElements(make(map[string]bool), context.RequiredElements)

		// Generate ValueSets that were added to context.RequiredValueSetBindings during the generation of other structs.
		generateValueSets()

		// Generate ResourceTypeToIResourceMap, this will be used when unmarshalling fields with type "Resource".
		generateResourceTypeToIResourceMap()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
