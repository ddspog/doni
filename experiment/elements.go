package experiment

// Experiment stores various values to configure an experiment where
// commands are called, and metrics are collected through various
// phases of execution.
type Experiment struct {
	Experiment string
	Desc       string
	Data       Data
	Collect    map[string]*Collect
	Phases     []*Phase
	Result     []map[string]interface{}
	Iteration  int
	Vars       map[string]interface{}
}

// Data stores values to be used as parameters during tests execution.
// It stores values in a matrix with defined number of rows and
// columns.
type Data struct {
	Columns []string
	Values  [][]interface{}
}

// Collect represents methods of collecting metrics along experiment.
type Collect struct {
	Collect    string
	File       string
	ParsedFile string
	Param      string
}

// Phase represents stages for the experiment executed.
type Phase struct {
	Phase string
	Desc  string
	Cmds  []*Cmd
}

// Cmd is a phase command
type Cmd struct {
	Cmd       string
	ParsedCmd string
	Collect   []*string
	Moment    string
}
