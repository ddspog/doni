package project

//go:generate go run gen_version.go

// Version returns the version of the project.
func Version() string {
	return version
}
