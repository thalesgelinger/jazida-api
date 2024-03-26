package db

import (
	"database/sql"
	"fmt"
	"jazida-api/entities"
	"log"

	_ "github.com/lib/pq"
)

func OpenConn() (*sql.DB, error) {
	const (
		host     = "localhost"
		port     = 5432
		user     = "jazida"
		password = "jazida123"
		dbname   = "jazidadb"
	)

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error conecting database")
	}

	return db, err
}

func CreateLoadsTable() {
	db, err := OpenConn()
	if err != nil {
		return
	}
	defer db.Close()

	createTableSQL := `
        CREATE TABLE IF NOT EXISTS loads (
            id SERIAL PRIMARY KEY,
            client    VARCHAR(50) NOT NULL,   
            plate      VARCHAR(50) NOT NULL,  
            material     VARCHAR(50) NOT NULL,
            quantity     VARCHAR(50) NOT NULL,
            paymentmethod VARCHAR(50) NOT NULL,
            signature    VARCHAR(100) NOT NULL
        )
    `

	_, err = db.Exec(createTableSQL)

	if err != nil {
		log.Fatal("Error creating table: ", err)
	}

	log.Println("Table created successfully")
}

func SaveLoad(load *entities.Load) {
	db, err := OpenConn()

	if err != nil {
		return
	}

	defer db.Close()

	var loadId int

	sql := `INSERT INTO loads (
		        client
		       ,plate
		       ,material
		       ,quantity
		       ,paymentmethod
		       ,signature
		    ) VALUES ($1,$2,$3,$4,$5,$6)
            RETURNING id
            `

	err = db.QueryRow(sql,
		load.Client,
		load.Plate,
		load.Material,
		load.Quantity,
		load.PaymentMethod,
		load.Signature,
	).Scan(&loadId)

	if err != nil {
		log.Fatal("Error creating new load", err)
	}
}

func GetLoads() ([]entities.Load, error) {
	db, err := OpenConn()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	query := `SELECT 
            client,    
            plate,    
            material,
            quantity,     
            paymentmethod,
            signature
            FROM loads`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error getting loads", err)
	}
	defer rows.Close()

	loads := []entities.Load{}

	for rows.Next() {
		var load entities.Load

		if err := rows.Scan(
			&load.Client,
			&load.Plate,
			&load.Material,
			&load.Quantity,
			&load.PaymentMethod,
			&load.Signature,
		); err != nil {
			log.Fatal("Error reading load row", err)
		}

		loads = append(loads, load)
	}

	return loads, nil
}
