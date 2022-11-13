-- name: prod-data-country
INSERT INTO Country
SELECT DISTINCT * FROM
(SELECT Home_Team_Name AS name, Home_Team_Initials AS team_initials FROM WorldCupMatchesCSV
UNION
SELECT Away_Team_Name AS name, Away_Team_Initials AS team_initials FROM WorldCupMatchesCSV) AS all_countries;