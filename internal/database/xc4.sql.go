// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: xc4.sql

package database

import (
	"context"
)

const getAllDevices = `-- name: GetAllDevices :many
SELECT id, device_type, status, created_at, updated_at FROM devices
`

func (q *Queries) GetAllDevices(ctx context.Context) ([]Device, error) {
	rows, err := q.db.QueryContext(ctx, getAllDevices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Device
	for rows.Next() {
		var i Device
		if err := rows.Scan(
			&i.ID,
			&i.DeviceType,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
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
