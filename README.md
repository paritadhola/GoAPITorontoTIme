## Go API for Current Toronto Time with MySQL Logging

This project provides an API that returns the current time in Toronto in JSON format. It also logs each request's timestamp into a MySQL database. The project uses Go for the API development and Docker for easy deployment.

## Prerequisites

Before starting, make sure you have the following installed on your machine:

**Docker:** Handles multiple users and concurrent requests.

**Go:** For running the Go application.

**MySQL:** Database for storing the time logs which handled by Docker.

## Setup and Installation

**1. Clone the Repository**

Clone this repository to your local machine:

    git clone https://github.com/yourusername/go-timezone-api.git
    cd go-timezone-api

**2. Install Dependencies (if not using Docker)**

If you are running the application locally without Docker, make sure to install Go dependencies.

    go mod tidy

## API Endpoints

**GET /time**

This endpoint returns the current time in Toronto in JSON format.

**Response Example**

    {
        "current_time": "2024-11-28 14:30:45",
        "location": "Toronto"
    }

## Database Setup

**1. Create the MySQL Database and Table**

If you are using MySQL locally instead of Docker, run the following SQL commands to create the database and table:

    CREATE DATABASE GoTimeZoneAPI;

    USE GoTimeZoneAPI;

    CREATE TABLE time_log (
    id INT AUTO_INCREMENT PRIMARY KEY,
    timestamp DATETIME NOT NULL
    );

The time_log table stores the timestamp for each API request.

## Database Setup

**1. Running Without Docker**

To run the Go application locally (without Docker), follow these steps:

1 . Make sure MySQL is running locally, and your database and table are set up.

2 . Modify the main.go file’s database connection string with your local MySQL credentials (if necessary).

3 . Run the application:

    go run main.go

The server will start on port 8080. You can visit http://localhost:8080/time to see the current time in Toronto.

## Docker Setup

**1. Build the Docker Imagesr**

To build and run both the Go application and the MySQL database using Docker, follow these steps:

1 . Make sure you have the Dockerfile, docker-compose.yml, and the main.go file in your project directory.

2 . Build the containers:

    docker build -t go-timezone-api:v.2 .

This command will build and start both the Go application and MySQL containers.

**2. Access the API**

Once the application is running, visit the following URL in your browser or use a tool like curl to test the API:

     http://localhost:8080/time

## Error Handling

The application includes basic error handling, including:

- connection failures.
- Time zone loading errors.
- API request failures.
