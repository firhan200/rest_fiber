package models

import "time"

// gorm.Model definition
type Product struct {
	ID        uint
	Name      string
	Price     int32
	CreatedAt time.Time
	UpdatedAt time.Time
}
