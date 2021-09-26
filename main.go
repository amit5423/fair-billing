package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"fair-billing/calculation"
)

var out io.Writer = os.Stdout

func main() {
	numOfArgs := len(os.Args)
	if numOfArgs < 2 {
		fmt.Fprintln(out, errors.New("Log file path not provided"))
		os.Exit(1)
	} else if numOfArgs > 2 {
		fmt.Fprintln(out, errors.New("Too may arguments"))
		os.Exit(1)
	}
	fileName := os.Args[1]
	fileName = filepath.FromSlash(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Fprintln(out, "Failed to open file at path", pErr.Path)
			os.Exit(1)
		}
		fmt.Fprintln(out, "Generic error", err)
		os.Exit(1)
	}
	defer file.Close()

	keys, report, err := calculation.Billing(file)
	if err != nil {
		fmt.Fprintln(out, err)
		os.Exit(1)
	} else {
		for _, k := range keys {
			value := report[k]
			fmt.Fprintln(out, k, value.TotalSession, value.TotalDuration)
		}
	}
}
