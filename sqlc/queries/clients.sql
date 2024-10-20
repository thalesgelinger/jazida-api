-- name: GetClients :many
SELECT 
    c.id,
    c.name
FROM clients c;

-- name: AddClient :exec
INSERT INTO clients (name) 
VALUES (?);

-- name: GetClientById :one
SELECT name 
FROM clients
WHERE id = ?;
