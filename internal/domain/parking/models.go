package parking

import (
	"gorm.io/gorm"
)

type ParkingZone struct {
	gorm.Model
	Name          string  `json:"name" gorm:"not null"`
	Type          string  `json:"type" gorm:"not null,oneof=general ev_charging covered"`
	TotalCapacity uint    `json:"total_capacity" gorm:"not null"`
	PricePerHour  float64 `json:"price_per_hour" gorm:"not null"`
}
