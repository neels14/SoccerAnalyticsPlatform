-- Average number of goals in a match for each country over all World Cups

-- name: create-view-CountryGoalsScored
CREATE OR REPLACE VIEW CountryGoalsScored AS
SELECT player_country as country_name, SUM(goals_scored) as country_goals_scored
FROM PlayerPlaysInMatch
GROUP BY player_country;

-- name: create-view-CountryTotalAppearances
CREATE OR REPLACE VIEW CountryTotalAppearances AS
SELECT country_name, SUM(country_matches) AS num_matches
FROM (
    (SELECT home_team_country as country_name, COUNT(*) as country_matches
     FROM SoccerMatch
     GROUP BY home_team_country)
    UNION ALL
    (SELECT away_team_country as country_name, COUNT(*) as country_matches
     FROM SoccerMatch
     GROUP BY away_team_country)) AS TotalMatches
GROUP BY country_name;

-- name: create-view-CountryAvgData
CREATE OR REPLACE VIEW CountryAvgData AS
SELECT
    CountryTotalAppearances.country_name AS country_name,
    IFNULL(country_goals_scored,0) AS goals_scored,
    num_matches AS matches_played
FROM CountryGoalsScored RIGHT OUTER JOIN CountryTotalAppearances
ON CountryGoalsScored.country_name = CountryTotalAppearances.country_name;

-- name: create-table-AvgGoalsInMatchByCountry
SELECT country_name, goals_scored/matches_played AS avg_goals_match
FROM CountryAvgData
ORDER BY avg_goals_match DESC, country_name;

