-- Top Winning Countries (First Place)
-- name: top-winning-countries-first-view
CREATE OR REPLACE VIEW FirstPlaceWinners AS
SELECT first_place_country AS country_name, COUNT(*) AS total_wins FROM WorldCup
GROUP BY country_name
ORDER BY total_wins DESC, country_name;
-- name: top-winning-countries-first
SELECT * FROM FirstPlaceWinners
WHERE country_name = ?;