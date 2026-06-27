package dto

/*
	{
	    "id": 5,
	    "name": "Terminal 1 EV Charging",
	    "type": "ev_charging",
	    "total_capacity": 20,
	    "price_per_hour": 5.50,
	    "created_at": "2026-06-20T10:30:00Z",
	    "updated_at": "2026-06-20T10:30:00Z"
	  }
*/
type CreatedParkingZoneResponse struct {
	Id            uint    `json:"id"`
	Name          string  `json:"name"`
	Type          string  `json:"type"`
	TotalCapacity uint    `json:"total_capacity"`
	PricePerHour  float64 `json:"price_per_hour"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}
type ParkingZoneResponse struct {
	Id             uint    `json:"id"`
	Name           string  `json:"name"`
	Type           string  `json:"type"`
	TotalCapacity  uint    `json:"total_capacity"`
	AvailableSpots uint    `json:"available_spots"`
	PricePerHour   float64 `json:"price_per_hour"`
	CreatedAt      string  `json:"created_at"`
}
