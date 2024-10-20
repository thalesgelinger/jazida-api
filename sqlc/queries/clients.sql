-- name: GetClients :many
SELECT 
    c.id,
    c.name,    
    p.plate    
FROM clients c
LEFT JOIN plates p ON c.id = p.client_id;

-- name: AddClient :exec
INSERT INTO clients (name) 
VALUES (?);

