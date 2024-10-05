package cmd

import (
	"github.com/spf13/cobra"
)

// NewRootCommand returns a new instance of the root command
func NewRootCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "argocd-oci-plugin",
		Short: "This is a plugin to get files from OCI-based registries",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	command.AddCommand(NewPullCommand())
	command.AddCommand(NewVersionCommand())

	return command
}
