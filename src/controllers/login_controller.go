package controllers

import (
	"net/http"
	"zpeTest/src/database"
	"zpeTest/src/models"
	"zpeTest/src/request_handler"
	"zpeTest/src/services"
	"zpeTest/src/utils"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Login to the application
// @Description Logs in a user and returns an access token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param login body models.Login true "Login credentials"
// @Success 200 {object} request_handler.RequestSuccess "Successfully logged in"
// @Failure 400 {object} request_handler.Error "invalid credentials"
// @Router /login [post]
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
		request_handler.RequestHandler(ctx, nil, http.StatusBadRequest, "invalid credentials")
		return
	}

	if !utils.CheckPasswordHash(l.Password, u.Password) {
		request_handler.RequestHandler(ctx, nil, http.StatusBadRequest, "invalid credentials")
		return
	}
	t, err := services.NewJWTService().GenerateToken(u.ID, u.Name, u.Email, u.Role)
	if err != nil {
		request_handler.RequestHandler(ctx, nil, 500, err.Error())
		return
	}
	request_handler.RequestHandler(ctx, t, 200, "")
}
