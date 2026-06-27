package parking

import (
	"spotsync/internal/auth"
	"spotsync/internal/config"
	"spotsync/internal/middlewares"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, env *config.Env) {
	repo := NewRepository(db)
	jwtService := auth.NewJWTService(env.JwtSecret)
	service := NewService(repo)
	handler := NewHandler(service)

	api := e.Group("/api/v1/zones")
	api.POST("", handler.CreateParkingZone, middlewares.AuthMiddleware(jwtService, "admin"))
	api.GET("", handler.GetAllParkingZones)
	api.GET("/:id", handler.GetParkingZoneByID)
}
