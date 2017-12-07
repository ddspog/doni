package experiment

import "fmt"

type experimentNotFoundError struct {
	experimentName string
}

func (err *experimentNotFoundError) Error() string {
	return fmt.Sprintf(`experiment: Experiment "%s" not found`, err.experimentName)
}

type taskRunError struct {
	experimentName string
	err            error
}

func (err *taskRunError) Error() string {
	return fmt.Sprintf(`task: Failed to run task "%s": %v`, err.experimentName, err.err)
}
