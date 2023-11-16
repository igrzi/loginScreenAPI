package main

import (
	"github.com/gin-contrib/cors"
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

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:5500"}
	r.Use(cors.New(config))

	r.POST("/submitUser", controllers.UserCreate)
	r.GET("/checkUser", controllers.UserCheck)

	r.Run() // listen and serve on 0.0.0.0:8080
}
