package user

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}
type Repository interface {
	RegisterUser(user *User) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) RegisterUser(user *User) error {
	result := r.db.Create(user).Error
	return result
}

// Add repository methods here, e.g., CreateUser, GetUserByEmail, etc.kk
