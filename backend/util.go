package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/caarlos0/env/v6"
)

type Config struct {
    User string `env:"DBUSER"`
    Password string `env:"DBPASS"`
    Nutritionix_appid string `env:"nutrition__appid"`
    Nutritionix_appkey string `env:"nutrition__appkey"`
    Nutritionix_domain string `env:"nutrition__domain"`
    Nutritionix_naturalLanguage string `env:"nutrition__naturalLanguage"`
    Nutritionix_contentType string `env:"nutrition__contentType"`
}

func getEnvironmentVariables() Config {
    err := godotenv.Load("/home/lyon/Documents/Sight/.env")
    if err != nil {
        log.Fatalf("unable to load .env file")
    }
	cfg := Config{}
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}

	return cfg
}