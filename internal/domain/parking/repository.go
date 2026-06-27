package parking

import (
	"spotsync/internal/domain/parking/dto"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}
type Repository interface {
	CreateParkingZone(parkingZone *ParkingZone) error
	GetAllParkingZones() ([]dto.ParkingZoneResponse, error)
	GetParkingZoneByID(id uint) (*dto.ParkingZoneResponse, error)
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
func (r *repository) GetAllParkingZones() ([]dto.ParkingZoneResponse, error) {
	var parkingZones []dto.ParkingZoneResponse
	result := r.db.Model(&ParkingZone{}).Select(`
		parking_zones.id,
		parking_zones.name,
		parking_zones.type,
		parking_zones.total_capacity,
		parking_zones.price_per_hour,
		parking_zones.created_at,
		parking_zones.total_capacity - (
			SELECT COUNT(*)
			FROM reservations
			WHERE reservations.zone_id = parking_zones.id
			AND reservations.status = 'active'
		) AS available_spots
	`).Scan(&parkingZones)
	if result.Error != nil {
		return nil, result.Error
	}
	return parkingZones, nil
}

func (r *repository) GetParkingZoneByID(id uint) (*dto.ParkingZoneResponse, error) {
	var parkingZone dto.ParkingZoneResponse
	result := r.db.Model(&ParkingZone{}).Select(`
		parking_zones.id,
		parking_zones.name,
		parking_zones.type,
		parking_zones.total_capacity,
		parking_zones.price_per_hour,
		parking_zones.created_at,
		parking_zones.total_capacity - (
			SELECT COUNT(*)
			FROM reservations
			WHERE reservations.zone_id = parking_zones.id
			AND reservations.status = 'active'
		) AS available_spots
	`).Where("parking_zones.id = ?", id).Scan(&parkingZone)

	if result.Error != nil {
		return nil, result.Error
	}
	return &parkingZone, nil
}
