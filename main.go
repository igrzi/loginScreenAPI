package main

import (
	"github.com/gin-gonic/gin"
	"github.com/igrzi/loginScreenAPI/controllers"
	"github.com/igrzi/loginScreenAPI/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/submitUser", controllers.UserCreate)

	r.Run() // listen and serve on 0.0.0.0:8080
}
