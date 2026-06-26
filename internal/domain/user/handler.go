package user

import (
	"net/http"
	"spotsync/internal/domain/user/dto"
	"spotsync/internal/httpResponse"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service Service
}

type Handler interface {
	RegisterUser(c *echo.Context) error
}

func NewHandler(service Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) RegisterUser(c *echo.Context) error {
	var req dto.UserCreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.Error{
			Success: false,
			Message: "Invalid request payload",
			Errors:  err,
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.Error{
			Success: false,
			Message: "Validation failed",
			Errors:  err,
		})
	}

	res, err := h.service.RegisterUser(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpResponse.Error{
			Success: false,
			Message: "Failed to register user",
			Errors:  err,
		})
	}

	return c.JSON(201, httpResponse.Success{
		Success: true,
		Message: "User registered successfully",
		Data:    res,
	})
}
