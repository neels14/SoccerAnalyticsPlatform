INSERT INTO PlayerPlaysInMatch
SELECT DISTINCT Players.Player_Name AS player_name, Players.Shirt_Number AS shirt_number,
                Country.name AS player_country, Matches.Year AS year,
                Matches.Stage AS stage, Matches.Home_Team_Name AS home_team_country,
                Matches.Away_Team_Name AS away_team_country, IF(Players.Lineup='S', 1, 0) AS is_starter,
                IF(Players.Event IS NULL, 0, LENGTH(Players.Event)-LENGTH(REPLACE(Players.Event, ' ', ''))+1) AS goals_scored
FROM
WorldCupPlayersCSV AS Players, WorldCupMatchesCSV AS Matches, Country
WHERE Players.MatchID = Matches.MatchID
AND Players.Team_Initials = Country.team_initials;