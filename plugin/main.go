package main

import (
	"github.com/resessh/mapkeyexistscheck"
	"golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

func (a analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		mapkeyexistscheck.Analyzer,
	}
}

// nolint
var AnalyzerPlugin analyzerPlugin
