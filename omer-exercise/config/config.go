package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	DB     DatabaseConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
    viper.BindEnv("db.host", "DB_HOST")
    viper.BindEnv("db.port", "DB_PORT")
    viper.BindEnv("db.user", "DB_USER")
    viper.BindEnv("db.password", "DB_PASSWORD")
    viper.BindEnv("db.name", "DB_NAME")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
