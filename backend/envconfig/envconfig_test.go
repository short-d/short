package envconfig

import (
	"testing"
	"time"

	"github.com/short-d/app/fw"
	"github.com/short-d/app/mdtest"
)

var _ fw.Environment = (*EnvironmentFake)(nil)

type EnvironmentFake struct {
	envs map[string]string
}

func (e EnvironmentFake) GetEnv(key string, defaultValue string) string {
	val, ok := e.envs[key]
	if !ok {
		return defaultValue
	}
	return val
}

func (e EnvironmentFake) AutoLoadDotEnvFile() {
	panic("implement me")
}

func TestParseConfigFromEnv(t *testing.T) {
	t.Parallel()

	t.Run("string", func(t *testing.T) {
		type config struct {
			DBUser     string `env:"DB_USER" default:"beta"`
			DBPassword string `env:"DB_PASSWORD" default:"password"`
		}

		testCases := []struct {
			name           string
			envs           map[string]string
			config         config
			expectedConfig config
		}{
			{
				name: "parse from environmental variables",
				envs: map[string]string{
					"DB_USER":     "alpha",
					"DB_PASSWORD": "alpha_pw",
				},
				config: config{},
				expectedConfig: config{
					DBUser:     "alpha",
					DBPassword: "alpha_pw",
				},
			},
			{
				name: "use default value",
				envs: map[string]string{
					"DB_USER": "alpha",
				},
				config: config{},
				expectedConfig: config{
					DBUser:     "alpha",
					DBPassword: "password",
				},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				envFake := EnvironmentFake{
					envs: testCase.envs,
				}
				envConfig := EnvConfig{environment: envFake}
				err := envConfig.ParseConfigFromEnv(&testCase.config)
				mdtest.Equal(t, nil, err)
				mdtest.Equal(t, testCase.expectedConfig, testCase.config)
			})
		}
	})

	t.Run("int", func(t *testing.T) {
		type config struct {
			CacheSize int `env:"CACHE_SIZE" default:"2"`
		}

		testCases := []struct {
			name           string
			envs           map[string]string
			config         config
			expectHasError bool
			expectedConfig config
		}{
			{
				name: "parse from environmental variables",
				envs: map[string]string{
					"CACHE_SIZE": "10",
				},
				config: config{},
				expectedConfig: config{
					CacheSize: 10,
				},
			},
			{
				name:   "use default value",
				envs:   map[string]string{},
				config: config{},
				expectedConfig: config{
					CacheSize: 2,
				},
			},
			{
				name: "incorrect format",
				envs: map[string]string{
					"CACHE_SIZE": "random",
				},
				config:         config{},
				expectHasError: true,
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				envFake := EnvironmentFake{
					envs: testCase.envs,
				}
				envConfig := EnvConfig{environment: envFake}
				err := envConfig.ParseConfigFromEnv(&testCase.config)
				if testCase.expectHasError {
					mdtest.NotEqual(t, nil, err)
					return
				}
				mdtest.Equal(t, nil, err)
				mdtest.Equal(t, testCase.expectedConfig, testCase.config)
			})
		}
	})

	t.Run("bool", func(t *testing.T) {
		type config struct {
			IsCacheEnabled bool `env:"IS_CACHE_ENABLED" default:"false"`
		}

		testCases := []struct {
			name           string
			envs           map[string]string
			config         config
			expectHasError bool
			expectedConfig config
		}{
			{
				name: "parse from environmental variables",
				envs: map[string]string{
					"IS_CACHE_ENABLED": "true",
				},
				config: config{},
				expectedConfig: config{
					IsCacheEnabled: true,
				},
			},
			{
				name:   "use default value",
				envs:   map[string]string{},
				config: config{},
				expectedConfig: config{
					IsCacheEnabled: false,
				},
			},
			{
				name: "incorrect format",
				envs: map[string]string{
					"IS_CACHE_ENABLED": "random",
				},
				config:         config{},
				expectHasError: true,
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				envFake := EnvironmentFake{
					envs: testCase.envs,
				}
				envConfig := EnvConfig{environment: envFake}
				err := envConfig.ParseConfigFromEnv(&testCase.config)
				if testCase.expectHasError {
					mdtest.NotEqual(t, nil, err)
					return
				}
				mdtest.Equal(t, nil, err)
				mdtest.Equal(t, testCase.expectedConfig, testCase.config)
			})
		}
	})

	t.Run("time.Duration", func(t *testing.T) {
		type config struct {
			AuthTokenLifetime time.Duration `env:"AUTH_TOKEN_LIFETIME" default:"1w"`
		}

		testCases := []struct {
			name           string
			envs           map[string]string
			config         config
			expectHasError bool
			expectedConfig config
		}{
			{
				name: "parse from environmental variables",
				envs: map[string]string{
					"AUTH_TOKEN_LIFETIME": "2h",
				},
				config: config{},
				expectedConfig: config{
					AuthTokenLifetime: 2 * time.Hour,
				},
			},
			{
				name:   "use default value",
				envs:   map[string]string{},
				config: config{},
				expectedConfig: config{
					AuthTokenLifetime: 168 * time.Hour,
				},
			},
			{
				name: "incorrect format",
				envs: map[string]string{
					"AUTH_TOKEN_LIFETIME": "random",
				},
				config:         config{},
				expectHasError: true,
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				envFake := EnvironmentFake{
					envs: testCase.envs,
				}
				envConfig := EnvConfig{environment: envFake}
				err := envConfig.ParseConfigFromEnv(&testCase.config)
				if testCase.expectHasError {
					mdtest.NotEqual(t, nil, err)
					return
				}
				mdtest.Equal(t, nil, err)
				mdtest.Equal(t, testCase.expectedConfig, testCase.config)
			})
		}
	})

	t.Run("Duration", func(t *testing.T) {
		type fakeDuration int64

		type fakeConfig struct {
			AuthTokenLifetime fakeDuration `env:"AUTH_TOKEN_LIFETIME" default:"1w"`
		}

		testCases := []struct {
			name       string
			envs       map[string]string
			fakeConfig fakeConfig
		}{
			{
				name: "incorrect type",
				envs: map[string]string{
					"AUTH_TOKEN_LIFETIME": "1h",
				},
				fakeConfig: fakeConfig{AuthTokenLifetime: fakeDuration(time.Hour)},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				envFake := EnvironmentFake{
					envs: testCase.envs,
				}
				envConfig := EnvConfig{environment: envFake}
				err := envConfig.ParseConfigFromEnv(&testCase.fakeConfig)
				mdtest.NotEqual(t, nil, err)
				return
			})
		}
	})
}
