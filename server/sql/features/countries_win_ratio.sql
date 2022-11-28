-- Countries win ratio
-- name: View-Country-Goal-Per-Match
CREATE OR REPLACE VIEW CountryGoalsPerMatch AS
SELECT year, stage, home_team_country, away_team_country, player_country, SUM(goals_scored) AS match_goals FROM PlayerPlaysInMatch
GROUP BY 1,2,3,4,5
ORDER BY 1,2,3,4;
-- name: View-Country-Winning-Matches
CREATE OR REPLACE VIEW CountryWinningMatches AS
SELECT c1.player_country AS country, COUNT(*) AS num_wins
FROM CountryGoalsPerMatch AS c1, CountryGoalsPerMatch AS c2
WHERE c1.player_country != c2.player_country AND
        c1.match_goals > c2.match_goals AND
        c1.year = c2.year AND
        c1.stage = c2.stage AND
        c1.home_team_country = c2.home_team_country AND
        c1.away_team_country = c2.away_team_country
GROUP BY 1;

-- name: Country-Win-Ratio
SELECT CountryWinningMatches.country AS country, CountryWinningMatches.num_wins / CountryTotalAppearances.num_matches * 100 AS win_ratio
FROM CountryWinningMatches, CountryTotalAppearances
WHERE CountryWinningMatches.country = CountryTotalAppearances.country_name
LIMIT ? OFFSET ?;

-- name: Count-Country-Win-Ratio
SELECT COUNT(*)
FROM CountryWinningMatches, CountryTotalAppearances
WHERE CountryWinningMatches.country = CountryTotalAppearances.country_name


-- name: Country-Win-Ratio-specific-country
SELECT CountryWinningMatches.country AS country, CountryWinningMatches.num_wins / CountryTotalAppearances.num_matches * 100 AS win_ratio
FROM CountryWinningMatches, CountryTotalAppearances
WHERE CountryWinningMatches.country = CountryTotalAppearances.country_name AND CountryWinningMatches.country = ?
LIMIT ? OFFSET ?;


-- name: Count-Country-Win-Ratio-specific-country
SELECT COUNT(*)
FROM CountryWinningMatches, CountryTotalAppearances
WHERE CountryWinningMatches.country = CountryTotalAppearances.country_name AND CountryWinningMatches.country = ?