-- Most started players in world cup history
-- name: view-player-games-started
CREATE OR REPLACE VIEW PlayerGamesStarted AS
SELECT player_name, shirt_number, player_country, COUNT(*) AS num_starts
FROM PlayerPlaysInMatch
WHERE is_starter=1
GROUP BY 1,2,3
ORDER BY 1,2,3;

-- name: view-games-played
CREATE OR REPLACE VIEW PlayerGamesPlayed AS
SELECT player_name, shirt_number, player_country, COUNT(*) AS num_games_played
FROM PlayerPlaysInMatch
GROUP BY 1,2,3
ORDER BY 1,2,3;

-- name: player-most-started
SELECT p1.player_name, p1.shirt_number, p1.player_country, p1.num_starts / p2.num_games_played * 100 AS start_ratio
FROM
    PlayerGamesStarted AS p1, PlayerGamesPlayed AS p2
WHERE p1.player_name = p2.player_name AND
        p1.shirt_number = p2.shirt_number AND
        p1.player_country = p2.player_country
ORDER BY 4 DESC, 1,2,3
LIMIT ? OFFSET ?;

-- name: player-most-started-specific-country
SELECT p1.player_name, p1.shirt_number, p1.player_country, p1.num_starts / p2.num_games_played * 100 AS start_ratio
FROM
    PlayerGamesStarted AS p1, PlayerGamesPlayed AS p2
WHERE p1.player_name = p2.player_name AND
        p1.shirt_number = p2.shirt_number AND
        p1.player_country = p2.player_country AND
        p1.player_country = ?
ORDER BY 4 DESC, 1,2,3
LIMIT ? OFFSET ?;

-- name: count-player-most-started
SELECT COUNT(*)
FROM
    PlayerGamesStarted AS p1, PlayerGamesPlayed AS p2
WHERE p1.player_name = p2.player_name AND
        p1.shirt_number = p2.shirt_number AND
        p1.player_country = p2.player_country

-- name: count-player-most-started-specific-country
SELECT COUNT(*)
FROM
    PlayerGamesStarted AS p1, PlayerGamesPlayed AS p2
WHERE p1.player_name = p2.player_name AND
        p1.shirt_number = p2.shirt_number AND
        p1.player_country = p2.player_country AND
        p1.player_country = ?
