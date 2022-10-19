SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS Player;
SET FOREIGN_KEY_CHECKS=1;

CREATE TABLE IF NOT EXISTS Player (
    name VARCHAR(64),
    shirt_number INT,
    country VARCHAR(64),
    PRIMARY KEY(name, shirt_number, country),
    FOREIGN KEY (country) REFERENCES Country(name)
) ENGINE=INNODB;