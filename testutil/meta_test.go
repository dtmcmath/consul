// It's important to keep the example in README.md in sync with the
// test in readme_test.go so that compilation tests will keep the
// README in sync.  There's more than one way to do it.
//
// One way is to actually make the README include the test.  That's
// what https://github.com/dave/rebecca, for example, is about, but it
// requires an extra build step.
//
// Another way it to have _another_ test that checks the bytes and
// complains if they're not the same.  That's sort of crazy, except
// that it sort of works.  That's what this test does.

package testutil_test

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"testing"
)

func TestReadmeMatches(t *testing.T) {
	readme, err := ioutil.ReadFile("README.md")
	if err != nil {
		t.Fatal(err)
	}
	codeSectionRe := regexp.MustCompile("```go\n([^`]*)```")
	codeSections := codeSectionRe.FindAllStringSubmatch(fmt.Sprintf("%s", readme), -1)
	if len(codeSections) != 1 {
		t.Errorf("README.md has the wrong number of code blocks; expected %d, got %d",
			1, len(codeSections),
		)
	}
	readmeCodeLines := strings.Split(codeSections[0][1], "\n")

	testfile, err := ioutil.ReadFile("readme_test.go")
	if err != nil {
		t.Fatal(err)
	}
	testfileCodeLines := strings.Split(fmt.Sprintf("%s", testfile), "\n")

	if len(testfileCodeLines) != len(readmeCodeLines) {
		t.Errorf("Expected README.md code block to have %d lines, got %d",
			len(testfileCodeLines), len(readmeCodeLines),
		)
	}

	readmePackageLine := strings.Replace(
		readmeCodeLines[0],
		"my_program",
		"testutil_test",
		-1,
	)
	if readmePackageLine != testfileCodeLines[0] {
		t.Errorf("Mismatch in README.md code, line %d.", 0)
	}
	for lineNum := 1; lineNum < len(testfileCodeLines) && lineNum < len(readmeCodeLines); lineNum++ {
		if readmeCodeLines[lineNum] != testfileCodeLines[lineNum] {
			t.Errorf("Mismatch in README.md code, line %d.  Expected\n\t'%s'\nGot\n\t'%s'\n",
				lineNum, testfileCodeLines[lineNum], readmeCodeLines[lineNum],
			)
		}
	}
}
