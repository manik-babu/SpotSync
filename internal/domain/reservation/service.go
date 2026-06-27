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
	GetMyReservations(userId uint) ([]dto.MyReservationsResponse, error)
	CancelReservation(id uint, userId uint) error
	GetAllReservations() ([]dto.AllReservationsResponse, error)
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
func (s *service) GetMyReservations(userId uint) ([]dto.MyReservationsResponse, error) {
	reservations, err := s.repo.GetMyReservations(userId)
	if err != nil {
		return nil, err
	}
	responses := make([]dto.MyReservationsResponse, len(reservations))
	for i, reservation := range reservations {
		responses[i] = dto.MyReservationsResponse{
			Id:           reservation.ID,
			LicensePlate: reservation.LicensePlate,
			Status:       reservation.Status,
			Zone: dto.Zone{
				Id:   reservation.Zone.ID,
				Name: reservation.Zone.Name,
				Type: reservation.Zone.Type,
			},
			CreatedAt: reservation.CreatedAt.UTC().Format(time.RFC3339),
		}
	}

	return responses, nil
}
func (s *service) CancelReservation(id uint, userId uint) error {
	err := s.repo.CancelReservation(id, userId)
	if err != nil {
		return err
	}
	return nil
}
func (s *service) GetAllReservations() ([]dto.AllReservationsResponse, error) {
	reservations, err := s.repo.GetAllReservations()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.AllReservationsResponse, len(reservations))
	for i, reservation := range reservations {
		responses[i] = dto.AllReservationsResponse{
			Id:           reservation.ID,
			LicensePlate: reservation.LicensePlate,
			Status:       reservation.Status,
			CreatedAt:    reservation.CreatedAt.UTC().Format(time.RFC3339),
			Zone: dto.Zone{
				Id:   reservation.Zone.ID,
				Name: reservation.Zone.Name,
				Type: reservation.Zone.Type,
			},
			User: dto.User{
				Id:    reservation.User.ID,
				Name:  reservation.User.Name,
				Email: reservation.User.Email,
			},
		}
	}

	return responses, nil
}
