package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var db *sql.DB
var cfg mysql.Config

func init() {
	dbName := os.Getenv("MYSQL_DATABASE")
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASS")
	dbAddr := fmt.Sprintf("%s:%s", os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"))

	cfg = mysql.Config{
		User:      dbUser,
		Passwd:    dbPass,
		Net:       "tcp",
		Addr:      dbAddr,
		DBName:    dbName,
		ParseTime: true,
	}
}

func testInit() {
	dbName := os.Getenv("MYSQL_DATABASE")
	dbTestName := fmt.Sprintf("%s_test", dbName)
	cfg.MultiStatements = true
	cfg.DBName = dbTestName
}

func ConnectDB() {
	var openErr error

	db, openErr = sql.Open("mysql", cfg.FormatDSN())
	if openErr != nil {
		log.Fatal(openErr)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Printf("Connected to %s database!\n", cfg.DBName)
}

func DisconnectDB() {
	closeErr := db.Close()
	if closeErr != nil {
		log.Fatal(closeErr)
	}
	log.Printf("Disconnected from %s database!\n", cfg.DBName)
}
