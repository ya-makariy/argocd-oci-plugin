package main

import (
	"os"

	cmd "github.com/ya-makariy/argocd-oci-plugin/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
