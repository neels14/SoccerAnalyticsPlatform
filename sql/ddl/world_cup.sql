DROP TABLE IF EXISTS WorldCup;

CREATE TABLE IF NOT EXISTS WorldCup (
    year YEAR PRIMARY KEY,
    host_country VARCHAR(64) NOT NULL REFERENCES Country(name),
    first_place_country VARCHAR(64) NOT NULL REFERENCES Country(name),
    second_place_country VARCHAR(64) NOT NULL REFERENCES Country(name),
    third_place_country VARCHAR(64) NOT NULL REFERENCES Country(name)
)