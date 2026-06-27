package reservation

import (
	"spotsync/internal/domain/reservation/dto"
	"time"
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
	reservation, err := s.repo.CreateReservationWithCapacityCheck(req, userId)
	if err != nil {
		return nil, err
	}

	res := dto.ReservationResponse{
		Id:           reservation.ID,
		UserId:       reservation.UserId,
		ZoneId:       reservation.ZoneId,
		LicensePlate: reservation.LicensePlate,
		Status:       reservation.Status,
		CreatedAt:    reservation.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt:    reservation.UpdatedAt.UTC().Format(time.RFC3339),
	}

	return &res, nil
}
