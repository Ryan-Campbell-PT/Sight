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

	Azure_User     string `env:"Azure_User"`
	Azure_Password string `env:"Azure_Password"`
	Azure_Database string `env:"Azure_Database"`
	Azure_Server   string `env:"Azure_Server"`
	Azure_Port     int64  `env:"Azure_Port"`
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

var server = "pt-industries.database.windows.net"
var port = 1433
var user = "PT-Industries-Lyon"
var password = "Pod-Around-Relation"
var database = "consume"

func getMsSqlConnectionString() string {
	return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	// cfg := getEnvironmentVariables()
	// return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
	// 	cfg.Azure_Server, cfg.Azure_User, cfg.Azure_Password, cfg.Azure_Port, cfg.Azure_Database)
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
