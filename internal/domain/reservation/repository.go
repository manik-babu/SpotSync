package reservation

import (
	"errors"
	"spotsync/internal/domain/parking"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}
type Repository interface {
	GetParkingZoneByID(id uint) (*parking.ParkingZone, error)
	CreateReservation(reservation *Reservation) error
	GetReservationByID(id uint) (*Reservation, error)
	GetAllReservations() ([]Reservation, error)
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetParkingZoneByID(id uint) (*parking.ParkingZone, error) {
	var parkingZone parking.ParkingZone
	result := r.db.First(&parkingZone, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("parking zone not found")
		}
		return nil, result.Error
	}
	return &parkingZone, nil
}

func (r *repository) CreateReservation(reservation *Reservation) error {
	return r.db.Create(reservation).Error
}

func (r *repository) GetReservationByID(id uint) (*Reservation, error) {
	var reservation Reservation
	err := r.db.First(&reservation, id).Error
	if err != nil {
		return nil, err
	}
	return &reservation, nil
}

func (r *repository) GetAllReservations() ([]Reservation, error) {
	var reservations []Reservation
	err := r.db.Find(&reservations).Error
	if err != nil {
		return nil, err
	}
	return reservations, nil
}
