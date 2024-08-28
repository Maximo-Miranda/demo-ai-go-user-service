package middleware

import (
	"net/http"
	"user-service/db"
	"user-service/models"

	"github.com/labstack/echo/v4"
)

func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token no proporcionado"})
		}

		var user models.User
		if err := db.DB.Where("user_token = ?", token).First(&user).Error; err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token inválido"})
		}

		c.Set("user_id", user.ID)
		return next(c)
	}
}
