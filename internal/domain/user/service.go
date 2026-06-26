package user

import "spotsync/internal/domain/user/dto"

type service struct {
	repo Repository
}

type Service interface {
	RegisterUser(user *dto.UserCreateRequest) (*dto.UserResponse, error)
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
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
