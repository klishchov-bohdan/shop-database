package models

import "time"

type Product struct {
	ID          uint64
	CategoryID  uint64
	Title       string
	Description string
	Price       float64
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
