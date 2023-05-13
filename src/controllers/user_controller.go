package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "zpeTest/src/models"
    "zpeTest/src/request_handler"
    "zpeTest/src/services"
)

func SaveUser(ctx *gin.Context) {
    var u models.User
    if err := ctx.ShouldBindJSON(&u); err != nil {
        request_handler.RequestHandler(ctx, nil, http.StatusBadRequest, err.Error())
        return
    }
    data, err := services.CreateUser(u)
    if len(err) > 0 {
        request_handler.RequestHandler(ctx, nil, http.StatusBadRequest, err)
        return
    }
    request_handler.RequestHandler(ctx, data, http.StatusCreated, "")
}
