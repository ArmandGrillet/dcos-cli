package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var configString = `
[core]
ssl_verify = "/home/mesosphere/.dcos/clusters/456934a6-fc1f-4d52-9419-a4c722abe2da/dcos_ca.crt"
dcos_url = "https://dcos.example.com"
dcos_acs_token = "eyJ0eXAiOFJHFADCJhbGciOiJSUzI1NiJ9"

[marathon]
url = "https://marathon.dcos.example.com"
dummy = "qwerty"

[cluster]
name = "mr-cluster-1835928138"
`

func TestStringAt(t *testing.T) {
	config, err := FromString(configString)
	require.NoError(t, err)

	expectedOutputs := []struct {
		key []string
		out string
	}{
		{[]string{"core", "dcos_url"}, "https://dcos.example.com\n"},
		{[]string{"marathon", "url"}, "https://marathon.dcos.example.com\n"},
		{[]string{"marathon"}, "marathon.dummy qwerty\nmarathon.url https://marathon.dcos.example.com\n"},
		{[]string{"cluster", "name"}, "mr-cluster-1835928138\n"},
		{[]string{"cluster"}, "cluster.name mr-cluster-1835928138\n"},
	}

	for _, expectedOutput := range expectedOutputs {
		out, err := config.StringAt(expectedOutput.key)
		require.NoError(t, err)
		require.Equal(t, expectedOutput.out, out)
	}
}

func TestString(t *testing.T) {
	config, err := FromString(configString)
	require.NoError(t, err)

	expectedOutput := `core.dcos_acs_token ********
core.dcos_url https://dcos.example.com
core.ssl_verify /home/mesosphere/.dcos/clusters/456934a6-fc1f-4d52-9419-a4c722abe2da/dcos_ca.crt
cluster.name mr-cluster-1835928138
marathon.dummy qwerty
marathon.url https://marathon.dcos.example.com
`
	require.Equal(t, expectedOutput, config.String())
}

func TestSSL(t *testing.T) {

	expectedSSL := []struct {
		config string
		secure bool
		capath string
	}{
		{"[core]\nssl_verify = \"True\"", false, ""},
		{"[core]\nssl_verify = \"true\"", false, ""},
		{"[core]\nssl_verify = \"False\"", true, ""},
		{"[core]\nssl_verify = \"false\"", true, ""},
		{"[core]\nssl_verify = \"/path/to/ca\"", false, "/path/to/ca"},
	}

	for _, exp := range expectedSSL {
		config, err := FromString(exp.config)
		require.NoError(t, err)
		require.Equal(t, exp.secure, config.Core.SSL.Insecure())
		require.Equal(t, exp.capath, config.Core.SSL.CAPath())
	}
}
