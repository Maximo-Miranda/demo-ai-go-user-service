package handlers

import (
	"net/http"
	"user-service/db"
	"user-service/models"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Payload de solicitud inválido"})
	}

	// Cifrar la contraseña antes de guardarla en la base de datos
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error al cifrar la contraseña"})
	}
	user.Password = string(hashedPassword)

	// Generar token único cifrado
	token := generateToken(user.Email)
	user.UserToken = token

	// Guardar usuario en la base de datos
	if err := db.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error al crear el usuario"})
	}

	user.UserToken = ""
	user.Password = ""

	// Retornar el token en la respuesta
	return c.JSON(http.StatusCreated, echo.Map{
		"user":  user,
		"token": token,
	})
}

func Login(c echo.Context) error {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&loginData); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Payload de solicitud inválido"})
	}

	var user models.User
	if err := db.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Credenciales inválidas"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Credenciales inválidas"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user_token": user.UserToken,
	})
}

func ListUsers(c echo.Context) error {
	var users []models.User
	db.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func ValidateToken(c echo.Context) error {

	userID := c.Get("user_id").(uint)

	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error al validar el token"})
	}

	user.UserToken = ""
	user.Password = ""

	return c.JSON(http.StatusOK, echo.Map{
		"user": user,
	})
}

func generateToken(email string) string {
	// Implementar lógica para generar un token único cifrado basado en el email del usuario
	// Puedes utilizar librerías como jwt-go o crear tu propia implementación
	// Aquí se muestra un ejemplo básico usando el email y una clave secreta
	secret := "mi_clave_secreta"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})
	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}
