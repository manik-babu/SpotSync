package parking

import (
	"spotsync/internal/auth"
	"spotsync/internal/domain/user/dto"
)

type service struct {
	repo       Repository
	jwtService auth.JWTService
}

type Service interface {
	RegisterUser(user *dto.UserCreateRequest) (*dto.UserResponse, error)
	LoginUser(user *dto.UserLoginRequest) (*dto.LoginResponse, error)
}

func NewService(repo Repository, jwtService auth.JWTService) Service {
	return &service{
		repo:       repo,
		jwtService: jwtService,
	}
}

func (s *service) RegisterUser(req *dto.UserCreateRequest) (*dto.UserResponse, error) {
	user := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}
	// Hash the password before saving to the database
	user.HashPassword()

	err := s.repo.RegisterUser(&user)
	if err != nil {
		return nil, err
	}

	res := dto.UserResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return &res, nil
}
func (s *service) LoginUser(req *dto.UserLoginRequest) (*dto.LoginResponse, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	// Verify the password
	if !user.CheckPassword(req.Password) {
		return nil, ErrorUserNotFound
	}

	// Generate a JWT token
	token, err := s.jwtService.GenerateToken(user.ID, user.Email, user.Name, user.Role)
	if err != nil {
		return nil, err
	}

	res := dto.LoginResponse{
		Token: token,
		User: dto.UserData{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
	}
	return &res, nil
}
