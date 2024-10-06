package config_test

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/ya-makariy/argocd-oci-plugin/pkg/config"
)

func TestNewConfig(t *testing.T) {
	testCases := []struct {
		environment    map[string]interface{}
		expectedFsPath string
	}{
		{
			map[string]interface{}{
				"AOP_USERNAME": "user",
				"AOP_PASSWORD": "password",
				"AOP_FS_PATH":  "values",
			},
			"values",
		},
		{
			map[string]interface{}{
				"AOP_USERNAME": "user",
				"AOP_PASSWORD": "password",
			},
			"./oci-files/",
		},
		{
			map[string]interface{}{},
			"./oci-files/",
		}, // empty config-file
	}
	for _, tc := range testCases {
		for k, v := range tc.environment {
			os.Setenv(k, v.(string))
		}
		viper := viper.New()
		config, err := config.New(viper, &config.Options{})
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		xFsPath := config.FsPath
		if xFsPath != tc.expectedFsPath {
			t.Errorf("expected: %s, got: %s.", tc.expectedFsPath, xFsPath)
		}
		for k := range tc.environment {
			os.Unsetenv(k)
		}
	}
}
