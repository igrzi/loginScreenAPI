package main

import (
	"github.com/igrzi/loginScreenAPI/initializers"
	"github.com/igrzi/loginScreenAPI/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
