package main

import (
	"github.com/gin-gonic/gin"

	"github.com/vermakmanish001/otp-auth-app/config"
	"github.com/vermakmanish001/otp-auth-app/handlers"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	r.Run(":8080")
}
