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

// FindUserByFilters godoc
// @Summary Find users by filters
// @Description Find users based on query parameters
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id query string false "User ID"
// @Param name query string false "name"
// @Param email query string false "User Email"
// @Param role query string false "User Role"
// @Success 200 {object} request_handler.RequestSuccess
// @Failure 400 {object} request_handler.RequestError
// @Failure 401
// @Router /user [get]
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

// SaveUser godoc
// @Summary Save user
// @Description save a user to the database
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param User body models.User true "User model"
// @Success 200 {object} request_handler.RequestSuccess
// @Failure 400 {object} request_handler.RequestError "Bad RequestError"
// @Failure 401
// @Router /user [post]
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

// EditUser godoc
// @Summary Edit user without Role
// @Description With this endpoint it is possible to change the user in the database but without changing its "Role"
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "user Id"
// @Param User body models.UserUpdateDTO true "User model"
// @Success 200 {object} request_handler.RequestSuccess
// @Failure 400 {object} request_handler.RequestError "Bad RequestError"
// @Failure 401
// @Router /user [patch]
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

// EditUserWithRole godoc
// @Summary Edit user with Role
// @Description With this endpoint it is possible to change the user's "Role" in the database
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "user Id"
// @Param User body models.UserUpdateRoleDTO true "User model"
// @Success 200 {object} request_handler.RequestSuccess
// @Failure 400 {object} request_handler.RequestError "Bad RequestError"
// @Failure 401
// @Router /user/editRole [patch]
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

// DeleteUser godoc
// @Summary Delete user
// @Description Delete a user from the database
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "user Id"
// @Success 200 {object} request_handler.RequestSuccess
// @Failure 400 {object} request_handler.RequestError "Bad RequestError"
// @Failure 401
// @Router /user [delete]
func DeleteUser(ctx *gin.Context) {
	id := ctx.Query("id")
	data, err := services.DeleteUser(id)
	if len(err) > 0 {
		request_handler.RequestHandler(ctx, nil, http.StatusBadRequest, err)
		return
	}
	request_handler.RequestHandler(ctx, data, http.StatusOK, "")
}
