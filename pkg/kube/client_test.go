package kube

import "testing"

func TestSecretNamespaceName(t *testing.T) {
	testCases := []struct {
		input             string
		expectedNamespace string
		expectedName      string
	}{
		{
			"secretwithoutnamespace",
			"argocd",
			"secretwithoutnamespace",
		},
		{
			"secretnamespace:secretname",
			"secretnamespace",
			"secretname",
		},
	}

	for _, tc := range testCases {
		namespace, name := secretNamespaceName(tc.input)
		if namespace != tc.expectedNamespace {
			t.Errorf("expected namespace: %s, got: %s.", tc.expectedNamespace, namespace)
		}
		if name != tc.expectedName {
			t.Errorf("expected name: %s, got: %s.", tc.expectedName, name)
		}
	}
}
