package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"zpeTest/src/controllers"
	"zpeTest/src/server/middlewares"
)

func HandleRequests() {
	server := gin.Default()

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

	err := server.Run(":8080")
	if err != nil {
		log.Panic("Error when trying to run server on port :8080")
		return
	}
}
