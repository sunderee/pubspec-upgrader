package yamlparser

func ParsePubspecDependencies(pubspecYamlFilePath string) (*PubspecYamlFileModel, error) {
	rawYamlFile, yamlParsingError := parseRawPubspecModel(pubspecYamlFilePath)
	if yamlParsingError != nil {
		return nil, yamlParsingError
	}

	rawDependencies := rawYamlFile["dependencies"].(map[interface{}]interface{})
	rawDevDependencies := rawYamlFile["dev_dependencies"].(map[interface{}]interface{})

	return &PubspecYamlFileModel{
		Dependencies:    parseDependencies(rawDependencies),
		DevDependencies: parseDevDependencies(rawDevDependencies),
	}, nil
}
