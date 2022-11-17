-- name: ddl-world-cup-set-foreign-0
SET FOREIGN_KEY_CHECKS=0;
-- name: ddl-world-cup-drop-if-exists
DROP TABLE IF EXISTS WorldCup;
-- name: ddl-world-cup-set-foreign-0
SET FOREIGN_KEY_CHECKS=1;
-- name: ddl-world-cup-create-table
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