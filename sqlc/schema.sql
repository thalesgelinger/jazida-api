CREATE TABLE IF NOT EXISTS clients (
    id SERIAL PRIMARY KEY,
    name    VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS plates (
    id SERIAL PRIMARY KEY,
    client_id INTEGER NOT NULL,
    plate VARCHAR(50) NOT NULL,
    FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS materials (
    id SERIAL PRIMARY KEY,
    name    VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS loads (
    id SERIAL PRIMARY KEY,
    client_id    INTEGER NOT NULL,   
    plate_id    INTEGER NOT NULL,   
    material_id     INTEGER NOT NULL,
    quantity     VARCHAR(50) NOT NULL,
    payment_method VARCHAR(50) NOT NULL,
    signature    VARCHAR(100) NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE,
    FOREIGN KEY (plate_id) REFERENCES plates(id) ON DELETE CASCADE,
    FOREIGN KEY (material_id) REFERENCES materials(id) ON DELETE CASCADE,
    CHECK (payment_method IN ('CASH', 'INSTALLMENT'))
);
