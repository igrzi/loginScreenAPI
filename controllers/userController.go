package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/igrzi/loginScreenAPI/initializers"
	"github.com/igrzi/loginScreenAPI/models"
	"gorm.io/gorm"
)

func UserCreate(c *gin.Context) {

	// Get data of user from request body
	var userData struct {
		Email    string
		Password string
	}
	c.Bind(&userData)

	if !checkIfUserExists(userData.Email) {
		// Encrypt password

		encryptedPassword := sha256.New()
		encryptedPassword.Write([]byte(userData.Password))
		hashedPassword := hex.EncodeToString(encryptedPassword.Sum(nil))

		// Create a user on the database
		user := models.User{Email: userData.Email, Password: hashedPassword}
		result := initializers.DB.Create(&user) // pass pointer of data to Create

		if result.Error != nil {
			c.Status(400)
			return
		}

		// Return a 200 status

		c.Status(200)
	} else {
		c.JSON(409, "User already on database")
	}
}

func checkIfUserExists(email string) bool {
	var user models.User

	// Check if the user with the provided email exists
	result := initializers.DB.Where("email = ?", email).First(&user)
	userExists := !errors.Is(result.Error, gorm.ErrRecordNotFound)

	return userExists
}

func UserCheck(c *gin.Context) {
	// Get email from the URL
	email := c.Param("email")

	var user models.User
	err := initializers.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(404)
		}
		return
	}

	c.Status(200)
}
