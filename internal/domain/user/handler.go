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
	LoginUser(c *echo.Context) error
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
			Errors:  err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.Error{
			Success: false,
			Message: "Validation failed",
			Errors:  err.Error(),
		})
	}

	res, err := h.service.RegisterUser(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpResponse.Error{
			Success: false,
			Message: "Failed to register user",
			Errors:  err.Error(),
		})
	}

	return c.JSON(201, httpResponse.Success{
		Success: true,
		Message: "User registered successfully",
		Data:    res,
	})
}
func (h *handler) LoginUser(c *echo.Context) error {
	var req dto.UserLoginRequest
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
			Message: "Validation failed",
			Errors:  err.Error(),
		})
	}

	res, err := h.service.LoginUser(&req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, httpResponse.Error{
			Success: false,
			Message: "Login failed",
			Errors:  err.Error(),
		})
	}

	return c.JSON(200, httpResponse.Success{
		Success: true,
		Message: "Login successful",
		Data:    res,
	})
}
