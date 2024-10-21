-- name: AddPlate :exec
INSERT INTO plates (client_id, plate) 
VALUES ($1, $2);

-- name: GetPlatesByClientId :many
SELECT id, plate 
FROM plates 
WHERE client_id = $1;

-- name: GetPlateById :one
SELECT plate
FROM plates
WHERE id = $1;

