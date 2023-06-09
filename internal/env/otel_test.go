package env_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/angelokurtis/go-starter-otel/_test"
	"github.com/angelokurtis/go-starter-otel/internal/env"
)

func TestNewFromEnv(t *testing.T) {
	t.Run(`Given no environment variables have been set
When a new configuration object is created
Then the object should have default values and no errors should occur`, func(t *testing.T) {
		cfg, err := env.LookupOTel()
		require.NoError(t, err)

		assert.NotNil(t, cfg)
	})
	t.Run(`Given an environment variable has been set with an invalid value
When a new configuration object is created
Then an error should be returned`, func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{"OTEL_TRACES_SAMPLER_ARG": "example.com"})
		defer envvars.Unset()

		cfg, err := env.LookupOTel()
		assert.Error(t, err)

		assert.Nil(t, cfg)
	})
}
