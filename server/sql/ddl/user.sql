-- name: ddl-user-set-foreign-0
SET FOREIGN_KEY_CHECKS=0;
-- name: ddl-user-drop-if-exists
DROP TABLE IF EXISTS User;
-- name: ddl-user-set-foreign-1
SET FOREIGN_KEY_CHECKS=1;
-- name: ddl-user-create-user
CREATE TABLE IF NOT EXISTS User (
    username VARCHAR(64) PRIMARY KEY,
    first_name VARCHAR(64) NOT NULL,
    last_name VARCHAR(64) NOT NULL,
    email VARCHAR(64) NOT NULL UNIQUE,
    password VARCHAR(200) NOT NULL
) ENGINE=INNODB;
