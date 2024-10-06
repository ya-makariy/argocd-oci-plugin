package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ya-makariy/argocd-oci-plugin/version"
)

// NewVersionCommand returns a new instance of the version command
func NewVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "version",
		Short: "Print argocd-oci-plugin version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "argocd-oci-plugin %s (%s) BuildDate: %s\n", version.Version, version.CommitSHA, version.BuildDate)
		},
	}

	return command
}
