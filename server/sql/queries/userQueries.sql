-- name: get-user
SELECT * FROM User
WHERE username = ?;

-- name: insert-user
INSERT INTO User VALUES (?, ?, ?, ?, MD5(?));
