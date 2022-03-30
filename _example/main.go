package main

import (
	"os"

	"github.com/zero-pkg/tpl"
)

func main() {
	tmpl := tpl.Must(tpl.New().ParseDir("templates", ".html"))

	if err := tmpl.Execute(os.Stdout, "content.html", ""); err != nil {
		panic(err)
	}
}
