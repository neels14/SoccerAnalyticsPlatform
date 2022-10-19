-- Insert sample mock data here as needed to validate your feature sql query works correctly.

INSERT INTO Country
VALUES
    ('Canada', 'CAN'),
    ('United States of America', 'USA'),
    ('Brazil', 'BRA'),
    ('Argentina', 'ARG');

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

INSERT INTO WorldCup
VALUES
    (2000, 'Canada', 'Canada', 'United States of America', 'Brazil'),
    (2004, 'United States of America', 'Canada', 'Argentina', 'Brazil'),
    (2008, 'Argentina', 'Argentina', 'Canada', 'Brazil'),
    (2012, 'Argentina', 'Canada', 'Argentina', 'Brazil');

INSERT INTO SoccerMatch
VALUES
    (2000, 'Group A', 'Canada', 'United States of America', '2000-01-01', 'Toronto', 'Rogers Centre', 30000, 'Mr. Jeff', 'Mr. Pauliza'),
    (2000, 'Group A', 'Canada', 'Brazil', '2000-01-01', 'Vancouver', 'Bell Arena', 40000, 'Mr. Jeff', 'Mr. Pauliza'),
    (2000, 'Group C', 'United States of America', 'Brazil', '2000-01-01', 'Calgary', 'Telus Arena', 12000, 'Mr. Jeff', 'Mr. Pauliza'),
    (2000, 'Group C', 'Argentina', 'Brazil', '2000-01-01', 'Edmonton', 'Telus Arena', 21000, 'Mr. Jeff', 'Mr. Pauliza');

INSERT INTO PlayerPlaysInMatch
VALUES
    ('Messi', 7, 'Argentina', 2000, 'Group C', 'Argentina', 'Brazil', TRUE, 2),
    ('Neymar', 1, 'Brazil', 2000, 'Group C', 'Argentina', 'Brazil', TRUE, 1),
    ('Romero', 8, 'Argentina', 2000, 'Group C', 'Argentina', 'Brazil', TRUE, 0),
    ('Hamza', 3, 'Brazil', 2000, 'Group C', 'Argentina', 'Brazil', TRUE, 3);