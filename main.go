package main

import (
	"database/sql"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"os"
)

func main() {

	connectToDb()

}

func connectToDb() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "db", 5432, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	fmt.Println(psqlconn)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("Database was not closed gracefully. %s\n", err)
		}
	}(db)

	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
