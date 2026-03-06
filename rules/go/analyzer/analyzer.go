package analyzer

import (
	"fmt"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/modernize"
)

var analyzers = modernize.Suite

var (
	// Set via x_defs in the Bazel target for each exported analyzer wrapper.
	name = "dummy value please replace using x_defs"

	Analyzer = findAnalyzerByName(name)
)

func findAnalyzerByName(name string) *analysis.Analyzer {
	for _, analyzer := range analyzers {
		if analyzer.Name == name {
			return analyzer
		}
	}
	panic(fmt.Sprintf("not a valid modernize analyzer: %s", name))
}
