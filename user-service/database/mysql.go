package database

import (
	"fmt"
	"log"
	"database/sql"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	fmt.Printf("Connecting with: user=%s, host=%s, port=%s, dbname=%s\n", user, host, port, name)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, name,
	)

	database, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	if err := database.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
	}

	fmt.Println("Connected to MySQL database")
	return database
}
