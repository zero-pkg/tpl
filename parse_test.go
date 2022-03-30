package tpl

import (
	"testing"
)

func TestParseFile(t *testing.T) {
	file, err := parseFile("testdata/", "testdata/standalone/standalone.html")
	ok(t, err)
	equals(t, "testdata/standalone/standalone.html", file.abspath)
	equals(t, "standalone/standalone.html", file.path)
	assert(t, file.parent == nil, "file.parent is not nil")
	assert(t, len(file.content) > 0, "file.content is zero len")
}

func TestParseFileExtended(t *testing.T) {
	file, err := parseFile("testdata/", "testdata/middle.html")
	ok(t, err)
	equals(t, "testdata/middle.html", file.abspath)
	equals(t, "middle.html", file.path)
	assert(t, file.parent != nil, "file.parent is nil")
	assert(t, *file.parent == "top.html", "file.parent is not equal top.html")
	assert(t, len(file.content) > 0, "file.content is zero len")
}

func TestParseFileNotFound(t *testing.T) {
	file, err := parseFile("testdata/", "testdata/foo.html")
	assert(t, err != nil, "err is nil")
	assert(t, file == nil, "file is not nil")
}
