package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig_Success(t *testing.T) {
	os.Setenv("APP_ENV", EnvDevelopment)
	os.Setenv("APP_PORT", "8080")
	os.Setenv("DB_URL", "postgres://user:pass@localhost:5432/dbname")
	viper.Reset()

	cfg, err := New(NewViper())

	assert.NoError(t, err, "Expected no error when loading config")
	assert.NotNil(t, cfg, "Config should not be nil")
	assert.Equal(t, EnvDevelopment, cfg.Env, "Expected ENV to be development")
	assert.Equal(t, "8080", cfg.Port, "Expected PORT to be 8080")
	assert.Equal(t, "postgres://user:pass@localhost:5432/dbname", cfg.DbUrl, "Expected correct Postgres URL")
}

func TestIsProduction(t *testing.T) {
	os.Setenv("APP_ENV", EnvProduction)
	viper.Reset()

	cfg, _ := New(NewViper())

	assert.True(t, cfg.IsProduction())
}

func TestNew_ReadInConfigError(t *testing.T) {
	viper.Reset()

	v := viper.New()
	v.SetConfigFile("invalid.yaml")

	_, err := New(v)

	assert.Error(t, err, "Expected an error due to invalid config type")
}
