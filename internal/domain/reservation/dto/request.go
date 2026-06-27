package dto

type CreateReservationRequest struct {
	ZoneId       uint   `json:"zone_id" validate:"required"`
	LicensePlate string `json:"license_plate" validate:"required"`
}
