-- Most popular world cups by attendance

CREATE OR REPLACE VIEW WorldCupAttendances AS
SELECT year, SUM(attendance) as year_attendance
 FROM SoccerMatch
 GROUP BY year;

SELECT * 
FROM WorldCupAttendances
WHERE year_attendance >= ALL(SELECT year_attendance FROM WorldCupAttendances)
ORDER BY  year_attendance;
