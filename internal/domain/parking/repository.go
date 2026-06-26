package parking

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}
type Repository interface {
	CreateParkingZone(parkingZone *ParkingZone) error
	GetAllParkingZones() ([]ParkingZone, error)
	GetParkingZoneByID(id uint) (*ParkingZone, error)
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateParkingZone(parkingZone *ParkingZone) error {
	result := r.db.Create(parkingZone)
	return result.Error
}
func (r *repository) GetAllParkingZones() ([]ParkingZone, error) {
	var parkingZones []ParkingZone
	result := r.db.Find(&parkingZones)
	if result.Error != nil {
		return nil, result.Error
	}
	return parkingZones, nil
}

func (r *repository) GetParkingZoneByID(id uint) (*ParkingZone, error) {
	var parkingZone ParkingZone
	result := r.db.First(&parkingZone, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &parkingZone, nil
}
