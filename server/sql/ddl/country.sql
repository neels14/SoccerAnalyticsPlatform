-- name: ddl-country-set-foreign-0
SET FOREIGN_KEY_CHECKS=0;
-- name: ddl-country-drop-if-exists
DROP TABLE IF EXISTS Country;
-- name: ddl-country-set-foreign-1
SET FOREIGN_KEY_CHECKS=1;
-- name: ddl-country-create-country
CREATE TABLE IF NOT EXISTS Country (
    name VARCHAR(64) PRIMARY KEY,
    team_initials CHAR(3) NOT NULL UNIQUE
) ENGINE=INNODB;