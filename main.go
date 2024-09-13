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
	// db.Connect()
	db.Connect()

	gin.SetMode(gin.DebugMode)
	engine := gin.Default()
	engine.Use(middlewares.Cors())
	routes(engine)

	host := os.Getenv("HOST") + ":" + os.Getenv("PORT")

	err = engine.Run(host)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// Routes
    r.POST("/register", middlewares.Register)
    r.POST("/login", middlewares.Login)

}

func routes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}
