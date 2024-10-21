-- name: GetClients :many
SELECT 
    c.id,
    c.name
FROM clients c;

-- name: AddClient :exec
INSERT INTO clients (name) 
VALUES ($1);

-- name: GetClientById :one
SELECT name 
FROM clients
WHERE id = $1;

