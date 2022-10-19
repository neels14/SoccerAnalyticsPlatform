DROP TABLE IF EXISTS Player;

CREATE TABLE IF NOT EXISTS Player (
    name VARCHAR(64),
    shirt_number INT,
    country VARCHAR(64) REFERENCES Country(name),
    PRIMARY KEY(name, shirt_number, country)
)