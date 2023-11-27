package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func GetConn() *sql.DB {
	return db
}

func ConnectToDatabase() error {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, dbhost, dbport, dbname)

	db, _ = sql.Open("mysql", connStr)

	err := db.Ping()

	if err != nil {
		log.Fatalf("Failed to ping to the database: %v", err)
		return err
	}

	return nil
}
