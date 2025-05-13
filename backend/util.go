package main

import (
	"fmt"
	"log"
	"math"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	User                        string `env:"DBUSER"`
	Password                    string `env:"DBPASS"`
	Nutritionix_appid           string `env:"nutrition__appid"`
	Nutritionix_appkey          string `env:"nutrition__appkey"`
	Nutritionix_domain          string `env:"nutrition__domain"`
	Nutritionix_naturalLanguage string `env:"nutrition__naturalLanguage"`
	Nutritionix_contentType     string `env:"nutrition__contentType"`

	Azure_User     string `env:"Azure_User"`
	Azure_Password string `env:"Azure_Password"`
	Azure_Database string `env:"Azure_Database"`
	Azure_Server   string `env:"Azure_Server"`
	Azure_Port     int64  `env:"Azure_Port"`
}

func getConfig() Config {
	err := godotenv.Load("../.env")
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

func handleError(printStr string, err error) bool {
	if err != nil {
		fmt.Println(printStr, err)
		return true
	}

	return false
}

func roundToNearestDecimal(num float64, decimal float64) float64 {
	factor := math.Pow(10, decimal)
	return math.Round(num*factor) / factor
}
