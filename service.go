package main

import (
	"github.com/sunderee/pubspec-upgrader/versions"
	"github.com/sunderee/pubspec-upgrader/yamlparser"
)

type PubspecUpgraderResultModel struct {
	Dependencies    []DependencyInfoModel
	DevDependencies []DependencyInfoModel
}

type DependencyInfoModel struct {
	Name          string
	LocalVersion  string
	RemoteVersion string
}

func parsePubspecYaml(filePath string) *yamlparser.PubspecYamlFileModel {
	pubspecYamlFile, yamlParsingError := yamlparser.ParsePubspecDependencies(filePath)
	if yamlParsingError != nil {
		panic(yamlParsingError)
	}

	return pubspecYamlFile
}

func requestDependenciesVersions(pubspecYaml yamlparser.PubspecYamlFileModel, ignoreUnstable bool) *PubspecUpgraderResultModel {
	var dependenciesList []DependencyInfoModel
	for _, dependency := range pubspecYaml.Dependencies {
		latestVersion := versions.GetLatestPackageVersion(dependency.Name, ignoreUnstable)
		dependenciesList = append(dependenciesList, DependencyInfoModel{
			Name:          dependency.Name,
			LocalVersion:  dependency.Version,
			RemoteVersion: latestVersion,
		})
	}

	var devDependenciesList []DependencyInfoModel
	for _, devDependency := range pubspecYaml.DevDependencies {
		latestVersion := versions.GetLatestPackageVersion(devDependency.Name, ignoreUnstable)
		devDependenciesList = append(devDependenciesList, DependencyInfoModel{
			Name:          devDependency.Name,
			LocalVersion:  devDependency.Version,
			RemoteVersion: latestVersion,
		})
	}

	return &PubspecUpgraderResultModel{
		Dependencies:    dependenciesList,
		DevDependencies: devDependenciesList,
	}
}
