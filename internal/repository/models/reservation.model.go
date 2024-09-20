package repository

import "database/sql"

type ReservationMapping struct {
	ReservationID        uint         `gorm:"column:reservation_id;primaryKey" json:"reservation_id"`
	LineUUID             string       `gorm:"column:lineUUID;primaryKey" json:"lineUUID"`
	CreateAt             sql.NullTime `gorm:"column:create_at;default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAt             sql.NullTime `gorm:"column:update_at;default:CURRENT_TIMESTAMP" json:"update_at"`
	Status               string       `gorm:"column:status;not null;default:pending" json:"status"`
	RentalType           string       `gorm:"column:rental_type;not null;default:monthly" json:"rental_type"`
	ExtraPeriodFeeStamp  int32        `gorm:"column:extra_period_fee_stamp" json:"extra_period_fee_stamp"`
	ExtraDepositFeeStamp int32        `gorm:"column:extra_deposit_fee_stamp" json:"extra_deposit_fee_stamp"`
	PickupDate           sql.NullTime `gorm:"column:pickup_date" json:"pickup_date"`
	ReturnDate           sql.NullTime `gorm:"column:return_date" json:"return_date"`
	TotalPrice           string       `gorm:"column:total_price" json:"total_price"`
	OutstandingBalance   string       `gorm:"column:outstanding_balance" json:"outstanding_balance"`
	RentalDays           int32        `gorm:"column:rental_days" json:"rental_days"`
}

type Fleet struct {
	ID     int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	JSON   string `gorm:"column:json" json:"json"`
	Active bool   `gorm:"column:active" json:"active"`
}
