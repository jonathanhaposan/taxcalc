package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Init() *sql.DB {
	host := os.Getenv("DBHOST")
	user := os.Getenv("DBUSER")
	pass := os.Getenv("DBPASS")
	dbName := os.Getenv("DBNAME")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", user, pass, dbName, host)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	log.Println("Succes connect to DB")

	return db
}
