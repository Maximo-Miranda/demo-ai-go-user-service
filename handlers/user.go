package handlers

import (
	"net/http"
	"user-service/db"
	"user-service/models"

	"github.com/golang-jwt/jwt"   // Para generar y manejar tokens JWT
	"github.com/labstack/echo/v4" // Framework web Echo
	"golang.org/x/crypto/bcrypt"  // Para el cifrado de contraseñas
)

// Register maneja el registro de nuevos usuarios
func Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Payload de solicitud inválido"})
	}

	// Cifra la contraseña antes de guardarla en la base de datos
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error al cifrar la contraseña"})
	}
	user.Password = string(hashedPassword)

	// Genera un token único cifrado para el usuario
	token := generateToken(user.Email)
	user.UserToken = token

	// Guarda el usuario en la base de datos
	if err := db.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error al crear el usuario"})
	}

	// Limpia datos sensibles antes de enviar la respuesta
	user.UserToken = ""
	user.Password = ""

	// Retorna el usuario creado y el token
	return c.JSON(http.StatusCreated, echo.Map{
		"user":  user,
		"token": token,
	})
}

// Login maneja la autenticación de usuarios
func Login(c echo.Context) error {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&loginData); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Payload de solicitud inválido"})
	}

	// Busca el usuario por email
	var user models.User
	if err := db.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Credenciales inválidas"})
	}

	// Verifica la contraseña
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Credenciales inválidas"})
	}

	// Retorna el token del usuario
	return c.JSON(http.StatusOK, echo.Map{
		"user_token": user.UserToken,
	})
}

// ListUsers obtiene la lista de todos los usuarios
func ListUsers(c echo.Context) error {
	var users []models.User
	db.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

// ValidateToken verifica la validez del token de un usuario
func ValidateToken(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error al validar el token"})
	}

	// Limpia datos sensibles antes de enviar la respuesta
	user.UserToken = ""
	user.Password = ""

	return c.JSON(http.StatusOK, echo.Map{
		"user": user,
	})
}

// generateToken crea un token JWT para el usuario
func generateToken(email string) string {
	//TODO: Implementar la obtencion de la clave secreta desde una variable de entorno
	secret := "mi_clave_secreta"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})
	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}
