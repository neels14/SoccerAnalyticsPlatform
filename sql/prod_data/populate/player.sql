-- name: prod-data-player
INSERT INTO Player
SELECT DISTINCT WorldCupPlayersCSV.Player_Name  AS name, WorldCupPlayersCSV.Shirt_Number AS shirt_number, Country.name AS country
FROM WorldCupPlayersCSV, Country
WHERE WorldCupPlayersCSV.Team_Initials = Country.team_initials;
