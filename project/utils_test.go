package project

import (
	"io/ioutil"
	"os"
	"strings"
)

var pathToProject = "$GOPATH/src/github.com/ddspog/doni/"

func projectFileAsString(f string) string {
	p := os.ExpandEnv(pathToProject) + f

	if c, err := ioutil.ReadFile(p); err == nil {
		return strings.TrimSuffix(string(c), "\r\n")
	}

	return ""
}
