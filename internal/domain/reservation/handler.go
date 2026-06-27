package reservation

import (
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
		return c.JSON(http.StatusInternalServerError, httpResponse.Error{
			Success: false,
			Message: "Failed to create reservation",
			Errors:  err.Error(),
		})
	}

	return c.JSON(201, httpResponse.Success{
		Success: true,
		Message: "Reservation confirmed successfully",
		Data:    res,
	})
}
