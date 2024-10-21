-- name: GetMaterials :many
SELECT 
    id, 
    name
FROM materials;

-- name: AddMaterial :exec
INSERT INTO materials (name) 
VALUES ($1);

-- name: GetMaterialById :one
SELECT name
FROM materials
WHERE id = $1;

