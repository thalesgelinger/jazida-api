-- name: CreateLoad :exec
INSERT INTO loads (client, plate, material, quantity, paymentmethod, signature) 
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetLoads :many
SELECT 
    id,
    client,    
    plate,    
    material,
    quantity,     
    paymentmethod,
    signature
FROM loads;
