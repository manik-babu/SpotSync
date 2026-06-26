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
}

func NewService(repo Repository, jwtService auth.JWTService) Service {
	return &service{
		repo:       repo,
		jwtService: jwtService,
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

// func (s *service) LoginUser(req *dto.UserLoginRequest) (*dto.LoginResponse, error) {
// 	user, err := s.repo.GetUserByEmail(req.Email)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Verify the password
// 	if !user.CheckPassword(req.Password) {
// 		return nil, ErrorUserNotFound
// 	}

// 	// Generate a JWT token
// 	token, err := s.jwtService.GenerateToken(user.ID, user.Email, user.Name, user.Role)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res := dto.LoginResponse{
// 		Token: token,
// 		User: dto.UserData{
// 			Id:    user.ID,
// 			Name:  user.Name,
// 			Email: user.Email,
// 			Role:  user.Role,
// 		},
// 	}
// 	return &res, nil
// }
