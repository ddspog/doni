// +build ignore

// This program generates version.go. It can be invoked by running
// go generate.
package main

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"
	"time"
)

var pathToVersion = os.ExpandEnv("$GOPATH/src/github.com/ddspog/doni/version")

var versionTemplate = parseTemplate(`// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// {{ .Timestamp }}
// using data from
// {{ .VersionPath }}
package project

const version = "{{ .VersionValue }}"
`)

func parseTemplate(t string) *template.Template {
	return template.Must(template.New("").Parse(t))
}

func readVersion() string {
	if c, err := ioutil.ReadFile(pathToVersion); err == nil {
		return string(c)
	}

	return ""
}

func main() {
	f, err := os.Create("version.go")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	versionTemplate.Execute(f, struct {
		Timestamp    time.Time
		VersionPath  string
		VersionValue string
	}{
		Timestamp:    time.Now(),
		VersionPath:  pathToVersion,
		VersionValue: readVersion(),
	})
}