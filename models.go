package main

import "errors"

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

type ContinentModel struct {
	continentCode ContinentCode
	continentName string
}

type ContinentMapModel struct {
	countryCode   CountryCode
	continentCode ContinentCode
}

type PerCapitalModel struct {
	countryCode  CountryCode
	year         Year
	gdpPerCapita float64
}
