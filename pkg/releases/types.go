package releases

const releaseBaseURL = "https://releases.hashicorp.com"

// ReleaseProduct holds all releases for a product
type ReleaseProduct struct {
	Name     string
	Versions map[string]*ReleaseVersion
}

// ReleaseVersion holds all builds for a product version
type ReleaseVersion struct {
	Name      string
	Version   string
	Shasums   string
	Signature string `json:"shasums_signature"`
	Builds    []*ReleaseBuild
}

// ReleaseBuild holds information on a single product build
type ReleaseBuild struct {
	Name     string
	Version  string
	OS       string
	Arch     string
	Filename string
}
