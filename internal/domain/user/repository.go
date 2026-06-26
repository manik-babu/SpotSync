package user

import (
	"errors"

	"gorm.io/gorm"
)

var ErrorUserAlreadyExists = errors.New("User with this email already exists")
var ErrorUserNotFound = errors.New("Email or password is incorrect")

type repository struct {
	db *gorm.DB
}
type Repository interface {
	RegisterUser(user *User) error
	GetUserByEmail(email string) (*User, error)
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) RegisterUser(user *User) error {
	result := r.db.Create(user)
	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return ErrorUserAlreadyExists
		}

		return result.Error
	}

	return nil
}
func (r *repository) GetUserByEmail(email string) (*User, error) {
	var user User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrorUserNotFound // User not found
		}
		return nil, result.Error // Other database error
	}
	return &user, nil
}

// Add repository methods here, e.g., CreateUser, GetUserByEmail, etc.kk
