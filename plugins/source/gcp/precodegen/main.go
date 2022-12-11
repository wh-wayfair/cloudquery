package main

import (
	"log"

	"github.com/cloudquery/plugins/source/gcp/precodegen/internal/parser"
)

func main() {
	pkgs, err := parser.DiscoverSubpackages()
	if err != nil {
		log.Fatal(err)
	}
	for _, pkg := range pkgs {
		log.Println(pkg)
	}
}