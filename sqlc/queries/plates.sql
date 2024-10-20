-- name: AddPlate :exec
INSERT INTO plates (client_id, plate) 
VALUES (?, ?);

-- name: GetPlatesByClientId :many
SELECT id, plate 
FROM plates 
WHERE client_id = ?;
