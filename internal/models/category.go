package models

import "time"

type Category struct {
	ID        uint64
	Title     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
