package reservation

import (
	"errors"
	"net/http"
	"spotsync/internal/domain/reservation/dto"
	"spotsync/internal/httpResponse"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service Service
}

type Handler interface {
	CreateReservation(c *echo.Context) error
	GetMyReservations(c *echo.Context) error
}

func NewHandler(service Service) Handler {
	return &handler{
		service: service,
	}
}
func (h *handler) CreateReservation(c *echo.Context) error {
	var req dto.CreateReservationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.Error{
			Success: false,
			Message: "Invalid request payload",
			Errors:  err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.Error{
			Success: false,
			Message: "Invalid input type or required field is empty",
			Errors:  err.Error(),
		})
	}

	userId, ok := c.Get("userID").(uint)
	if !ok {
		return c.JSON(http.StatusInternalServerError, httpResponse.Error{
			Success: false,
			Message: "Failed to retrieve user ID from context",
		})
	}
	res, err := h.service.CreateReservation(&req, userId)
	if err != nil {
		statusCode := http.StatusInternalServerError
		message := "Failed to create reservation"

		switch {
		case errors.Is(err, ErrParkingZoneNotFound):
			statusCode = http.StatusNotFound
			message = "Parking zone not found"
		case errors.Is(err, ErrParkingZoneFull):
			statusCode = http.StatusConflict
			message = "Parking zone is full"
		}

		return c.JSON(statusCode, httpResponse.Error{
			Success: false,
			Message: message,
			Errors:  err.Error(),
		})
	}

	return c.JSON(201, httpResponse.Success{
		Success: true,
		Message: "Reservation confirmed successfully",
		Data:    res,
	})
}
func (h *handler) GetMyReservations(c *echo.Context) error {
	userId, ok := c.Get("userID").(uint)
	if !ok {
		return c.JSON(http.StatusInternalServerError, httpResponse.Error{
			Success: false,
			Message: "Failed to retrieve user ID from context",
		})
	}

	reservations, err := h.service.GetMyReservations(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpResponse.Error{
			Success: false,
			Message: "Failed to retrieve reservations",
			Errors:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, httpResponse.Success{
		Success: true,
		Message: "My reservations retrieved successfully",
		Data:    reservations,
	})
}
