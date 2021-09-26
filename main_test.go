package main

import (
	"bytes"
	"os"
	"testing"
)

func TestCLI(t *testing.T) {

	for _, testcases := range []struct {
		Args   []string
		Output string
	}{
		{
			Args:   []string{"./fair-billing", "logfiles/logs0"},
			Output: "ALICE99 4 240\nCHARLIE 3 37\n",
		},
		{
			Args:   []string{"./fair-billing", "logfiles/logs1"},
			Output: "ALICE99 1 31\nCHARLIE 1 2\n",
		},
		{
			Args:   []string{"./fair-billing", "logfiles/logs2"},
			Output: "ALICE99 2 18\nCHARLIE 1 0\n",
		},
		{
			Args:   []string{"./fair-billing", "logfiles/logs3"},
			Output: "ALICE 1 275\nRUBY 3 461\nTOM 2 157\n",
		},
		{
			Args:   []string{"./fair-billing", "logfiles/logs4"},
			Output: "Harry 1 3312\nHarshita 1 46851\nJACOB01 1 7148\nNick07 2 57318\n",
		},
	} {
		t.Run("", func(t *testing.T) {
			os.Args = testcases.Args
			out = bytes.NewBuffer(nil)
			main()

			if got := out.(*bytes.Buffer).String(); got != testcases.Output {
				t.Errorf("expected %s, but got %s", testcases.Output, got)
			}
		})
	}

}
