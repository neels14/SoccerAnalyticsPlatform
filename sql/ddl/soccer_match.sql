DROP TABLE IF EXISTS SoccerMatch;

CREATE TABLE IF NOT EXISTS SoccerMatch (
    year YEAR REFERENCES WorldCup(year),
    stage VARCHAR(64),
    home_team_country VARCHAR(64) REFERENCES Country(name),
    away_team_country VARCHAR(64) REFERENCES Country(name),
    datetime DATETIME NOT NULL,
    city VARCHAR(64) NOT NULL,
    stadium VARCHAR(64) NOT NULL,
    attendance INT NOT NULL,
    home_coach VARCHAR(64) NOT NULL,
    away_coach VARCHAR(64) NOT NULL,
    PRIMARY KEY(year, stage, home_team_country, away_team_country)
)