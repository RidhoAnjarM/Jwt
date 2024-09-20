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

	//user
	r.POST("/users/create", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.CreateUser)
	r.GET("/users/get", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.GetUsers)
	r.GET("/users/get/:id", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.GetUser)
	r.PUT("/users/update/:id", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.UpdateUser)
	r.DELETE("/users/delete/:id", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.DeleteUser)

	//role
	r.POST("/role/create", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.CreateRole)
	r.GET("/role/get", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.GetRoles)
	r.GET("/role/get/:id", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.GetRole)
	r.PUT("/role/update/:id", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.UpdateRole)
	r.DELETE("/role/delete/:id", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.DeleteRole)

	//ac
	r.POST("/ac/create", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.CreateAC)
	r.GET("/ac/get", middleware.AuthMiddleware(), controllers.GetACs)
	r.GET("/ac/get/:id", middleware.AuthMiddleware(), controllers.GetAC)
	r.PUT("/ac/update/:id", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.UpdateAC)
	r.DELETE("/ac/delete/:id", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.DeleteAC)

	//service
	r.POST("/service/create", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2), controllers.CreateService)
	r.GET("/service/get", middleware.AuthMiddleware(), controllers.GetServices)
	r.GET("/service/get/:id", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2), controllers.GetService)
	r.PUT("/service/update/:id", middleware.AuthMiddleware(), middleware.RoleMiddleware(1, 2), controllers.UpdateService)
	r.DELETE("/service/delete/:id", middleware.AuthMiddleware(), middleware.RoleMiddleware(1), controllers.DeleteService)
}
