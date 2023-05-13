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
            user.POST("", middlewares.Auth("Modifier"), controllers.SaveUser)
        }
    }

    err := server.Run(":8080")
    if err != nil {
        log.Panic("Error when trying to run server on port :8080")
        return
    }
}
