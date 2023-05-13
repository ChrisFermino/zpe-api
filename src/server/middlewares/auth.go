package middlewares

import (
    "github.com/gin-gonic/gin"
    "zpeTest/src/services"
)

func Auth(role string) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        const BearerSchema = "Bearer "
        header := ctx.GetHeader("Authorization")
        if header == "" {
            ctx.AbortWithStatus(401)
            return
        }
        token := header[len(BearerSchema):]
        ParseToken, err := services.NewJWTService().ValidateToken(token)
        if !err {
            ctx.AbortWithStatus(401)
            return
        }
        switch role {
        case "Admin":
            if !(ParseToken.Role == "Admin") {
                ctx.AbortWithStatus(401)
            }
        case "Modifier":
            if !(ParseToken.Role == "Modifier" || ParseToken.Role == "Admin") {
                ctx.AbortWithStatus(401)
            }
        case "Watcher":
            if !(ParseToken.Role == "Watcher" || ParseToken.Role == "Modifier" || ParseToken.Role == "Admin") {
                ctx.AbortWithStatus(401)
            }
        default:
            ctx.AbortWithStatus(401)
        }
    }
}
