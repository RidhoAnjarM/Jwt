package main

import (
	"main/controllers"
	"main/db"
	"main/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	
	db.ConnectDatabase()

	gin.SetMode(gin.DebugMode)
	engine := gin.Default()
	engine.Use(middleware.Cors())
	routes(engine)

	host := os.Getenv("HOST") + ":" + os.Getenv("PORT")

	err = engine.Run(host)
	if err != nil {
		panic(err)
	}
}

func routes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/me", middleware.AuthMiddleware(), controllers.Me)
}
