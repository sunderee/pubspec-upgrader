package main

import (
	"flag"
	"fmt"
)

const (
	reset = "\033[0m"
	red   = "\033[31m"
	green = "\033[32m"
)

func main() {
	filePath := flag.String("file", "", "Path to the pubspec.yaml file")
	ignoreUnstable := flag.Bool("ignore-unstable", true, "Whether unstable versions are ignored")
	flag.Parse()

	if len(*filePath) == 0 {
		fmt.Println(red + "Path to the pubspec.yaml file is required" + reset)
		flag.Usage()
	}

	pubspecYaml := parsePubspecYaml(*filePath)
	if pubspecYaml == nil {
		fmt.Println(red + "Something went wrong while parsing" + reset)
	}

	fmt.Println("Fetching results, please be patient...")
	if pubspecYaml != nil {
		pubspecUpgraderResult := requestDependenciesVersions(*pubspecYaml, *ignoreUnstable)
		fmt.Println("Results for dependencies:")
		for _, dependency := range pubspecUpgraderResult.Dependencies {
			if dependency.LocalVersion != dependency.RemoteVersion {
				fmt.Printf(green+"%s can be upgraded from %s to %s"+reset+"\n", dependency.Name, dependency.LocalVersion, dependency.RemoteVersion)
			} else {
				fmt.Printf("%s is up to date\n", dependency.Name)
			}
		}

		fmt.Println("Results for dev_dependencies:")
		for _, devDependency := range pubspecUpgraderResult.DevDependencies {
			if devDependency.LocalVersion != devDependency.RemoteVersion {
				fmt.Printf(green+"%s can be upgraded from %s to %s"+reset+"\n", devDependency.Name, devDependency.LocalVersion, devDependency.RemoteVersion)
			} else {
				fmt.Printf("%s is up to date\n", devDependency.Name)
			}
		}

		fmt.Println("Done!")
	}

}
