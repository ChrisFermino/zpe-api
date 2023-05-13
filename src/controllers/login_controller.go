package controllers

import (
    "net/http"
    "zpeTest/src/database"
    "zpeTest/src/models"
    "zpeTest/src/services"
    "zpeTest/src/utils"

    "github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
    var l models.Login
    if err := ctx.ShouldBindJSON(&l); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "Error": "cannot bind JSON: " + err.Error()})
        return
    }

    var u models.User
    dbErr := database.DB.Where("email = ?", l.Email).First(&u).Error
    if dbErr != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"Error": "invalid credentials"})
        return
    }

    if !utils.CheckPasswordHash(l.Password, u.Password) {
        ctx.JSON(http.StatusUnauthorized, gin.H{"Error": "invalid credentials"})
        return
    }
    t, err := services.NewJWTService().GenerateToken(u.ID, u.Name, u.Email, u.Role)
    if err != nil {
        ctx.JSON(500, gin.H{"Error": err.Error()})
        return
    }
    ctx.JSON(200, gin.H{"token": t})
}
