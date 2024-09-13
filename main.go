package main

import (
	"main/db"
	"main/middlewares"
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
	engine.Use(middlewares.Cors())
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
	r.POST("/register", middlewares.Register)
	r.POST("/login", middlewares.Login)
}
