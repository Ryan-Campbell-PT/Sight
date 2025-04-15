package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
)

type Daily struct {
	ID			int64
	foodString	string
	date		string
}

// this dbOnce variable makes it so no matter how many times you call the function getDatabase()
// the code inside will only run once
var (
	db *sql.DB
	dbOnce sync.Once
)

func IRRELEVANT() {
	// Capture connection properties.
	// Get a database handle
	cfg := getSqlConfig()
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

/*
func visualizationTest_queryForDailyCalories() ([]Daily, error) {
	duh()
	var daily []Daily;
	rows, err := db.Query("select * from daily")
	if err != nil {
		return nil, fmt.Errorf("dailyQuery %q: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var day Daily
		if err := rows.Scan(&day.ID, &day.foodString, &day.date, &day.calories); err != nil {
			return nil, fmt.Errorf("dailyQuery %q: %v", err)
		}
		daily = append(daily, day)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("dailyQuery %q: %v", err)
	}
	return daily, nil
}
*/

func getDatabase() *sql.DB {
	dbOnce.Do(func() {
		cfg := getSqlConfig()
		db, err := sql.Open("mysql", cfg.FormatDSN())
		if(err != nil) {
			log.Fatal(err)
		}

		pingErr := db.Ping()
		if(pingErr != nil) {
			log.Fatal(pingErr)
		}
	})

	return db
}

func saveToDatabase(data BodyResponse) error {
	db := getDatabase()
	
	_, err := db.Exec("INSERT INTO Daily(Date, FoodString) VALUES(?, ?)", data.Date, data.FoodListString)
	if handleError("Error inserting body values into database: ", err) {
		return err
	}

	return nil
}