// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: materials.sql

package db

import (
	"context"
)

const addMaterial = `-- name: AddMaterial :exec
INSERT INTO materials (name) 
VALUES (?)
`

func (q *Queries) AddMaterial(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, addMaterial, name)
	return err
}

const getMaterials = `-- name: GetMaterials :many
SELECT 
    id, 
    name
FROM materials
`

func (q *Queries) GetMaterials(ctx context.Context) ([]Material, error) {
	rows, err := q.db.QueryContext(ctx, getMaterials)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Material
	for rows.Next() {
		var i Material
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
