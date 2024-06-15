-- name: GetClients :many
SELECT 
    name,    
    plate    
FROM clients;

-- name: AddClient :exec
INSERT INTO clients (name,plate) 
VALUES ($1,$2);
