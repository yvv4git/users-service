package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	envFileName   = fmt.Sprintf("%s_DB.FILENAMEDB", prefixEnvironmet)
	envServerPort = fmt.Sprintf("%s_SERVER.PORT", prefixEnvironmet)
)

func TestConfig(t *testing.T) {
	type environment struct {
		FileNameDB string
		Port       int
	}

	setEnvironvent := func(env environment) {
		// key - value
		os.Setenv(envFileName, env.FileNameDB)
		os.Setenv(envServerPort, fmt.Sprint(env.Port))
	}

	unsetEnvironment := func() {
		os.Unsetenv(envFileName)
		os.Unsetenv(envServerPort)
	}

	testCases := []struct {
		name           string
		useEnvironment bool
		env            environment
		filePathConfig string
		expect         *Config
		wantError      bool
	}{
		{
			name:           "File from file, then from environment",
			useEnvironment: true,
			env: environment{
				FileNameDB: "db/other.db",
				Port:       1234,
			},
			filePathConfig: "fixtures/test",
			expect: &Config{
				DB: DBConfig{
					FileNameDB: "db/other.db",
				},
				Server: ServerConfig{
					Port: 1234,
				},
			},
			wantError: false,
		},
		{
			name:           "Use only file",
			useEnvironment: false,
			env: environment{
				FileNameDB: "db/service.db",
				Port:       8080,
			},
			filePathConfig: "fixtures/test",
			expect: &Config{
				DB: DBConfig{
					FileNameDB: "db/service.db",
				},
				Server: ServerConfig{
					Port: 8080,
				},
			},
			wantError: false,
		},
		{
			name:           "File not found, then use environment",
			useEnvironment: true,
			env: environment{
				FileNameDB: "db/unknown.db",
				Port:       8181,
			},
			filePathConfig: "fixtures/unknown",
			expect: &Config{
				DB: DBConfig{
					FileNameDB: "db/unknown.db",
				},
				Server: ServerConfig{
					Port: 8181,
				},
			},
			wantError: false,
		},
		{
			name:           "File not found, environment empty, then use default",
			useEnvironment: false,
			env: environment{
				FileNameDB: "db/unknown.db",
				Port:       8181,
			},
			filePathConfig: "fixtures/unknown",
			expect: &Config{
				DB: DBConfig{
					FileNameDB: defaultFileNameDB,
				},
				Server: ServerConfig{
					Port: defaultServerPort,
				},
			},
			wantError: false,
		},
		{
			name:           "Bad file stucture, generate error",
			useEnvironment: false,
			filePathConfig: "fixtures/bad",
			wantError:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup environment from test case.
			if tc.useEnvironment {
				setEnvironvent(tc.env)
			} else {
				unsetEnvironment()
			}

			// Init config with config file.
			result, err := Init(tc.filePathConfig)
			if tc.wantError {
				assert.NotEmpty(t, err)
			} else {
				assert.Nil(t, err)
			}

			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestParseFile(t *testing.T) {
	type parsedPath struct {
		dir      string
		fileName string
	}

	testCases := []struct {
		filePath string
		expect   parsedPath
	}{
		{
			filePath: "fixture/test",
			expect: parsedPath{
				dir:      "fixture",
				fileName: "test",
			},
		},
	}

	for _, tc := range testCases {
		dir, filename, err := parseFilePath(tc.filePath)
		assert.Nil(t, err)

		result := parsedPath{
			dir:      dir,
			fileName: filename,
		}
		assert.Equal(t, tc.expect, result, "WTF")
	}
}
