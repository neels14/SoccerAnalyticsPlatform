-- Top Winning Countries (1st Place)
SELECT first_place_country, COUNT(*) FROM WorldCup
GROUP BY first_place_country
ORDER BY 2 DESC;

-- Top Winning Countries (Podium)
-- TODO