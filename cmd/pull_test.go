package cmd

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {

	t.Run("will throw an error expecting arguments", func(t *testing.T) {
		args := []string{}
		cmd := NewPullCommand()

		c := bytes.NewBufferString("")
		cmd.SetArgs(args)
		cmd.SetErr(c)
		cmd.SetOut(bytes.NewBufferString(""))
		cmd.Execute()
		out, err := io.ReadAll(c)
		if err != nil {
			t.Fatal(err)
		}
		expected := "<name>{:<tag>|@<digest>} argument required to pull files"
		if !strings.Contains(string(out), expected) {
			t.Fatalf("expected to contain: %s but got %s", expected, out)
		}
	})
}
