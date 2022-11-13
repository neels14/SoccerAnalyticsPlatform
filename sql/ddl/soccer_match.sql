-- name: ddl-match-set-foreign-0
SET FOREIGN_KEY_CHECKS=0;
-- name: ddl-match-drop-if-exists
DROP TABLE IF EXISTS SoccerMatch;
-- name: ddl-match-set-foreign-1
SET FOREIGN_KEY_CHECKS=1;
-- name: ddl-match-create-matches
CREATE TABLE IF NOT EXISTS SoccerMatch (
    year YEAR REFERENCES WorldCup(year),
    stage VARCHAR(64),
    home_team_country VARCHAR(64),
    away_team_country VARCHAR(64),
    date VARCHAR(64) NOT NULL,
    city VARCHAR(64) NOT NULL,
    stadium VARCHAR(64) NOT NULL,
    attendance INT NOT NULL,
#     home_coach VARCHAR(64) NOT NULL,
#     away_coach VARCHAR(64) NOT NULL,
    PRIMARY KEY(year, stage, home_team_country, away_team_country),
    FOREIGN KEY (home_team_country) REFERENCES Country(name),
    FOREIGN KEY (away_team_country) REFERENCES Country(name)

) ENGINE=INNODB;