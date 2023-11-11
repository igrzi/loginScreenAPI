package controllers

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/gin-gonic/gin"
	"github.com/igrzi/loginScreenAPI/initializers"
	"github.com/igrzi/loginScreenAPI/models"
)

func UserCreate(c *gin.Context) {
	// Get data of user from request body
	var userData struct {
		Email    string
		Password string
	}
	c.Bind(&userData)

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
}

func UserCheck(c *gin.Context) {

}
