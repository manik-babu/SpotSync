package parking

import (
	"spotsync/internal/auth"
	"spotsync/internal/config"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, env *config.Env) {
	repo := NewRepository(db)
	jwtService := auth.NewJWTService(env.JwtSecret)
	service := NewService(repo, jwtService)
	handler := NewHandler(service)

	api := e.Group("/api/v1/auth")
	api.POST("/register", handler.RegisterUser)
	api.POST("/login", handler.LoginUser)
}
