package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/kallepan/go-jwt/common/controllers"
	"gitlab.com/kallepan/go-jwt/common/database"
	"gitlab.com/kallepan/go-jwt/common/middlewares"
	"gitlab.com/kallepan/go-jwt/env"
	"gitlab.com/kallepan/go-jwt/jwt"
)

func main() {
	connectionString := env.GetConnectionString()
	err := database.Connect(connectionString)

	if err != nil {
		println("Failed to connect to database!")
		panic(err)
	}

	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()

	router.NoRoute(middlewares.NoRouteHandler)
	router.Use(middlewares.ErrorHandler)
	router.Use(middlewares.CORSMiddleware())
	router.SetTrustedProxies(strings.Split(env.GetValueFromEnv("TRUSTED_PROXIES", ","), ","))

	auth := router.Group("/api")
	{
		auth.POST("/token", jwt.GenerateJWTTokenController)
		auth.POST("/register", jwt.RegisterUser)
	}

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", controllers.Ping)
	}

	secured := router.Group("/api/v1")
	secured.Use(middlewares.AuthMiddleware())
	{
		// Add your routes here
	}

	return router
}
