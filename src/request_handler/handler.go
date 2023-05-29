package request_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InternalServerError(msg string) Error {
	return Error{Message: "Internal Server Error", Status: 500, Error: msg}
}

func BadRequestError(msg string) Error {
	return Error{Message: "Error, check the data and try again!", Status: 400, Error: msg}
}

func NotFoundError(msg string) Error {
	return Error{Message: "Register not found", Status: 404, Error: msg}
}

func RequestHandler(ctx *gin.Context, data any, status int, msg string) {
	res := RequestError{
		Data: data,
	}
	if status >= 300 {
		switch status {
		case http.StatusInternalServerError:
			res.Error = InternalServerError(msg)
		case http.StatusBadRequest:
			res.Error = BadRequestError(msg)
		case http.StatusNotFound:
			res.Error = NotFoundError(msg)
		default:
			res.Error = BadRequestError(msg)
		}
		ctx.JSON(status, res)
		return
	}
	ctx.JSON(status, RequestSuccess{Data: data})
}
