-- name: ddl-player-set-foreign-0
SET FOREIGN_KEY_CHECKS=0;
-- name: ddl-player-drop-if-exists
DROP TABLE IF EXISTS Player;
-- name: ddl-player-set-foreign-1
SET FOREIGN_KEY_CHECKS=1;
-- name: ddl-player-create-player
CREATE TABLE IF NOT EXISTS Player (
    name VARCHAR(64),
    shirt_number INT,
    country VARCHAR(64),
    PRIMARY KEY(name, shirt_number, country),
    FOREIGN KEY (country) REFERENCES Country(name)
) ENGINE=INNODB;