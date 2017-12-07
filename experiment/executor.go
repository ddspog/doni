package experiment

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"text/template"
	"time"

	"github.com/ddspog/doni/exec"
	"github.com/spf13/viper"
)

const (
	// MaximumExperimentCall is the max number of times a experiment
	// can be called. This exists to prevent infinite loops on cyclic
	// dependencies
	MaximumExperimentCall = 100
)

// Executor executes a Expfile.
type Executor struct {
	Dir     string
	Verbose bool
	Output  string

	Experiments map[string]*Experiment

	Context context.Context

	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer

	experimentCallCount map[string]*int32

	dynamicCache   map[string]string
	muDynamicCache sync.Mutex
}

// ParseExpFile loads the experiment file defining various experiments.
func (e *Executor) ParseExpFile() error {
	defer e.FillNames()
	return viper.Unmarshal(&e.Experiments)
}

// FillNames fill some objects with their names for better logging.
func (e *Executor) FillNames() {
	for expName, exp := range e.Experiments {
		exp.Experiment = expName
		for colName, col := range exp.Collect {
			col.Collect = colName
		}
	}
}

// Run runs Experiment
func (e *Executor) Run(exp string) error {
	if e.Context == nil {
		e.Context = context.Background()
	}
	if e.Stdin == nil {
		e.Stdin = os.Stdin
	}
	if e.Stdout == nil {
		e.Stdout = os.Stdout
	}
	if e.Stderr == nil {
		e.Stderr = os.Stderr
	}

	e.experimentCallCount = make(map[string]*int32, len(e.Experiments))
	for k := range e.Experiments {
		e.experimentCallCount[k] = new(int32)
	}

	if e.dynamicCache == nil {
		e.dynamicCache = make(map[string]string, 10)
	}

	// check if given experiments exist
	if _, ok := e.Experiments[exp]; !ok {
		// FIXME: move to the main package
		e.PrintExperimentsHelp()
		return &experimentNotFoundError{experimentName: exp}
	}

	if err := e.RunExperiment(e.Context, exp); err != nil {
		return err
	}

	return nil
}

// RunExperiment runs a experiment by its name
func (e *Executor) RunExperiment(ctx context.Context, expName string) error {
	exp := e.Experiments[expName]

	exp.Result = make([]map[string]interface{}, len(exp.Data.Values))

	for exp.Iteration = 0; exp.Iteration < len(exp.Data.Values); exp.Iteration++ {
		exp.Vars = make(map[string]interface{})
		exp.Result[exp.Iteration] = make(map[string]interface{})

		exp.Vars["dataStr"] = fmt.Sprintf("%v", exp.Data.Values[exp.Iteration])
		for j := 0; j < len(exp.Data.Columns); j++ {
			exp.Vars[exp.Data.Columns[j]] = exp.Data.Values[exp.Iteration][j]
			exp.Result[exp.Iteration][exp.Data.Columns[j]] = exp.Data.Values[exp.Iteration][j].(interface{})
		}

		for i := range exp.Phases {
			if err := e.runPhase(ctx, exp, i); err != nil {
				return err
			}
		}
	}

	e.printOutput(exp)

	return nil
}

func (e *Executor) printOutput(exp *Experiment) {
	file, errCreate := os.Create(e.Output)
	_ = file.Truncate(0)

	if errCreate != nil {
		panic(fmt.Sprintf("error: invalid output, %s", errCreate))
	}

	defer file.Close()

	if b, err := json.MarshalIndent(exp.Result, "", "\t"); err != nil {
		panic(fmt.Sprintf("error: can't unmarshal result, %s", err))
	} else {
		file.WriteAt(b, 0)
	}
}

func (e *Executor) runPhase(ctx context.Context, exp *Experiment, id int) error {
	phs := exp.Phases[id]

	for i := range phs.Cmds {
		if err := e.runCommand(ctx, exp, phs, i); err != nil {
			return &taskRunError{exp.Experiment, err}
		}
	}
	return nil
}

func (e *Executor) runCommand(ctx context.Context, exp *Experiment, phs *Phase, i int) error {
	cmd := phs.Cmds[i]

	if cmd.Cmd == "" {
		return e.runCollect(ctx, exp, phs.Phase, cmd.Moment, cmd.Collect...)
	}

	var newCmd bytes.Buffer
	t, _ := template.New("").Parse(cmd.Cmd)
	if err := t.Execute(&newCmd, exp.Vars); err != nil {
		return err
	}
	cmd.ParsedCmd = newCmd.String()

	return exec.RunCommand(&exec.RunCommandOptions{
		Context: ctx,
		Command: cmd.ParsedCmd,
		Dir:     e.Dir,
		Stdin:   e.Stdin,
		Stdout:  e.Stdout,
		Stderr:  e.Stderr,
	})
}

func (e *Executor) runCollect(ctx context.Context, exp *Experiment, phs string, moment string, collectCalls ...*string) error {
	for _, c := range collectCalls {
		colCall := exp.Collect[*c]

		index := fmt.Sprintf("%s.%s.%s", phs, moment, colCall.Collect)

		switch colCall.Param {
		case "time":
			exp.Result[exp.Iteration][index] = time.Now()
		case "size":
			var newFile bytes.Buffer
			t, _ := template.New("").Parse(colCall.File)
			if err := t.Execute(&newFile, exp.Vars); err != nil {
				return err
			}
			colCall.ParsedFile = newFile.String()

			exp.Result[exp.Iteration][index] = e.getFileSize(colCall.ParsedFile)
		default:
			return fmt.Errorf("error: invalid param type")
		}
	}

	return nil
}

func (e *Executor) getFileSize(f string) string {
	file, err := os.Open(f)
	defer file.Close()
	if err != nil {
		e.verboseErrf("error: opening file %s, %s", f, err)
	}

	fst, err := file.Stat()
	if err != nil {
		e.verboseErrf("error: checking stats of file %s, %s", f, err)
	}
	return fmt.Sprintf("%v", fst.Size())
}
