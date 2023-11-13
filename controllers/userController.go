package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/igrzi/loginScreenAPI/initializers"
	"github.com/igrzi/loginScreenAPI/models"
	"github.com/igrzi/loginScreenAPI/utils"
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

		hashedPassword := utils.MakeSHA256(userData.Password)

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
	// Get email and password from the URL
	URLemail := c.Query("email")
	URLpassword := c.Query("password")

	hashedPassword := utils.MakeSHA256(URLpassword)

	// Query the database
	var user models.User
	err := initializers.DB.Where("email = ? AND password = ?", URLemail, hashedPassword).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(404)
		}
		return
	}

	c.Status(200)
}
