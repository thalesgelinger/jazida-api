-- name: CreateLoad :exec
INSERT INTO loads (client_id, plate_id, material_id, quantity, payment_method, signature) 
VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetLoads :many
SELECT 
    l.id,
    c.name AS client,    
    p.plate AS plate,    
    m.name AS material,
    l.quantity,     
    l.payment_method,
    l.signature
FROM loads l
JOIN clients c ON l.client_id = c.id
JOIN plates p ON l.plate_id = p.id
JOIN materials m ON l.material_id = m.id;

