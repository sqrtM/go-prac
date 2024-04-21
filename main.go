package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	connectToDb()
	countries := readCountries()
	fmt.Println(countries)
}

func readCountries() []CountryModel {
	filename := fmt.Sprintf("./materials/csv/countries.csv")

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("%s was not closed gracefully. %s\n", filename, err)
		}
	}(f)

	reader := csv.NewReader(f)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var countryList []CountryModel

	for i, line := range data {
		if i > 0 { // omit header line
			var rec CountryModel
			for j, field := range line {
				if j == 0 {
					code, err := NewCountryCode(field)
					if err != nil {
						rec.countryCode, _ = NewCountryCode("XXX")
					} else {
						rec.countryCode = code
					}
				} else if j == 1 {
					rec.countryName = field
				}
			}
			countryList = append(countryList, rec)
		}
	}
	return countryList
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
