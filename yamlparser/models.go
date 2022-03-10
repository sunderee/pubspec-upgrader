package yamlparser

type PubspecYamlFileModel struct {
	Dependencies    []DependencyModel
	DevDependencies []DependencyModel
}

type DependencyModel struct {
	Name    string
	Version string
}
