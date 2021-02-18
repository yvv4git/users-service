package config

import (
	"strings"

	"github.com/spf13/viper"
)

const (
	defaultConfigType = "yml"
	defaultFileNameDB = "db/default.db"
	defaultServerPort = 1357
	prefixEnvironmet  = "APP"
)

// Config is used as config of this app.
type Config struct {
	DB     DBConfig
	Server ServerConfig
}

type (
	// DBConfig is used as config struct for db.
	DBConfig struct {
		FileNameDB string `mapstructure:"FILENAMEDB"`
	}

	// ServerConfig is used for example.
	ServerConfig struct {
		Port int `mapstructure:"PORT"`
	}
)

// Init is used for create instance of config.
func Init(filePath string) (conf *Config, err error) {
	dir, fileName, err := parseFilePath(filePath)
	if err != nil {
		return nil, err
	}

	var runtimeViper = viper.New()

	// Set default values and init settings.
	populateDefaults(runtimeViper)

	// Load from file.
	runtimeViper.AddConfigPath(dir)
	runtimeViper.SetConfigName(fileName)
	runtimeViper.SetConfigType(defaultConfigType)

	// Load from evironment if exists.
	runtimeViper.SetEnvPrefix(prefixEnvironmet)
	runtimeViper.AutomaticEnv()

	// Fill viper map.
	if err = runtimeViper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return
		}
	}

	// Fill config from viper.
	err = runtimeViper.Unmarshal(&conf)
	return
}

func populateDefaults(viperRuntime *viper.Viper) {
	viperRuntime.SetDefault("DB.FileNameDB", defaultFileNameDB)
	viperRuntime.SetDefault("Server.Port", defaultServerPort)
}

func parseFilePath(filePath string) (dir string, fileName string, err error) {
	path := strings.Split(filePath, "/")
	if len(path) < 2 {
		return "", "", ErrFileNotFound
	}
	dir = path[0]
	fileName = path[1]

	return
}
