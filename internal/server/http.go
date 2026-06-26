package server

import (
	"fmt"
	"spotsync/internal/config"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func Start(env *config.Env) {
	e := echo.New()
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
