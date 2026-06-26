package server

import (
	"fmt"
	"spotsync/internal/config"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally return the error to let each route control the status code.
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}

func Start(db *gorm.DB, env *config.Env) {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.RequestLogger())

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(200, map[string]any{
			"ok":      true,
			"message": "Server running",
		})
	})
	fmt.Println("Server running on http://localhost:" + env.Port)
	if err := e.Start(":" + env.Port); err != nil {
		e.Logger.Error("Failed to start the server!", "Error", err.Error())
	}
}
