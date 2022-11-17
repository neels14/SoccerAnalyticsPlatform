-- name: ddl-player-play-set-foreign-0
SET FOREIGN_KEY_CHECKS=0;
-- name: ddl-player-drop-table-if-exists
DROP TABLE IF EXISTS PlayerPlaysInMatch;
-- name: ddl-player-play-set-foreign-1
SET FOREIGN_KEY_CHECKS=1;
-- name: ddl-player-play-create-table
CREATE TABLE IF NOT EXISTS PlayerPlaysInMatch (
    player_name VARCHAR(64),
    shirt_number INT,
    player_country VARCHAR(64),
    year YEAR,
    stage VARCHAR(64),
    home_team_country VARCHAR(64),
    away_team_country VARCHAR(64),
    is_starter INT NOT NULL,
    goals_scored INT NOT NULL,
    PRIMARY KEY(player_name, shirt_number, player_country, year, stage, home_team_country, away_team_country),
    FOREIGN KEY(player_name, shirt_number, player_country) REFERENCES Player(name, shirt_number, country),
    FOREIGN KEY(year, stage, home_team_country, away_team_country) REFERENCES SoccerMatch(year, stage, home_team_country, away_team_country)
) ENGINE=INNODB;