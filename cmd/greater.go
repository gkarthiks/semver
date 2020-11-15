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
	"github.com/gkarthiks/semver/utils"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// greaterCmd represents the greater command
var greaterCmd = &cobra.Command{
	Use:   "greater",
	Short: "Compare two given semver",
	Long: `Compares two given semver (A, B) and exists with 0 if A is greater than B, fatal errors if B is greater than A. 
    For example:
	$ semver greater 1.5.8 1.4.0
		
	will exits with 0.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		lhsSemver := args[0]
		rhsSemver := args[1]
		utils.VerbosePrintf("Obtained the LHS semver as %v", lhsSemver)
		utils.VerbosePrintf("Obtained the RHS semver as %v", rhsSemver)
		lhs, err := utils.MustParseVersion(lhsSemver)
		if err != nil {
			log.Fatalf("error occurred while parsing the given string %s to semver", lhs)
		}
		rhs, err := utils.MustParseVersion(rhsSemver)
		if err != nil {
			log.Fatalf("error occurred while parsing the given string %s to semver", rhs)
		}
		utils.VerbosePrintf("Executing the comparison")
		if lhs.GreaterThan(rhs) {
			os.Exit(0)
		} else {
			log.Fatal("Not Greater")
		}
	},
}

func init() {
	rootCmd.AddCommand(greaterCmd)
}
