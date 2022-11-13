CREATE OR REPLACE VIEW PlayerGamesStarted AS
SELECT player_name, shirt_number, player_country, COUNT(*) AS num_starts
FROM PlayerPlaysInMatch
WHERE is_starter=1
GROUP BY 1,2,3
ORDER BY 1,2,3;

CREATE OR REPLACE VIEW PlayerGamesPlayed AS
SELECT player_name, shirt_number, player_country, COUNT(*) AS num_games_played
FROM PlayerPlaysInMatch
GROUP BY 1,2,3
ORDER BY 1,2,3;

SELECT p1.player_name, p1.shirt_number, p1.player_country, p1.num_starts / p2.num_games_played * 100 AS start_ratio
FROM
    PlayerGamesStarted AS p1, PlayerGamesPlayed AS p2
WHERE p1.player_name = p2.player_name AND
        p1.shirt_number = p2.shirt_number AND
        p1.player_country = p2.player_country
ORDER BY 4 DESC, 1,2,3;
