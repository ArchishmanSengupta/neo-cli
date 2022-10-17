package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DbConn *sqlx.DB

// Connect - extracts the datafrom the .env file and establish the database connection
// Returns - db the database connection variable, error (if any)
func Connect() (*sqlx.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbServer := os.Getenv("DB_SERVER")

	dbUrl := "postgres://" + dbUser + ":" + dbPass + "@" + dbServer + "/" + dbName + "?sslmode=disable"

	fmt.Println(dbUrl)

	db, err := sqlx.Connect("postgres", dbUrl)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return db, nil
}
