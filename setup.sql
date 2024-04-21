CREATE TABLE IF NOT EXISTS countries
(
    country_code VARCHAR(3) PRIMARY KEY NOT NULL,
    country_name VARCHAR(64)            NOT NULL
);

CREATE TABLE IF NOT EXISTS continents
(
    continent_code VARCHAR(2) PRIMARY KEY NOT NULL,
    continent_name VARCHAR(32)            NOT NULL
);

CREATE TABLE IF NOT EXISTS continent_map
(
    country_code   VARCHAR(3) REFERENCES countries (country_code)    NOT NULL,
    continent_code VARCHAR(2) REFERENCES continents (continent_code) NOT NULL,
    PRIMARY KEY (country_code, continent_code)
);

CREATE TABLE IF NOT EXISTS per_capita
(
    country_code   VARCHAR(3) PRIMARY KEY REFERENCES countries (country_code) NOT NULL,
    year_column    INT                                                        NOT NULL,
    gdp_per_capita FLOAT                                                      NOT NULL
        CONSTRAINT check_year_range CHECK (year_column >= 2004 AND year_column <= 2006)
);
