package yamlparser

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

func parseRawPubspecModel(pubspecYamlFilePath string) (map[string]interface{}, error) {
	absoluteFilePath, absolutePathParsingError := filepath.Abs(pubspecYamlFilePath)
	if absolutePathParsingError != nil {
		return nil, absolutePathParsingError
	}

	rawYamlFile, readYamlFileError := ioutil.ReadFile(absoluteFilePath)
	if readYamlFileError != nil {
		return nil, readYamlFileError
	}

	var deserializationMap map[string]interface{}
	yamlFileUnmarshalError := yaml.Unmarshal(rawYamlFile, &deserializationMap)
	if yamlFileUnmarshalError != nil {
		return nil, yamlFileUnmarshalError
	}

	return deserializationMap, nil
}

func parseDependencies(rawDependencies map[interface{}]interface{}) []DependencyModel {
	var dependenciesList []DependencyModel
	for key, value := range rawDependencies {
		if key == "flutter" {
			continue
		}

		versionValue := fmt.Sprintf("%v", value)
		if strings.HasPrefix(versionValue, "^") {
			versionValue = strings.Replace(versionValue, "^", "", 1)
		}
		dependenciesList = append(dependenciesList, DependencyModel{
			Name:    fmt.Sprintf("%v", key),
			Version: versionValue,
		})
	}

	return dependenciesList
}

func parseDevDependencies(rawDevDependencies map[interface{}]interface{}) []DependencyModel {
	var devDependenciesList []DependencyModel
	for key, value := range rawDevDependencies {
		if key == "flutter_test" || key == "integration_test" {
			continue
		}

		versionValue := fmt.Sprintf("%v", value)
		if strings.HasPrefix(versionValue, "^") {
			versionValue = strings.Replace(versionValue, "^", "", 1)
		}
		devDependenciesList = append(devDependenciesList, DependencyModel{
			Name:    fmt.Sprintf("%v", key),
			Version: versionValue,
		})
	}

	return devDependenciesList
}
