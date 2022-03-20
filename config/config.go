package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		HTTP
		PG
		Log
	}

	// App -.
	// App struct {
	// 	Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
	// 	Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	// }

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// PG -.
	PG struct {
		Host     string `env-required:"true" env:"PG_HOST"`
		Port     string `env-required:"true" env:"PG_PORT"`
		User     string `env-required:"true" env:"PG_USER"`
		Password string `env-required:"true" env:"PG_PASSWORD"`
		SSLMode  string `env-required:"true" env:"PG_SSLMODE"`
	}
)

func (pg PG) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s  sslmode=%s",
		pg.Host,
		pg.Port,
		pg.User,
		pg.Password,
		pg.SSLMode,
	)
}

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(".env", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
