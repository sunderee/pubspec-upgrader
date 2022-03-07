package versions

type packageInfoModel struct {
	PackageName     string   `json:"name"`
	PackageVersions []string `json:"versions"`
}

type apiException struct {
	StatusCode int
	Message    string
}
