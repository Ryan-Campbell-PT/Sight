package main
/*
import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Daily struct {
	ID			int64
	foodString	string
	date		string
}

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User: os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "consume",
		AllowNativePasswords: true,
	}
	// Get a database handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if(err != nil) {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if(pingErr != nil) {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	dailyValues, err := dailyQuery("1/1/2025")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("daily Values: %v\n", dailyValues)
}

func dailyQuery(date string) ([]Daily, error) {
	var daily []Daily

	rows, err := db.Query("SELECT * FROM daily WHERE date = ?", date)
	if err != nil {
		return nil, fmt.Errorf("dailyQuery %q: %v", date, err)
	}
	defer rows.Close()

	for rows.Next() {
		var day Daily
		if err := rows.Scan(&day.ID, &day.foodString, &day.date); err != nil {
			return nil, fmt.Errorf("dailyQuery %q: %v", date, err)
		}
		daily = append(daily, day)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("dailyQuery %q: %v", date, err)
	}
	return daily, nil
}
	*/