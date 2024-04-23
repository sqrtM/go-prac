CREATE DOMAIN COUNTRY_CODE AS VARCHAR(3);
CREATE DOMAIN CONTINENT_CODE AS VARCHAR(2);
CREATE DOMAIN FISCAL_YEAR AS INT CONSTRAINT check_year_range CHECK (VALUE >= 2004 AND VALUE <= 2012);

CREATE TABLE IF NOT EXISTS countries
(
    country_code COUNTRY_CODE PRIMARY KEY NOT NULL,
    country_name VARCHAR(64)              NOT NULL
);

CREATE TABLE IF NOT EXISTS continents
(
    continent_code CONTINENT_CODE PRIMARY KEY NOT NULL,
    continent_name VARCHAR(32)                NOT NULL
);

CREATE TABLE IF NOT EXISTS continent_map
(
    country_code   COUNTRY_CODE REFERENCES countries (country_code)      NOT NULL,
    continent_code CONTINENT_CODE REFERENCES continents (continent_code) NOT NULL,
    PRIMARY KEY (country_code, continent_code)
);

CREATE TABLE IF NOT EXISTS per_capita
(
    country_code   COUNTRY_CODE PRIMARY KEY REFERENCES countries (country_code) NOT NULL,
    year_column    FISCAL_YEAR                                                  NOT NULL,
    gdp_per_capita FLOAT
);
