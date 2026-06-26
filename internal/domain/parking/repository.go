package parking

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}
type Repository interface {
	CreateParkingZone(parkingZone *ParkingZone) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateParkingZone(parkingZone *ParkingZone) error {
	result := r.db.Create(parkingZone)
	return result.Error
}

// func (r *repository) GetUserByEmail(email string) (*User, error) {
// 	var user User
// 	result := r.db.Where("email = ?", email).First(&user)
// 	if result.Error != nil {
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			return nil, ErrorUserNotFound // User not found
// 		}
// 		return nil, result.Error // Other database error
// 	}
// 	return &user, nil
// }
