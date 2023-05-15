package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"zpeTest/src/models"
	"zpeTest/src/request_handler"
	"zpeTest/src/services"
)

// Control layer handling http requests

func FindUserByFilters(ctx *gin.Context) {
	queryValues := ctx.Request.URL.Query()

	filters := make(map[string]string)
	for key, values := range queryValues {
		filters[key] = strings.Join(values, ",")
	}
	data, err := services.FindUserByFilters(filters)
	if err != nil {
		request_handler.RequestHandler(ctx, nil, http.StatusBadRequest, err.Error())
		return
	}
	request_handler.RequestHandler(ctx, data, http.StatusOK, "")
}

func SaveUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		request_handler.RequestHandler(ctx, nil, http.StatusBadRequest, err.Error())
		return
	}
	data, err := services.CreateUser(user)
	if len(err) > 0 {
		request_handler.RequestHandler(ctx, nil, http.StatusBadRequest, err)
		return
	}
	request_handler.RequestHandler(ctx, data, http.StatusCreated, "")
}

func EditUser(ctx *gin.Context) {
	var user models.UserUpdateDTO
	id := ctx.Query("id")
	if len(id) == 0 {
		request_handler.RequestHandler(ctx, nil, http.StatusBadRequest, "the id field was not informed")
		return
	}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		request_handler.RequestHandler(ctx, nil, http.StatusBadRequest, err.Error())
		return
	}
	data, err := services.EditUserDTO(user, id)
	if len(err) > 0 {
		request_handler.RequestHandler(ctx, nil, http.StatusBadRequest, err)
		return
	}
	request_handler.RequestHandler(ctx, data, http.StatusOK, "")
}

func EditUserWithRole(ctx *gin.Context) {
	var user models.UserUpdateRoleDTO
	id := ctx.Query("id")
	if len(id) == 0 {
		request_handler.RequestHandler(ctx, nil, http.StatusBadRequest, "the id field was not informed")
		return
	}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		request_handler.RequestHandler(ctx, nil, http.StatusBadRequest, err.Error())
		return
	}
	data, err := services.EditUserWithRoleDTO(user, id)
	if len(err) > 0 {
		request_handler.RequestHandler(ctx, nil, http.StatusBadRequest, err)
		return
	}
	request_handler.RequestHandler(ctx, data, http.StatusOK, "")
}

func DeleteUser(ctx *gin.Context) {
	id := ctx.Query("id")
	data, err := services.DeleteUser(id)
	if len(err) > 0 {
		request_handler.RequestHandler(ctx, nil, http.StatusBadRequest, err)
		return
	}
	request_handler.RequestHandler(ctx, data, http.StatusOK, "")
}
