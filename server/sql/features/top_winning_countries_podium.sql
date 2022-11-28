-- Top Winning Countries (Podium)
-- name: top-winning-countries-first-view-podium
CREATE OR REPLACE VIEW FirstPlaceWinners AS
SELECT first_place_country AS country_name, COUNT(*) AS num_wins FROM WorldCup
GROUP BY country_name
ORDER BY num_wins DESC;
-- name: top-winning-countries-second-view-podium
CREATE OR REPLACE VIEW SecondPlaceWinners AS
SELECT second_place_country AS country_name, COUNT(*) AS num_wins FROM WorldCup
GROUP BY country_name
ORDER BY num_wins DESC;
-- name: top-winning-countries-third-view-podium
CREATE OR REPLACE VIEW ThirdPlaceWinners AS
SELECT third_place_country AS country_name, COUNT(*) AS num_wins FROM WorldCup
GROUP BY country_name
ORDER BY num_wins DESC;
-- name: top-winning-countries-podium-with-country-name
SELECT country_name, SUM(num_wins) AS total_wins FROM (
    SELECT * FROM FirstPlaceWinners
    UNION
    SELECT * FROM SecondPlaceWinners
    UNION
    SELECT * FROM ThirdPlaceWinners) AS podium_winners
WHERE country_name = ?
GROUP BY country_name
ORDER BY total_wins DESC, country_name
LIMIT ? OFFSET ?;

-- name: top-winning-countries-podium
SELECT country_name, SUM(num_wins) AS total_wins FROM (
    SELECT * FROM FirstPlaceWinners
    UNION
    SELECT * FROM SecondPlaceWinners
    UNION
    SELECT * FROM ThirdPlaceWinners) AS podium_winners
GROUP BY country_name
ORDER BY total_wins DESC, country_name
LIMIT ? OFFSET ?;

-- name: count-top-winning-countries-podium-with-country-name
SELECT COUNT(*) FROM (
    SELECT * FROM FirstPlaceWinners
    UNION
    SELECT * FROM SecondPlaceWinners
    UNION
    SELECT * FROM ThirdPlaceWinners) AS podium_winners
WHERE country_name = ?
GROUP BY country_name

-- name: count-top-winning-countries-podium
SELECT COUNT(*) FROM (
    SELECT * FROM FirstPlaceWinners
    UNION
    SELECT * FROM SecondPlaceWinners
    UNION
    SELECT * FROM ThirdPlaceWinners) AS podium_winners
GROUP BY country_name;


