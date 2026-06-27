package parking

import (
	"spotsync/internal/auth"
	"spotsync/internal/domain/parking/dto"
)

type service struct {
	repo       Repository
	jwtService auth.JWTService
}

type Service interface {
	CreateParkingZone(req *dto.CreateParkingZoneRequest) (*dto.ParkingZoneResponse, error)
	GetAllParkingZones() ([]dto.ParkingZoneResponse, error)
	GetParkingZoneByID(id uint) (*dto.ParkingZoneResponse, error)
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateParkingZone(req *dto.CreateParkingZoneRequest) (*dto.ParkingZoneResponse, error) {
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
	res := dto.ParkingZoneResponse{
		Id:            parkingZone.ID,
		Name:          parkingZone.Name,
		Type:          parkingZone.Type,
		TotalCapacity: parkingZone.TotalCapacity,
		PricePerHour:  parkingZone.PricePerHour,
		CreatedAt:     parkingZone.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:     parkingZone.UpdatedAt.Format("2006-01-02T15:04:05Z"),
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
		res = append(res, dto.ParkingZoneResponse{
			Id:            parkingZone.ID,
			Name:          parkingZone.Name,
			Type:          parkingZone.Type,
			TotalCapacity: parkingZone.TotalCapacity,
			PricePerHour:  parkingZone.PricePerHour,
			CreatedAt:     parkingZone.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:     parkingZone.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		})
	}
	return res, nil
}

func (s *service) GetParkingZoneByID(id uint) (*dto.ParkingZoneResponse, error) {
	parkingZone, err := s.repo.GetParkingZoneByID(id)
	if err != nil {
		return nil, err
	}
	res := dto.ParkingZoneResponse{
		Id:            parkingZone.ID,
		Name:          parkingZone.Name,
		Type:          parkingZone.Type,
		TotalCapacity: parkingZone.TotalCapacity,
		PricePerHour:  parkingZone.PricePerHour,
		CreatedAt:     parkingZone.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:     parkingZone.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
	return &res, nil
}
