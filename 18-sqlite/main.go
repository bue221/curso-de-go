package main

import (
	"database/sql"
	"log"

	"sql-in-go/database"

	_ "github.com/mattn/go-sqlite3"
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		return nil, err
	}

	createTables(db)

	return db, nil
}

func createTables(db *sql.DB) error {
	orderRepository := &database.OrderRepository{Db: db}
	err := orderRepository.CreateTable()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	dbConnection, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}

	defer dbConnection.Close()

	orderRepository := &database.OrderRepository{Db: dbConnection}
	orderRepository.CreateOrder(database.Order{Product: "Laptop", Amount: 1000})

	
}