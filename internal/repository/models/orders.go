package repository

import (
	"time"
)

// Order is a struct that represent the order model
type Order struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Status     string    `gorm:"column:status;not null;default:pending" json:"status"`
	RentalType string    `gorm:"column:rental_type;not null;default:monthly" json:"rental_type"`
	PickupDate time.Time `gorm:"column:pickup_date" json:"pickup_date"`
	ReturnDate time.Time `gorm:"column:return_date" json:"return_date"`
	TotalPrice string    `gorm:"column:total_price" json:"total_price"`
	RentalDays int32     `gorm:"column:rental_days" json:"rental_days"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
	VehicleID  int32     `gorm:"column:vehicle_id" json:"vehicle_id"`
}
