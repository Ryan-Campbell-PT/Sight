package util

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"time"

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

func GetEnvConfig() Config {
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

func HandleError(printStr string, err error) bool {
	if err != nil {
		fmt.Println(printStr, err)
		return true
	}

	return false
}

func RoundToNearestDecimal(num float64, decimal float64) float64 {
	factor := math.Pow(10, decimal)
	return math.Round(num*factor) / factor
}

func GetDate(date time.Time) string {
	return date.Format(time.DateOnly)
}

func ReadRequestBody(body io.ReadCloser) ([]byte, error) {
	defer body.Close()
	data, err := io.ReadAll(body)
	return data, err
}

func SendHttpRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if HandleError("SendHttpRequest/Error sending http request: ", err) {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func IntPtr(i int) *int {
	return &i
}
