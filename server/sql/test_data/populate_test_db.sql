-- Insert sample mock data here as needed to validate your feature sql query works correctly.
-- name: test-data-into-countries
INSERT INTO Country
VALUES
    ('Canada', 'CAN'),
    ('United States of America', 'USA'),
    ('Brazil', 'BRA'),
    ('Argentina', 'ARG');
-- name: test-data-into-players
INSERT INTO Player
VALUES
    ('Messi', 7, 'Argentina'),
    ('Aashrit', 24, 'Canada'),
    ('Anu', 11, 'Canada'),
    ('Neel', 21, 'Canada'),
    ('Hamza', 3, 'Brazil'),
    ('Neymar', 1, 'Brazil'),
    ('Romero', 8, 'Argentina'),
    ('Paul', 9, 'United States of America'),
    ('Jacob', 17, 'United States of America');
-- name: test-data-into-world-cup
INSERT INTO WorldCup
VALUES
    (2000, 'Canada', 'Canada', 'United States of America', 'Brazil'),
    (2004, 'United States of America', 'Canada', 'Argentina', 'Brazil'),
    (2008, 'Argentina', 'Argentina', 'Canada', 'Brazil'),
    (2012, 'Argentina', 'Canada', 'Argentina', 'Brazil');
-- name: test-data-into-match
INSERT INTO SoccerMatch
VALUES
    (2000, 'Group A', 'Canada', 'United States of America', '2000-01-01', 'Toronto', 'Rogers Centre', 30000),
    (2000, 'Group A', 'Canada', 'Brazil', '2000-01-01', 'Vancouver', 'Bell Arena', 40000),
    (2000, 'Group C', 'United States of America', 'Brazil', '2000-01-01', 'Calgary', 'Telus Arena', 12000),
    (2000, 'Group C', 'Argentina', 'Brazil', '2000-01-01', 'Edmonton', 'Telus Arena', 21000),
    (2000, 'Group B', 'Argentina', 'Brazil', '2000-01-01', 'Edmonton', 'Telus Arena', 21000),
    (2004, 'Group A', 'Canada', 'United States of America', '2000-01-01', 'Toronto', 'Rogers Centre', 30000),
    (2004, 'Group A', 'Canada', 'Brazil', '2000-01-01', 'Vancouver', 'Bell Arena', 40000),
    (2004, 'Group C', 'United States of America', 'Brazil', '2000-01-01', 'Calgary', 'Telus Arena', 12000),
    (2004, 'Group C', 'Argentina', 'Brazil', '2000-01-01', 'Edmonton', 'Telus Arena', 21000);
-- name: test-data-into-player-in-match
INSERT INTO PlayerPlaysInMatch
VALUES
    ('Messi', 7, 'Argentina', 2000, 'Group C', 'Argentina', 'Brazil', TRUE, 2),
    ('Neymar', 1, 'Brazil', 2000, 'Group C', 'Argentina', 'Brazil', FALSE, 1),
    ('Romero', 8, 'Argentina', 2000, 'Group C', 'Argentina', 'Brazil', TRUE, 2),
    ('Hamza', 3, 'Brazil', 2000, 'Group C', 'Argentina', 'Brazil', FALSE, 3),
    ('Messi', 7, 'Argentina', 2000, 'Group B', 'Argentina', 'Brazil', FALSE, 2),
    ('Neymar', 1, 'Brazil', 2000, 'Group B', 'Argentina', 'Brazil', FALSE, 1),
    ('Romero', 8, 'Argentina', 2000, 'Group B', 'Argentina', 'Brazil', TRUE, 2),
    ('Hamza', 3, 'Brazil', 2000, 'Group B', 'Argentina', 'Brazil', FALSE, 3),
    ('Messi', 7, 'Argentina', 2004, 'Group C', 'Argentina', 'Brazil', TRUE, 2),
    ('Neymar', 1, 'Brazil', 2004, 'Group C', 'Argentina', 'Brazil', TRUE, 1),
    ('Romero', 8, 'Argentina', 2004, 'Group C', 'Argentina', 'Brazil', TRUE, 2),
    ('Hamza', 3, 'Brazil', 2004, 'Group C', 'Argentina', 'Brazil', FALSE, 4),
    ('Aashrit', 24, 'Canada', 2004, 'Group A', 'Canada', 'United States of America', TRUE, 5),
    ('Jacob', 17, 'United States of America', 2004, 'Group A', 'Canada', 'United States of America', TRUE, 1);

