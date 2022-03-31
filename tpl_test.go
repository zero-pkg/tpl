package tpl

import (
	"bytes"
	"strings"
	"testing"
)

func TestExecute(t *testing.T) {
	tmpl := New()

	_, err := tmpl.ParseDir("testdata", ".html")
	ok(t, err)

	cases := map[string]string{
		"top.html":                   "This is content of top.html",
		"middle.html":                "This is content of middle.html and content of partial.html",
		"bottom.html":                "This is content of bottom.html and content of partial.html",
		"standalone/standalone.html": "This is content of standalone.html and content of partial.html",
		"partials/empty.html":        "",
	}

	for k := range cases {
		t.Run(k, func(t *testing.T) {
			var b bytes.Buffer

			err = tmpl.Execute(&b, k, "")
			ok(t, err)
			equals(t, cases[k], strings.TrimSpace(b.String()))
		})
	}
}

func TestExecuteError(t *testing.T) {
	tmpl := New()

	_, err := tmpl.ParseDir("testdata", ".html")
	ok(t, err)

	var b bytes.Buffer

	err = tmpl.Execute(&b, "notfound", "")
	equals(t, "template notfound is not found", err.Error())
}

func TestLookup(t *testing.T) {
	tmpl := New()

	_, err := tmpl.ParseDir("testdata", ".html")
	ok(t, err)

	template := tmpl.Lookup("standalone/standalone.html")
	assert(t, template != nil, "template is nil")
}

func TestLookupNotFound(t *testing.T) {
	tmpl := New()

	_, err := tmpl.ParseDir("testdata", ".html")
	ok(t, err)

	template := tmpl.Lookup("notfound")
	assert(t, template == nil, "template is not nil")
}
