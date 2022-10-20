-- Top Winning Countries (First Place)

CREATE OR REPLACE VIEW FirstPlaceWinners AS
SELECT first_place_country AS country_name, COUNT(*) AS total_wins FROM WorldCup
GROUP BY country_name
ORDER BY total_wins DESC, country_name;

SELECT * FROM FirstPlaceWinners
WHERE country_name = "Canada";

SELECT * FROM FirstPlaceWinners;