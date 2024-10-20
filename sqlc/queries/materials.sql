-- name: GetMaterials :many
SELECT 
    id, 
    name
FROM materials;

-- name: AddMaterial :exec
INSERT INTO materials (name) 
VALUES (?);

