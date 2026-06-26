package parking

import (
	"net/http"
	"spotsync/internal/domain/parking/dto"
	"spotsync/internal/httpResponse"
	"strconv"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service Service
}

type Handler interface {
	CreateParkingZone(c *echo.Context) error
	GetAllParkingZones(c *echo.Context) error
	GetParkingZoneByID(c *echo.Context) error
}

func NewHandler(service Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreateParkingZone(c *echo.Context) error {
	var req dto.CreateParkingZoneRequest
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

	res, err := h.service.CreateParkingZone(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpResponse.Error{
			Success: false,
			Message: "Failed to create parking zone",
			Errors:  err.Error(),
		})
	}

	return c.JSON(201, httpResponse.Success{
		Success: true,
		Message: "Parking zone created successfully",
		Data:    res,
	})
}

func (h *handler) GetAllParkingZones(c *echo.Context) error {
	res, err := h.service.GetAllParkingZones()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpResponse.Error{
			Success: false,
			Message: "Failed to get parking zones",
			Errors:  err.Error(),
		})
	}

	return c.JSON(200, httpResponse.Success{
		Success: true,
		Message: "Parking zones retrieved successfully",
		Data:    res,
	})
}

func (h *handler) GetParkingZoneByID(c *echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.Error{
			Success: false,
			Message: "Invalid parking zone ID",
			Errors:  err.Error(),
		})
	}

	res, err := h.service.GetParkingZoneByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpResponse.Error{
			Success: false,
			Message: "Failed to get parking zone",
			Errors:  err.Error(),
		})
	}

	return c.JSON(200, httpResponse.Success{
		Success: true,
		Message: "Parking zone retrieved successfully",
		Data:    res,
	})
}
