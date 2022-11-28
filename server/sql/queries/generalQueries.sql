-- name: get-all-countries
SELECT * FROM Country
LIMIT ? OFFSET ?;

-- name: count-countries
SELECT COUNT(*) FROM Country

-- name: get-all-players
SELECT * FROM Player
LIMIT ? OFFSET ?;

-- name: count-players
SELECT COUNT(*) FROM Player

-- name: get-all-wc
SELECT * FROM WorldCup
LIMIT ? OFFSET ?;

-- name: count-wc
SELECT COUNT(*) FROM WorldCup;

-- name: get-all-matches
SELECT * FROM SoccerMatch
LIMIT ? OFFSET ?;

-- name: count-matches
SELECT COUNT(*) FROM SoccerMatch;