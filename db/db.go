package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// Initialize the database connection
func InitDB() {
	var err error
	// Update connection string with your database details
	connStr := "user=savanah password=Kanarada@22 dbname=customers_savannah_db sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Verify that the database is reachable
	if err := DB.Ping(); err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Create the customers table if it doesn't exist
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS customers (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			code VARCHAR(20) NOT NULL,
			phone_number VARCHAR(20) NOT NULL
		);
	`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	} else {
		log.Println("Table 'customers' is ready.")
	}

	// Create the orders table if it doesn't exist
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS orders (
			id SERIAL PRIMARY KEY,
			customer_id INTEGER NOT NULL,
			item VARCHAR(100) NOT NULL,
			amount NUMERIC(10, 2) NOT NULL,
			FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE
		);
	`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	} else {
		log.Println("Table 'orders' is ready.")
	}
}
