-- name: prod-data-match
INSERT INTO SoccerMatch
SELECT DISTINCT Year AS year, Stage AS stage, Home_Team_Name AS home_team_country, Away_Team_Name AS away_team_country, Datetime AS date, City AS city, Stadium AS stadium, Attendance AS attendance FROM WorldCupMatchesCSV;