package die

import (
	"fmt"
	"io"
	"os"
)

// FDie ends program execution and prints the given message to the selected channel
func FDief(w io.Writer, x string) {
	fmt.Fprintf(w, x)
	os.Exit(1)
}

// Die ends program execution and prints the given message to stderr
func Die(x string) {
	FDief(os.Stderr, x)
}

// DieW accepts err and tries to unwrap to Die
func DieW(help string, x error) {
	Die(fmt.Errorf("%s: %w", help, x).Error())
}

// Must dies if error is not nil
func Must(err error) {
	if err != nil {
		Die(fmt.Errorf("fatal: %w", err).Error())
	}
}

// AllMust dies if any error in the group
func AllMust(errs ...error) {
	for _, err := range errs {
		Must(err)
	}
}
