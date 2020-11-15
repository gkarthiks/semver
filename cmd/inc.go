/*
Package cmd hold the command logic for equal, greater, inc and lesser
*/
/*
Copyright © 2020 Karthikeyan Govindaraj <github.gkarthiks@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"semver/utils"
)

// incCmd represents the inc command
var incCmd = &cobra.Command{
	Use:   "inc",
	Short: "Increments the semver against the given flag",
	Long: `Increments the given semver corresponding to the provided flag.
    For example:

    $ semver inc 1.5.11 minor

will increments the minor version and results in 1.6.0
Similarly

	$ semver inc 1.1.8 patch

will increments the patch version and results in 1.1.9
Can also increment and drop an alpha tag, like

	$ semver inc 1.6.9 minor alpha

will results in 1.7.0-alpha
.`,
	Args: cobra.RangeArgs(2, 3),
	Run: func(cmd *cobra.Command, args []string) {

		parsedSemver, err := utils.MustParseVersion(args[0])
		if err != nil {
			log.Fatalf("error while parsing the given string %s to semver", args[0])
		}
		utils.VerbosePrintf("Parsed given semver string as %v", parsedSemver)
		if len(args) == 3 {
			utils.VerbosePrintf("Total arguments obtained is 3, proceeding with 3 args logic")
			if args[2] == "alpha" {
				if incVersion, err := utils.IncrementVersion(parsedSemver, args[1]); err != nil {
					log.Fatalf("error occurred while executing increment version, err: %v", err)
				} else {
					finalIncVersion, err := utils.IncrementVersion(&incVersion, "alpha")
					if err != nil {
						log.Fatalf("error occurred while doing incrementing and adding alpha %v", err)
					} else {
						fmt.Println(finalIncVersion)
					}
				}
			} else {
				log.Fatalf("the third argument must be rc followed by major/minor/patch")
			}
		} else if incVersion, err := utils.IncrementVersion(parsedSemver, args[1]); err != nil {
			log.Fatalf("error occurred while executing increment version, err: %v", err)
		} else {
			fmt.Println(incVersion)
		}
	},
}

func init() {
	rootCmd.AddCommand(incCmd)
}
