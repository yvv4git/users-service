package config

import "github.com/spf13/viper"

const (
	defaultFileNameDB = "db/example.db"
	defaultOtherVal   = 5
)

// Config is used as config of this app.
type Config struct {
	DB    DBConfig
	Other OtherConfig
}

type (
	// DBConfig is used as config struct for db.
	DBConfig struct {
		FileNameDB string
	}

	// OtherConfig is used for example.
	OtherConfig struct {
		OtherVal int
	}
)

func setUpByViper() error {
	fillDefaults()

	if err := parseConfigFile(); err != nil {
		return err
	}

	return parseDbEnvVariables()
}

func fillDefaults() {
	viper.SetDefault("DB.FileNameDB", defaultFileNameDB)
	viper.SetDefault("Other.OtherVal", defaultOtherVal)
}

func parseConfigFile() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("default")
	return viper.ReadInConfig()
}

func parseDbEnvVariables() error {
	viper.SetEnvPrefix("db")

	if err := viper.BindEnv("FileNameDB"); err != nil {
		return err
	}

	return nil
}
