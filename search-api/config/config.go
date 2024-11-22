package config

import (
	"database/sql"
	"fmt"
	"os"

	"search-api/internal/infrastructure/database"

	"github.com/joho/godotenv"
)

func GetDSN() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
}
func ConnectDb() *sql.DB {
	// Database connection
	dsn := GetDSN()
	db := database.ConnectPostgres(dsn)

	return db
}
