package main

// staticlint app allows running checks on the source code.
// The following checks will run:
// `shadow.Analyzer` - checks for shadowed variables.
// `printf.Analyzer` - checks consistency of Printf format strings and arguments.
// `structtag.Analyzer` - checks struct field tags are well formed.
// `assign.Analyzer` - detects useless assignments.
// `analyzers.OsExitCheck` - detects using os.Exit() in the main package.
import (
	_ "embed"
	"encoding/json"

	"github.com/cliffordsimak-76-cards/url-shortener/cmd/staticlint/customanalyzer"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/assign"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"honnef.co/go/tools/staticcheck"
)

type ConfigData struct {
	Staticchecks []string
}

//go:embed config.json
var data []byte

func main() {
	analyzers := []*analysis.Analyzer{
		shadow.Analyzer,
		printf.Analyzer,
		structtag.Analyzer,
		assign.Analyzer,
		customanalyzer.OsExitCheck,
	}

	analyzers = appendStaticCheckAnylyzers(analyzers)

	multichecker.Main(
		analyzers...,
	)
}

func appendStaticCheckAnylyzers(analyzers []*analysis.Analyzer) []*analysis.Analyzer {
	var cfg ConfigData
	if err := json.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	checks := make(map[string]bool)
	for _, c := range cfg.Staticchecks {
		checks[c] = true
	}

	for _, a := range staticcheck.Analyzers {
		if checks[a.Name] {
			analyzers = append(analyzers, a)
		}
	}

	return analyzers
}
