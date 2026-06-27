package reservation

import "gorm.io/gorm"

type Reservation struct {
	gorm.Model
	UserId       uint   `json:"user_id" gorm:"not null"`
	ZoneId       uint   `json:"zone_id" gorm:"not null"`
	LicensePlate string `json:"license_plate" gorm:"not null"`
	Status       string `json:"status" gorm:"type:varchar(10);check:status IN ('active','completed', 'cancelled');not null;default:'active'"`
}
