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
	"github.com/joho/godotenv"
)

// Structure for store time response in JSON format
type TimeResponse struct {
	CurrentTime string `json:"current_time"`
	Location    string `json:"location"`
}

var db *sql.DB

// Function to connect to MySQL database
func connectDB() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	// Connection string format
	dbSrcName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName)

	// Using backoff to retry DB connection

	db, err = sql.Open("mysql", dbSrcName)
	if err != nil {
		log.Printf("Unable to connect to database: %v", err)

	}

	// Ping to test the connection
	if err = db.Ping(); err != nil {
		log.Printf("Database ping failed: %v", err)

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

	// Insert current time into the database
	_, err = db.Exec("INSERT INTO time_log (timestamp) VALUES (?)", formattedTime)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}
	log.Println("Data inserted successfully!")

	// Prepare and send the response
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
	// Connect to DB
	connectDB()

	// Start the server for calling the API
	http.HandleFunc("/time", getTimeZoneHandler)

	port := ":8080"
	fmt.Println("Server is running on port", port)
	http.ListenAndServe(port, nil)
}
