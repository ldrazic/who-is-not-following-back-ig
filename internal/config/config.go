package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
)

type config struct {
	AppConfig
}

var Config config

func init() {
	if err := env.Parse(&Config); err != nil {
		logrus.Fatalf("Error initializing: %s", err.Error())
	}
}

type AppConfig struct {
	Base64InstagramCookie string `env:"BASE64_IG_COOKIE"`
	InstagramAppID        string `env:"IG_APP_ID"`
	InstagramUserID       string `env:"IG_USER_ID"`
}
