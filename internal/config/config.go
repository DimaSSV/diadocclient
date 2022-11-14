package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	IsDebug  bool   `env:"DEBUG" env-default:"true"`
	LogLevel string `env:"LOGLEVEL" env-default:"warn"`
	Diadoc   struct {
		Login    string `env:"DIADOC_LOGIN" env-required:"true"`
		Password string `env:"DIADOC_PASSWORD" env-required:"true"`
		Host     string `env:"DIADOC_HOST" env-default:"diadoc-api.kontur.ru"`
		ClientID string `env:"DIADOC_CLIENT_ID" env-required:"true"`
		Token    string `env:"DIADOC_AUTH_TOKEN"`
	}
}

func New() Config {
	cfg := Config{}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(err)
	}
	return cfg
}
