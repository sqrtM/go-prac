package prac

import (
	"database/sql"
	"errors"
)

type CountryCode string

func NewCountryCode(val string) (CountryCode, error) {
	code := CountryCode(val)
	if !code.IsValid() {
		return "", errors.New("invalid continent code value")
	}
	return code, nil
}

func (c CountryCode) IsValid() bool {
	return len(c) == 3
}

type ContinentCode string

func NewContinentCode(val string) (ContinentCode, error) {
	code := ContinentCode(val)
	if !code.IsValid() {
		return "", errors.New("invalid continent code value")
	}
	return code, nil
}

func (c ContinentCode) IsValid() bool {
	return len(c) == 2
}

type Year int

func NewYear(val int) (Year, error) {
	year := Year(val)
	if !year.IsValid() {
		return 0, errors.New("invalid year value")
	}
	return year, nil
}

func (y Year) IsValid() bool {
	return y >= 2004 && y <= 2012
}

type CountryModel struct {
	countryCode CountryCode
	countryName string
}

func (cm CountryModel) Insert(db *sql.DB) (int64, error) {
	sqlStatement := `
		INSERT INTO countries (country_code, country_name)
		VALUES ($1, $2)`

	cc, err := db.Exec(sqlStatement, cm.countryCode, cm.countryName)
	if err != nil {
		return 0, err
	}

	rows, re := cc.RowsAffected()
	if re != nil {
		return 0, re
	}

	return rows, nil
}

type ContinentModel struct {
	continentCode ContinentCode
	continentName string
}

func (cm ContinentModel) Insert(db *sql.DB) (int64, error) {
	sqlStatement := `
		INSERT INTO continents (continent_code, continent_name)
		VALUES ($1, $2)`

	cc, err := db.Exec(sqlStatement, cm.continentCode, cm.continentName)
	if err != nil {
		return 0, err
	}

	rows, re := cc.RowsAffected()
	if re != nil {
		return 0, re
	}

	return rows, nil
}

type ContinentMapModel struct {
	countryCode   CountryCode
	continentCode ContinentCode
}

func (cm ContinentMapModel) Insert(db *sql.DB) (int64, error) {
	sqlStatement := `
		INSERT INTO continent_map (country_code, continent_code)
		VALUES ($1, $2)`

	cc, err := db.Exec(sqlStatement, cm.countryCode, cm.continentCode)
	if err != nil {
		return 0, err
	}

	rows, re := cc.RowsAffected()
	if re != nil {
		return 0, re
	}

	return rows, nil
}

type PerCapitalModel struct {
	countryCode  CountryCode
	year         Year
	gdpPerCapita float64
}

func (cm PerCapitalModel) Insert(db *sql.DB) (int64, error) {
	sqlStatement := `
		INSERT INTO per_capita (country_code, year_column, gdp_per_capita)
		VALUES ($1, $2, $3)`

	cc, err := db.Exec(sqlStatement, cm.countryCode, cm.year, cm.gdpPerCapita)
	if err != nil {
		return 0, err
	}

	rows, re := cc.RowsAffected()
	if re != nil {
		return 0, re
	}

	return rows, nil
}
