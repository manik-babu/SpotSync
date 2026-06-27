package dto

type CreateReservationRequest struct {
	ZoneId       uint   `json:"zone_id" validate:"required"`
	LicensePlate string `json:"license_plate" validate:"required"`
}

type ReservationResponse struct {
	Id           uint   `json:"id"`
	UserId       uint   `json:"user_id"`
	ZoneId       uint   `json:"zone_id"`
	LicensePlate string `json:"license_plate"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
