DROP TABLE IF EXISTS PlayerPlaysInMatch;

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
)