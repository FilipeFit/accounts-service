package domain

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	ID             uint64 `gorm:"primary_key"`
	CustomerId     uint64
	Balance        float64
	Active         bool
	AllowOverdraft bool
}
