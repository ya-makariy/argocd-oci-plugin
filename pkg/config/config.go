package config

import (
	"bytes"
	"os"
	"strings"

	"github.com/spf13/viper"
	"github.com/ya-makariy/argocd-oci-plugin/pkg/backends"
	"github.com/ya-makariy/argocd-oci-plugin/pkg/kube"
	"github.com/ya-makariy/argocd-oci-plugin/pkg/types"
	"github.com/ya-makariy/argocd-oci-plugin/pkg/utils"
	"oras.land/oras-go/v2/registry/remote/auth"
	"oras.land/oras-go/v2/registry/remote/retry"
)

// Options options that can be passed to a Config struct
type Options struct {
	SecretName string
	ConfigPath string
	Registry   string
}

// Config is used to decide the backend and auth type
type Config struct {
	Backend types.Backend
	FsPath  string
}

func New(v *viper.Viper, co *Options) (*Config, error) {
	// Read in config file or kubernetes secret and set as env vars
	err := readConfigOrSecret(co.SecretName, co.ConfigPath, v)
	if err != nil {
		return nil, err
	}

	// Instantiate Env
	utils.VerboseToStdErr("reading configuration from environment, overriding any previous settings")
	v.AutomaticEnv()

	utils.VerboseToStdErr("AVP configured with the following settings:\n")
	for k, viperValue := range v.AllSettings() {
		utils.VerboseToStdErr("%s: %s\n", k, viperValue)
	}

	var backend types.Backend

	fspath := v.GetString(types.EnvFsPath)

	credential := auth.EmptyCredential
	if v.IsSet(types.EnvUsername) ||
		v.IsSet(types.EnvPassword) {
		credential = auth.Credential{
			Username: v.GetString(types.EnvUsername),
			Password: v.GetString(types.EnvPassword),
		}
	}
	backend = backends.NewRegistryBackend(auth.Client{
		Client:     retry.DefaultClient,
		Cache:      auth.NewCache(),
		Credential: auth.StaticCredential(co.Registry, credential),
	})

	return &Config{
		Backend: backend,
		FsPath:  fspath,
	}, nil
}

func readConfigOrSecret(secretName, configPath string, v *viper.Viper) error {
	// If a secret name is passed, pull config from Kubernetes
	if secretName != "" {
		utils.VerboseToStdErr("reading configuration from secret %s", secretName)

		localClient, err := kube.NewClient()
		if err != nil {
			return err
		}
		yaml, err := localClient.ReadSecret(secretName)
		if err != nil {
			return err
		}
		v.SetConfigType("yaml")
		v.ReadConfig(bytes.NewBuffer(yaml))
	}

	// If a config file path is passed, read in that file and overwrite all other
	if configPath != "" {
		utils.VerboseToStdErr("reading configuration from config file %s, overriding any previous settings", configPath)

		v.SetConfigFile(configPath)
		err := v.ReadInConfig()
		if err != nil {
			return err
		}
	}

	// Check for ArgoCD 2.4 prefixed environment variables
	for _, envVar := range os.Environ() {
		if strings.HasPrefix(envVar, types.EnvArgoCDPrefix) {
			envVarPair := strings.SplitN(envVar, "=", 2)
			key := strings.TrimPrefix(envVarPair[0], types.EnvArgoCDPrefix+"_")
			val := envVarPair[1]
			v.Set(key, val)
		}
	}

	return nil
}
