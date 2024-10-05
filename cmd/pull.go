package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewPullCommand() *cobra.Command {

	var command = &cobra.Command{
		Use:   "pull [flags] <name>{:<tag>|@<digest>}",
		Short: "Pull files from a registry",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("<name>{:<tag>|@<digest>} argument required to pull files")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}

	return command
}
