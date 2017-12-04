package project

import (
	"testing"

	"github.com/ddspog/mspec/bdd"
)

// Feature Print version of project
// - As a developer,
// - I want to consult the version of the project,
// - So that I can inform users on command.
func TestVersion(t *testing.T) {
	given, like, s := bdd.Sentences()

	given(t, "project has version v%[1]s", func(when bdd.When, args ...interface{}) {
		when("Version() is called", func(it bdd.It) {
			val := Version()

			it("should return %[1]s", func(assert bdd.Assert) {
				assert.Equal(args[0].(string), val)
			})
		})
	}, like(
		s(projectFileAsString("version")),
	))
}
