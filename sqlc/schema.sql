CREATE TABLE IF NOT EXISTS loads (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    client    VARCHAR(50) NOT NULL,   
    plate      VARCHAR(50) NOT NULL,  
    material     VARCHAR(50) NOT NULL,
    quantity     VARCHAR(50) NOT NULL,
    paymentmethod VARCHAR(50) NOT NULL,
    signature    VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS clients (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name    VARCHAR(50) NOT NULL,   
    plate      VARCHAR(50) NOT NULL  
);

