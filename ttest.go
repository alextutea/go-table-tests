package ttest

import (
	"fmt"
	"reflect"
	"strconv"
)

const (
	SuccessMark = "\u2713"
	FailureMark = "\u2717"
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
)

type Case struct {
	In               interface{}
	ExpectedOut      interface{}
	ErrTypeCheckFunc func(error) bool //function that determines whether the returned error is the correct error
	Desc             string           //test premise (e.g. "when adding 2 and 2 together...")
	ExpectedBehavior string           //test expectation (e.g. "...it should return 4")
}

func (c Case) Check(out interface{}, err error) (bool, string) {
	isCorrectOut := reflect.DeepEqual(out, c.ExpectedOut)
	isCorrectErr := c.ErrTypeCheckFunc == nil
	if err != nil {
		isCorrectErr = !isCorrectErr && c.ErrTypeCheckFunc(err)
	}
	msg := "OK"
	passed := isCorrectErr && isCorrectOut
	if passed {
		return passed, msg
	}
	msg = fmt.Sprintf(
		"\nFailed check:\n- Input: %s\n- Expected output: %s\n- Actual output: %s\n- Correct error: %s",
		c.In, c.ExpectedOut, out, strconv.FormatBool(isCorrectErr),
	)
	if !isCorrectErr {
		msg += "\n- Error: %s"
	}
	return passed, msg
}

func SuccessMessage(msg string) string {
	return statusMessage(msg, SuccessMark, colorGreen)
}

func FailureMessage(msg string) string {
	return statusMessage(msg, FailureMark, colorRed)
}

func statusMessage(msg, mark, color string) string {
	return fmt.Sprintf("%s\t%s\t%s%s", color, mark, msg, colorReset)
}
