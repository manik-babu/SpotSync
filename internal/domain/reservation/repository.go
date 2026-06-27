package reservation

import (
	"errors"
	"spotsync/internal/domain/parking"
	parkingdto "spotsync/internal/domain/parking/dto"
	reservationdto "spotsync/internal/domain/reservation/dto"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrParkingZoneNotFound = errors.New("parking zone not found")
	ErrParkingZoneFull     = errors.New("parking zone is full")
)

type repository struct {
	db *gorm.DB
}
type Repository interface {
	GetParkingZoneByID(id uint) (*parkingdto.ParkingZoneResponse, error)
	CreateReservation(reservation *Reservation) error
	CreateReservationWithCapacityCheck(req *reservationdto.CreateReservationRequest, userId uint) (*Reservation, error)
	GetReservationByID(id uint) (*Reservation, error)
	GetAllReservations() ([]Reservation, error)
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetParkingZoneByID(id uint) (*parkingdto.ParkingZoneResponse, error) {
	var parkingZone parkingdto.ParkingZoneResponse
	result := r.db.Model(&parking.ParkingZone{}).Select(`
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

func (r *repository) CreateReservation(reservation *Reservation) error {
	return r.db.Create(reservation).Error
}

func (r *repository) CreateReservationWithCapacityCheck(req *reservationdto.CreateReservationRequest, userId uint) (*Reservation, error) {
	var createdReservation Reservation

	err := r.db.Transaction(func(tx *gorm.DB) error {
		var zone parking.ParkingZone
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&zone, req.ZoneId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrParkingZoneNotFound
			}
			return err
		}

		var activeCount int64
		if err := tx.Model(&Reservation{}).
			Where("zone_id = ? AND status = ?", req.ZoneId, "active").
			Count(&activeCount).Error; err != nil {
			return err
		}

		if activeCount >= int64(zone.TotalCapacity) {
			return ErrParkingZoneFull
		}

		reservation := Reservation{
			UserId:       userId,
			ZoneId:       req.ZoneId,
			LicensePlate: req.LicensePlate,
			Status:       "active",
		}

		if err := tx.Create(&reservation).Error; err != nil {
			return err
		}

		createdReservation = reservation
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &createdReservation, nil
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
