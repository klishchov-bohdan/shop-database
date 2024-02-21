package models

import "time"

type Order struct {
	ID         uint64
	TotalPrice float64
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}
