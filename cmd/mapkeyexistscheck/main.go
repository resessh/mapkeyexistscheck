package main

import (
	"github.com/resessh/mapkeyexistscheck"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(mapkeyexistscheck.Analyzer)
}
