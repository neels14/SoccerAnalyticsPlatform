-- Most popular world cups by attendance

-- name: create-view-WorldCupAttendances
CREATE OR REPLACE VIEW WorldCupAttendances AS
SELECT year, SUM(attendance) as year_attendance
 FROM SoccerMatch
 GROUP BY year;

-- name: create-table-MostPopularWCByAttendance
SELECT * 
FROM WorldCupAttendances
WHERE year_attendance >= ALL(SELECT year_attendance FROM WorldCupAttendances)
ORDER BY  year_attendance
LIMIT ? OFFSET ?;

-- name: count-create-table-MostPopularWCByAttendance
SELECT COUNT(*) 
FROM WorldCupAttendances
WHERE year_attendance >= ALL(SELECT year_attendance FROM WorldCupAttendances)
ORDER BY  year_attendance

