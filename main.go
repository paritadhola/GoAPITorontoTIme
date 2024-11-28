package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Structure for store time response in jason format
type TimeResponse struct {
	CurrentTime string `json:"current_time"`
	Location    string `json:"location"`
}

var db *sql.DB

func connectDB() {
	var err error
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	dbSrcName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName)
	db, err = sql.Open("mysql", dbSrcName)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}
	fmt.Println("Database connected successfully!")

}

func getTimeZoneHandler(w http.ResponseWriter, r *http.Request) {

	loc, err := time.LoadLocation("America/Toronto") // Static input for Toronto time
	if err != nil {
		http.Error(w, "Error loading timezone", http.StatusBadRequest)
		return
	}

	// Store the value of current timezone
	currentTime := time.Now().In(loc)
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	log.Printf("Current Toronto time: %s", formattedTime)

	_, err = db.Exec("INSERT INTO time_log (timestamp) VALUES (?)", formattedTime)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}
	log.Println("Data inserted successfully!")

	// Get the response
	response := TimeResponse{
		CurrentTime: formattedTime,
		Location:    "Toronto",
	}

	// Print the response in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	log.Println("Response sent to client.")
}

func main() {

	dbSrcName := "root:12345678@tcp(localhost:3306)/GoTimeZoneAPI"
	db, err := sql.Open("mysql", dbSrcName)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	_, err = db.Exec("INSERT INTO time_log (timestamp) VALUES (NOW())")
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}
	log.Println("Data inserted successfully!")

	// Start the server for callling API
	http.HandleFunc("/time", getTimeZoneHandler)

	port := ":8080"
	println("Server is running on port", port)
	http.ListenAndServe(port, nil)
}
