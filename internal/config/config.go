package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Host     string `env:"HOST"`
	Port     string `env:"PORT"` // Порт http сервера
	DBHost   string `env:"DB_HOST"`
	DBPort   string `env:"DB_PORT"` // Порт базы данных
	Password string `env:"DB_PASSWORD"`
	User     string `env:"DB_USER"`
	DBName   string `env:"DB_NAME"`
	SSLMode  string `env:"DB_SSLMODE"`

	SECRET_KEY string `env:"SECRET_KEY"`

	REDIS_PASSWORD        string `env:"REDIS_PASSWORD"`
	REDIS_USER            string `env:"REDIS_USER"`
	REDIS_USER_PASSWORD   string `env:"REDIS_USER_PASSWORD"`
	REDIS_HOST            string `env:"REDIS_HOST"`
	REDIS_PORT            string `env:"REDIS_PORT"`
	REDIS_DURATION_SECOND string `env:"REDIS_DURATION_SECOND"` // Формат second

	MigrationsTable string `env:"DB_MIGRATIONTABLE"`
}

func MustLoadPath(configPath string) *Config {
	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}
