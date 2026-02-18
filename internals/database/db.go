package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/COZYTECH/PERSONALTRACKERAPI/internals/utils"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() {
	cfg := utils.LoadConfig()

	// Example DSN: "user:password@tcp(host:port)/dbname?parseTime=true"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Database ping failed:", err)
	}

	log.Println("Database connected successfully")
}
