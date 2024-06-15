-- name: CreateLoad :exec
INSERT INTO loads (client,plate,material,quantity,paymentmethod,signature) 
VALUES ($1,$2,$3,$4,$5,$6);

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
