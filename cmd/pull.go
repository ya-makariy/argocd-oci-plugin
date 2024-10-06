package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ya-makariy/argocd-oci-plugin/pkg/config"
	"github.com/ya-makariy/argocd-oci-plugin/pkg/types"
	"oras.land/oras-go/v2/registry/remote"
)

func NewPullCommand() *cobra.Command {
	var configPath, secretName string
	var verboseOutput bool
	var disableCache bool

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
			var repository, reg, tag string

			image := strings.Split(args[0], ":")
			repository = image[0]
			reg = strings.Split(repository, "/")[0]
			if len(image) == 1 {
				tag = "latest"
			} else {
				tag = image[1]
			}

			v := viper.New()
			viper.Set("verboseOutput", verboseOutput)
			viper.Set("disableCache", disableCache)
			viper.SetDefault(types.EnvFsPath, types.DefaultFsPath)
			cmdConfig, err := config.New(v, &config.Options{
				SecretName: secretName,
				ConfigPath: configPath,
				Registry:   reg,
			})
			if err != nil {
				return err
			}

			repo, err := remote.NewRepository(repository)
			if err != nil {
				return err
			}

			err = cmdConfig.Backend.PullFiles(repo, tag, cmdConfig.FsPath)
			if err != nil {
				return err
			}
			return nil
		},
	}

	command.Flags().StringVarP(&configPath, "config-path", "c", "", "path to a file containing Vault configuration (YAML, JSON, envfile) to use")
	command.Flags().StringVarP(&secretName, "secret-name", "s", "", "name of a Kubernetes Secret in the argocd namespace containing Vault configuration data in the argocd namespace of your ArgoCD host (Only available when used in ArgoCD). The namespace can be overridden by using the format <namespace>:<name>")
	command.Flags().BoolVar(&verboseOutput, "verbose-sensitive-output", false, "enable verbose mode for detailed info to help with debugging. Includes sensitive data (credentials), logged to stderr")
	command.Flags().BoolVar(&disableCache, "disable-token-cache", false, "disable the automatic token cache feature that store tokens locally")
	return command
}
