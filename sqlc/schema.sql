CREATE TABLE IF NOT EXISTS loads (
    id SERIAL PRIMARY KEY,
    client    VARCHAR(50) NOT NULL,   
    plate      VARCHAR(50) NOT NULL,  
    material     VARCHAR(50) NOT NULL,
    quantity     VARCHAR(50) NOT NULL,
    paymentmethod VARCHAR(50) NOT NULL,
    signature    VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS clients (
    id SERIAL PRIMARY KEY,
    name    VARCHAR(50) NOT NULL,   
    plate      VARCHAR(50) NOT NULL  
);
