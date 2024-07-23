package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/wasupalonely/reservapp/config"
	"github.com/wasupalonely/reservapp/internal/user"
	"github.com/wasupalonely/reservapp/internal/validation"
)

func generateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}

func LoginHandler(c *gin.Context) {
	var loginData validation.LoginData
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := user.GetByIdentifier(loginData.Identifier)
	if err != nil || !user.CheckPasswordHash(loginData.Password, u.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid identifier or password"})
		return
	}

	token, err := generateToken(u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func RegisterHandler(c *gin.Context) {
	var registrationData validation.RegistrationData
	if err := c.ShouldBindJSON(&registrationData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := user.HashPassword(registrationData.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}

	newUser := user.User{
		Username: registrationData.Username,
		Password: hashedPassword,
		Email:    registrationData.Email,
	}

	if err := user.CreateUser(&newUser); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
