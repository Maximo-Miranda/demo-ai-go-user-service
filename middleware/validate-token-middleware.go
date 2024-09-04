package middleware

import (
	"net/http"            // Proporciona constantes HTTP y tipos para servidores web
	"user-service/db"     // Módulo local para interactuar con la base de datos
	"user-service/models" // Módulo local que contiene las estructuras de datos, como User

	"github.com/labstack/echo/v4" // Framework web Echo para crear la API REST
)

// ValidateToken es un middleware que verifica la validez del token de usuario
func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Obtiene el token del encabezado de autorización
		token := c.Request().Header.Get("Authorization")

		// Verifica si el token está presente
		if token == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token no proporcionado"})
		}

		// Busca un usuario con el token proporcionado en la base de datos
		var user models.User
		if err := db.DB.Where("user_token = ?", token).First(&user).Error; err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token inválido"})
		}

		// Si el token es válido, almacena el ID del usuario en el contexto
		c.Set("user_id", user.ID)

		// Continúa con el siguiente manejador en la cadena
		return next(c)
	}
}
