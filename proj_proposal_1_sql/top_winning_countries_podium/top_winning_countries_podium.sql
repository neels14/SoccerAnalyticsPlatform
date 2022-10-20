-- Top Winning Countries (Podium)

CREATE OR REPLACE VIEW FirstPlaceWinners AS
SELECT first_place_country AS country_name, COUNT(*) AS num_wins FROM WorldCup
GROUP BY country_name
ORDER BY num_wins DESC;

CREATE OR REPLACE VIEW SecondPlaceWinners AS
SELECT second_place_country AS country_name, COUNT(*) AS num_wins FROM WorldCup
GROUP BY country_name
ORDER BY num_wins DESC;

CREATE OR REPLACE VIEW ThirdPlaceWinners AS
SELECT third_place_country AS country_name, COUNT(*) AS num_wins FROM WorldCup
GROUP BY country_name
ORDER BY num_wins DESC;

SELECT country_name, SUM(num_wins) AS total_wins FROM (
    SELECT * FROM FirstPlaceWinners
    UNION
    SELECT * FROM SecondPlaceWinners
    UNION
    SELECT * FROM ThirdPlaceWinners) AS podium_winners
WHERE country_name = "Brazil"
GROUP BY country_name
ORDER BY total_wins DESC, country_name;

SELECT country_name, SUM(num_wins) AS total_wins FROM (
    SELECT * FROM FirstPlaceWinners
    UNION
    SELECT * FROM SecondPlaceWinners
    UNION
    SELECT * FROM ThirdPlaceWinners) AS podium_winners
GROUP BY country_name
ORDER BY total_wins DESC, country_name;


