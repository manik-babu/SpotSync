package parking

import (
	"fmt"
	"spotsync/internal/auth"
	"spotsync/internal/domain/parking/dto"
	"time"
)

type service struct {
	repo       Repository
	jwtService auth.JWTService
}

type Service interface {
	CreateParkingZone(req *dto.CreateParkingZoneRequest) (*dto.CreatedParkingZoneResponse, error)
	GetAllParkingZones() ([]dto.ParkingZoneResponse, error)
	GetParkingZoneByID(id uint) (*dto.ParkingZoneResponse, error)
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func formatUTCZulu(t time.Time) string {
	return t.UTC().Format("2006-01-02T15:04:05Z")
}

func parseAndFormatUTCZulu(value string) (string, error) {
	parsed, err := time.Parse(time.RFC3339Nano, value)
	if err != nil {
		return "", err
	}
	return formatUTCZulu(parsed), nil
}

func (s *service) CreateParkingZone(req *dto.CreateParkingZoneRequest) (*dto.CreatedParkingZoneResponse, error) {
	parkingZone := ParkingZone{
		Name:          req.Name,
		Type:          req.Type,
		TotalCapacity: req.TotalCapacity,
		PricePerHour:  req.PricePerHour,
	}
	err := s.repo.CreateParkingZone(&parkingZone)
	if err != nil {
		return nil, err
	}
	res := dto.CreatedParkingZoneResponse{
		Id:            parkingZone.ID,
		Name:          parkingZone.Name,
		Type:          parkingZone.Type,
		TotalCapacity: parkingZone.TotalCapacity,
		PricePerHour:  parkingZone.PricePerHour,
		CreatedAt:     formatUTCZulu(parkingZone.CreatedAt),
		UpdatedAt:     formatUTCZulu(parkingZone.UpdatedAt),
	}
	return &res, nil

}

func (s *service) GetAllParkingZones() ([]dto.ParkingZoneResponse, error) {
	parkingZones, err := s.repo.GetAllParkingZones()
	if err != nil {
		return nil, err
	}
	var res []dto.ParkingZoneResponse
	for _, parkingZone := range parkingZones {
		createdAt, err := parseAndFormatUTCZulu(parkingZone.CreatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, dto.ParkingZoneResponse{
			Id:             parkingZone.Id,
			Name:           parkingZone.Name,
			Type:           parkingZone.Type,
			TotalCapacity:  parkingZone.TotalCapacity,
			AvailableSpots: parkingZone.AvailableSpots,
			PricePerHour:   parkingZone.PricePerHour,
			CreatedAt:      createdAt,
		})
	}
	return res, nil
}

func (s *service) GetParkingZoneByID(id uint) (*dto.ParkingZoneResponse, error) {
	parkingZone, err := s.repo.GetParkingZoneByID(id)
	if err != nil {
		return nil, err
	}
	fmt.Println(parkingZone) // Debugging line
	createdAt, err := parseAndFormatUTCZulu(parkingZone.CreatedAt)
	if err != nil {
		return nil, err
	}
	res := dto.ParkingZoneResponse{
		Id:             parkingZone.Id,
		Name:           parkingZone.Name,
		Type:           parkingZone.Type,
		TotalCapacity:  parkingZone.TotalCapacity,
		AvailableSpots: parkingZone.AvailableSpots,
		PricePerHour:   parkingZone.PricePerHour,
		CreatedAt:      createdAt,
	}
	return &res, nil
}
