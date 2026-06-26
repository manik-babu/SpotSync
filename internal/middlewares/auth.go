package middlewares

import (
	"slices"
	"spotsync/internal/auth"
	"spotsync/internal/httpResponse"
	"strings"

	"github.com/labstack/echo/v5"
)

func AuthMiddleware(jwtService auth.JWTService, UserRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(401, httpResponse.Error{
					Success: false,
					Message: "Missing Authorization header",
				})
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(401, httpResponse.Error{
					Success: false,
					Message: "Invalid Authorization header format",
				})
			}

			token := parts[1]
			claims, err := jwtService.ValidateToken(token)
			if err != nil {
				return c.JSON(401, httpResponse.Error{
					Success: false,
					Message: "Invalid or expired token",
					Errors:  err.Error(),
				})
			}

			// Check if the user has any of the required roles
			if len(UserRoles) > 0 {
				hasRole := slices.Contains(UserRoles, claims.Role)
				if !hasRole {
					return c.JSON(403, httpResponse.Error{
						Success: false,
						Message: "Forbidden: insufficient permissions",
					})
				}
			}

			c.Set("userID", claims.UserID)
			c.Set("userRole", claims.Role)
			c.Set("userEmail", claims.Email)
			return next(c)
		}
	}
}
