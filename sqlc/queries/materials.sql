-- name: GetMaterials :many
SELECT 
    name
FROM materials;

-- name: AddMaterial :exec
INSERT INTO materials (name) 
VALUES (?);

