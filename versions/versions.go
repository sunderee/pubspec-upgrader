package versions

import (
	"fmt"
	"strings"
)

func GetLatestPackageVersion(packageName string, ignoreUnstable bool) string {
	packageInfo, requestError := requestPackageInfo(packageName)
	if requestError != nil {
		fmt.Print(requestError.Error())
		return ""
	}

	var versions []string = packageInfo.PackageVersions
	if ignoreUnstable {
		return versions[0]
	}

	for _, version := range versions {
		if !strings.Contains(version, "beta") && !strings.Contains(version, "alpha") {
			return version
		}
	}
	return ""
}
