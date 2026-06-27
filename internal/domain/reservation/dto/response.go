package dto

type ReservationResponse struct {
	Id           uint   `json:"id"`
	UserId       uint   `json:"user_id"`
	ZoneId       uint   `json:"zone_id"`
	LicensePlate string `json:"license_plate"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
type MyReservationsResponse struct {
	Id           uint   `json:"id"`
	LicensePlate string `json:"license_plate"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
	Zone         Zone   `json:"zone"`
}
type Zone struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
