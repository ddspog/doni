package experiment

import (
	"fmt"
	"sort"
	"text/tabwriter"
)

// PrintExperimentsHelp prints help os experiments that have a description
func (e *Executor) PrintExperimentsHelp() {
	experiments := e.experimentsWithDesc()
	if len(experiments) == 0 {
		e.outf("experiment: No experiments with description available")
		return
	}
	e.outf("experiment: Available experiments for this project:")

	// Format in tab-separated columns with a tab stop of 8.
	w := tabwriter.NewWriter(e.Stdout, 0, 8, 0, '\t', 0)
	for _, experiment := range experiments {
		fmt.Fprintln(w, fmt.Sprintf("* %s: \t%s", experiment, e.Experiments[experiment].Desc))
	}
	w.Flush()
}

func (e *Executor) experimentsWithDesc() (experiments []string) {
	for name, experiment := range e.Experiments {
		if experiment.Desc != "" {
			experiments = append(experiments, name)
		}
	}
	sort.Strings(experiments)
	return
}
