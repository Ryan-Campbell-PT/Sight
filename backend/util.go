package main

import (
	"fmt"
	"log"
	"math"

	"github.com/caarlos0/env/v6"
	"github.com/go-sql-driver/mysql"
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
}

func getEnvironmentVariables() Config {
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

func getSqlConfig() mysql.Config {
	cfg := getEnvironmentVariables()
	return mysql.Config{
		User:                 cfg.User,
		Passwd:               cfg.Password,
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "consume",
		AllowNativePasswords: true,
	}
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
