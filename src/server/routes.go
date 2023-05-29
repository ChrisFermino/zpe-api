package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"zpeTest/docs"
	"zpeTest/src/controllers"
	"zpeTest/src/server/middlewares"
)

func HandleRequests() {
	server := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"

	api := server.Group("/api")
	{
		api.POST("login", controllers.Login)

		user := api.Group("user")
		{
			user.GET("", middlewares.Auth("Watcher"), controllers.FindUserByFilters)
			user.POST("", middlewares.Auth("Modifier"), controllers.SaveUser)
			user.PATCH("", middlewares.Auth("Modifier"), controllers.EditUser)
			user.PATCH("/editRole", middlewares.Auth("Admin"), controllers.EditUserWithRole)
			user.DELETE("", middlewares.Auth("Modifier"), controllers.DeleteUser)
		}
	}

	// Swagger documentation endpoint
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := server.Run(":8080")
	if err != nil {
		log.Panic("Error when trying to run server on port :8080")
		return
	}
}
