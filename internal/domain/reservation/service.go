package reservation

import (
	"fmt"
	"spotsync/internal/domain/reservation/dto"
)

type service struct {
	repo Repository
}

type Service interface {
	CreateReservation(req *dto.CreateReservationRequest, userId uint) (*dto.ReservationResponse, error)
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
func (s *service) CreateReservation(req *dto.CreateReservationRequest, userId uint) (*dto.ReservationResponse, error) {
	// Check if the parking zone exists
	parkingZone, err := s.repo.GetParkingZoneByID(req.ZoneId)
	if err != nil {
		return nil, err
	}
	if parkingZone == nil {
		return nil, fmt.Errorf("parking zone with ID %d not found", req.ZoneId)
	}

	reservation := Reservation{
		UserId:       userId,
		ZoneId:       req.ZoneId,
		LicensePlate: req.LicensePlate,
		Status:       "active",
	}

	err = s.repo.CreateReservation(&reservation)
	if err != nil {
		return nil, err
	}

	res := dto.ReservationResponse{
		Id:           reservation.ID,
		UserId:       reservation.UserId,
		ZoneId:       reservation.ZoneId,
		LicensePlate: reservation.LicensePlate,
		Status:       reservation.Status,
		CreatedAt:    reservation.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:    reservation.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}

	return &res, nil
}
