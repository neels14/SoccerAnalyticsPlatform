-- name: prod-data-world-cup
INSERT INTO WorldCup
SELECT DISTINCT Year AS year, Country AS host_country, Winner AS first_place_country, RunnersUp AS second_place_country, Third AS third_place_country
FROM WorldCupsCSV;
