package configs

import (
	"github.com/BurntSushi/toml"
)

// DBConfig represents database connection configuration information.
type DBConfig struct {
	User     string `toml:"DB_USER"`
	Password string `toml:"DB_PASS"`
	Host     string `toml:"DB_HOST"`
	Port     int    `toml:"DB_PORT"`
	Name     string `toml:"DB_NAME"`
}

// Config represents application configuration.
type Config struct {
	DB DBConfig `toml:"database"`
}

// NewConfig return configuration struct.
func NewConfig(path string) (Config, error) {
	var config Config
	_, err := toml.DecodeFile(path+"db_connect.tml", &config)
	if err != nil {
		panic(err)
	}

	return config, err
}
