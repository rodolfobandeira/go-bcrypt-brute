package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLConnection - Connects to MySQL
func MySQLConnection() *sql.DB {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
