-- Top goal scorer from each country

CREATE OR REPLACE VIEW CountryGoalScorers AS
SELECT player_country as goal_scorer_country, player_name as goal_scorer_name ,SUM(goals_scored) as goal_scorer_goals
FROM PlayerPlaysInMatch
GROUP BY player_country, player_name;

CREATE OR REPLACE VIEW CountryMostGoals AS
SELECT goal_scorer_country as top_scorer_country, MAX(goal_scorer_goals) as top_scorer_goals
FROM CountryGoalScorers
GROUP BY goal_scorer_country;

SELECT top_scorer_country, goal_scorer_name, top_scorer_goals
FROM CountryGoalScorers, CountryMostGoals
WHERE top_scorer_country = goal_scorer_country AND top_scorer_goals = goal_scorer_goals AND top_scorer_country = 'Argentina'
ORDER BY top_scorer_goals DESC, top_scorer_country, goal_scorer_name;

SELECT top_scorer_country, goal_scorer_name, top_scorer_goals
FROM CountryGoalScorers, CountryMostGoals
WHERE top_scorer_country = goal_scorer_country AND top_scorer_goals = goal_scorer_goals
ORDER BY top_scorer_goals DESC, top_scorer_country, goal_scorer_name;
