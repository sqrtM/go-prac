package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
)

func main() {
	db := connectToDb()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Printf("Database was not closed gracefully. %s\n", err)
		}
	}(db)

	countries := readCountries()
	cont := readContinents()
	ccmap := readContinentMap()
	percap := readPerCapita()

	for i := range countries {
		insert, err := countries[i].Insert(db)
		if err != nil {
			return
		}
		if insert > 0 {
			fmt.Println("successful insert")
		}
	}

	for i := range cont {
		insert, err := cont[i].Insert(db)
		if err != nil {
			return
		}
		if insert > 0 {
			fmt.Println("successful insert")
		}
	}

	for i := range ccmap {
		insert, err := ccmap[i].Insert(db)
		if err != nil {
			return
		}
		if insert > 0 {
			fmt.Println("successful insert")
		}
	}

	for i := range percap {
		insert, err := percap[i].Insert(db)
		if err != nil {
			return
		}
		if insert > 0 {
			fmt.Println("successful insert")
		}
	}
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

func readContinents() []ContinentModel {
	filename := fmt.Sprintf("./materials/csv/continents.csv")

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

	var continentList []ContinentModel

	for i, line := range data {
		if i > 0 { // omit header line
			var rec ContinentModel
			for j, field := range line {
				if j == 0 {
					code, err := NewContinentCode(field)
					if err != nil {
						rec.continentCode, _ = NewContinentCode("XXX")
					} else {
						rec.continentCode = code
					}
				} else if j == 1 {
					fmt.Println(field)
					rec.continentName = field
				}
			}
			continentList = append(continentList, rec)
		}
	}
	return continentList
}

func readContinentMap() []ContinentMapModel {
	filename := fmt.Sprintf("./materials/csv/continent_map.csv")

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

	var mapList []ContinentMapModel

	for i, line := range data {
		if i > 0 { // omit header line
			var rec ContinentMapModel
			for j, field := range line {
				if j == 0 {
					code, err := NewCountryCode(field)
					if err != nil {
						rec.countryCode, _ = NewCountryCode("XXX")
					} else {
						rec.countryCode = code
					}
				} else if j == 1 {
					fmt.Println(field)

					code, err := NewContinentCode(field)
					if err != nil {
						rec.continentCode, _ = NewContinentCode("XXX")
					} else {
						rec.continentCode = code
					}
				}
			}
			mapList = append(mapList, rec)
		}
	}
	return mapList
}

func readPerCapita() []PerCapitalModel {
	filename := fmt.Sprintf("./materials/csv/per_capita.csv")

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

	var perCapitaList []PerCapitalModel

	for i, line := range data {
		if i > 0 { // omit header line
			var rec PerCapitalModel
			for j, field := range line {
				if j == 0 {
					code, err := NewCountryCode(field)
					if err != nil {
						rec.countryCode, _ = NewCountryCode("XXX")
					} else {
						rec.countryCode = code
					}
				} else if j == 1 {
					val, intErr := strconv.Atoi(field)
					if intErr == nil {
						yr, err := NewYear(val)
						if err != nil {
							rec.year = yr
						} else {
							rec.year = 2006
						}
					} else {
						rec.year = 2006
					}
				} else if j == 2 {
					fmt.Println(field)

					val, floatErr := strconv.ParseFloat(field, 64)
					if floatErr == nil {
						rec.gdpPerCapita = val
					} else {
						rec.gdpPerCapita = 0.0
					}
				}
			}
			perCapitaList = append(perCapitaList, rec)
		}
	}
	return perCapitaList
}

func connectToDb() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "db", 5432, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	fmt.Println(psqlconn)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
