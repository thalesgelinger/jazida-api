-- name: GetClients :many
SELECT 
    c.id,
    c.name
FROM clients c;

-- name: AddClient :exec
INSERT INTO clients (name) 
VALUES (?);

