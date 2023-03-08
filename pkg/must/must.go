package must

import (
	"fmt"
	"io"
	"os"
)

// FDief ends program execution and prints the given message to the selected writer
func FDief(w io.Writer, x string) {
	fmt.Fprintf(w, x)
	os.Exit(1)
}

// Die ends program execution and prints the given message to stderr
func Die(x string) {
	FDief(os.Stderr, x)
}

// Dief accepts err and tries to unwrap to Die
func Dief(help string, x error) {
	Die(fmt.Errorf("%s: %w", help, x).Error())
}

// Must dies if error is not nil
func Must(err error) {
	if err != nil {
		Die(err.Error())
	}
}

// Mustf dies if error is not nil with a custom message
func Mustf(msg string, err error) {
	if err != nil {
		Die(fmt.Errorf("%s: %w", msg, err).Error())
	}
}

func Get[T any](x T, err error) T {
	Mustf("failed getting value", err)
	return x
}

// AllMust dies if any error in the group
func AllMust(errs ...error) {
	for _, err := range errs {
		Must(err)
	}
}
