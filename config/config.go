package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

const (
	EnvProduction  = "production"
	EnvDevelopment = "development"
)

type Config struct {
	Env   string
	Port  string
	DbUrl string
}

func (c *Config) IsProduction() bool {
	return c.Env == EnvProduction
}

func NewViper() *viper.Viper {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return v
}

func New(v *viper.Viper) (*Config, error) {
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
		fmt.Println("No config file found. Using defaults and environment variables")
	}

	return &Config{
		Env:   v.GetString("app.env"),
		Port:  v.GetString("app.port"),
		DbUrl: v.GetString("db.url"),
	}, nil
}
