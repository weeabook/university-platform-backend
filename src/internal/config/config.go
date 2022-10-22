package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug       bool `env:"IS_DEBUG" env-default:"false"`
	IsDevelopment bool `env:"IS_DEV" env-default:"false"`
	Listen        struct {
		Type   string `env:"LISTEN_TYPE" env-default:"port"`
		BindIP string `env:"BIND_IP" env-default:"0.0.0.0"`
		Port   string `env:"PORT" env-default:"10000"`
	}
	DatabaseConfig struct {
		Host     string `env:"PG_HOST" env-required:"true"`
		Port     string `env:"PG_PORT" env-required:"true"`
		Username string `env:"PG_USERNAME" env-required:"true"`
		DBName   string `env:"DB_NAME" env-required:"true"`
		SSLMode  string `env:"SSL_MODE" env-default:"false"`
		Password string `env:"PG_PASSWORD" env-required:"true"`
	}
	AppConfig struct {
		LogLevel  string
		AdminUser struct {
			Email    string `env:"ADMIN_EMAIL" env-required:"true"`
			Password string `env:"ADMIN_PASSWORD" env-required:"true"`
		}
	}
}

var instanse *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Print("gather config")

		instanse = &Config{}

		if err := cleanenv.ReadEnv(instanse); err != nil {
			helpText := "University Platform System"
			help, _ := cleanenv.GetDescription(instanse, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})

	return instanse
}
