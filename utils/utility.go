package utils

import (
	"errors"
	"fmt"
	"github.com/Masterminds/semver/v3"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

// MustParseVersion returns the parsed semver of the given string if no error
// else returns err
func MustParseVersion(versionString string) (*semver.Version, error) {
	VerbosePrintf("Parsing the given string %s to Semantic Version.\n", versionString)
	semParsedVersion, err := semver.NewVersion(versionString)
	if err != nil {
		return nil, err
	}
	return semParsedVersion, nil
}

// IncrementVersion will increment the given semver corresponding to major/minor/patch or alpha
func IncrementVersion(parsedSemver *semver.Version, incType string) (semver.Version, error) {
	switch incType {
	case "minor":
		VerbosePrintf("Obtained 'minor' as the incrementing argument")
		return parsedSemver.IncMinor(), nil
	case "major":
		VerbosePrintf("Obtained 'major' as the incrementing argument")
		return parsedSemver.IncMajor(), nil
	case "patch":
		VerbosePrintf("Obtained 'patch' as the incrementing argument")
		return parsedSemver.IncPatch(), nil
	case "alpha":
		VerbosePrintf("Obtained 'alpha' as the incrementing argument")
		preReleaseStr := parsedSemver.Prerelease()
		alphaSplits := strings.Split(preReleaseStr, ".")
		VerbosePrintf("String split happened and obtained the following values.\n %v\n", alphaSplits)
		if len(alphaSplits) > 1 {
			VerbosePrintf("alphaSplits is more than 1, proceeding with the string conversion.")

			versionInt, err := strconv.Atoi(alphaSplits[1])
			if err != nil {
				log.Fatalf("error occurred while converting the alpha number from string to integer, err: %v", err)
			}
			newIncVersion := versionInt + 1
			VerbosePrintf("New increment version is %v \n", newIncVersion)
			incAlphaVersion := alphaSplits[0]+"."+strconv.Itoa(newIncVersion)
			alphaIncVersion := fmt.Sprintf("%d.%d.%d-%s", parsedSemver.Major(), parsedSemver.Minor(), parsedSemver.Patch(), incAlphaVersion)
			VerbosePrintf("New concatenated complete version is %v \n", alphaIncVersion)
			parsedNewAlphaVersion, err := semver.NewVersion(alphaIncVersion)
			if err != nil {
				return *DefaultVersion, err
			}
			return *parsedNewAlphaVersion, nil
		}
		VerbosePrintf("Arguments is not greater than 1, so proceeding with basic alpha logic")
		preRelease, err := parsedSemver.SetPrerelease("alpha")
		if err != nil {
			return *DefaultVersion, err
		}

		VerbosePrintf("Parsed pre-release is %v", preRelease)
		return preRelease, nil

	default:
		VerbosePrintf("Executing command for the default case")
		err := errors.New("error in input, accepted inc inputs are major, minor, patch and alpha")
		return *DefaultVersion, err
	}
}

// VerbosePrintf will print the debug string message and arguments if the verbose flag is passed
func VerbosePrintf(printString string, args ...interface{}) {
	if Verbose {
		log.SetLevel(log.DebugLevel)
		log.Debugf(printString, args)
	}
}
