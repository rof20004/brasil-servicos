package database

import (
	"database/sql"
	"log"
)

// DB - database connection instance
var DB *sql.DB

// InitDatabase - initialize database connection
func InitDatabase() {
	db, err := sql.Open("sqlite3", "brasil-servicos.db")
	if err != nil {
		log.Fatalln(err)
	}

	// Verify connection
	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}

	DB = db
}

// InitMigrations - create tables
func InitMigrations() {
	createTableUser := `CREATE TABLE IF NOT EXISTS users (
												id INTEGER PRIMARY KEY,
												nome TEXT NOT NULL
											);`
	_, err := DB.Exec(createTableUser)
	if err != nil {
		log.Fatalln(err)
	}
}
