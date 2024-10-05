package main

import (
	"os"

	"github.com/ya-makriy/argocd-oci-plugin/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
