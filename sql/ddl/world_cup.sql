SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS WorldCup;
SET FOREIGN_KEY_CHECKS=1;

CREATE TABLE IF NOT EXISTS WorldCup (
    year YEAR PRIMARY KEY,
    host_country VARCHAR(64) NOT NULL,
    first_place_country VARCHAR(64) NOT NULL,
    second_place_country VARCHAR(64) NOT NULL,
    third_place_country VARCHAR(64) NOT NULL,
    FOREIGN KEY (host_country) REFERENCES Country(name),
    FOREIGN KEY (first_place_country) REFERENCES Country(name),
    FOREIGN KEY (second_place_country) REFERENCES Country(name),
    FOREIGN KEY (third_place_country) REFERENCES Country(name)
) ENGINE=INNODB;