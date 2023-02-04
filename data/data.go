// Main handler for database and CRUD operations
package data

import (
	"fmt"
	"log"
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

var BPDatabase *sql.DB

const BETA_DB_NAME = "blood_pressure"

func Connect() {
	config := mysql.Config{
		User:   "root",
		Passwd: "",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: BETA_DB_NAME,
		AllowNativePasswords: true,
	}

	var openErr error
	BPDatabase, openErr = sql.Open("mysql", config.FormatDSN())
	if openErr != nil {
		log.Fatal(openErr)
	}

	if pingErr := BPDatabase.Ping(); pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected to database!")
}
