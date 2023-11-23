package helper

import (
	"market-api/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(ctx *gin.Context, data any)  {
	ctx.JSON(http.StatusOK, web.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   data,
	})
}

func ResponseSuccessLogin(ctx *gin.Context, data any)  {
	ctx.JSON(http.StatusOK, web.ResponseLogin{
		Code:   http.StatusOK,
		Status: "success",
		Token:   data,
	})
}

func ResponseCategoryMessage(ctx *gin.Context, data any)  {
	ctx.JSON(http.StatusOK, web.ResponseCategoryMessage{
		Code:   http.StatusOK,
		Status: "success",
		Message:   data,
	})
}

func ResponseSuccessTopup(ctx *gin.Context, data any)  {
	ctx.JSON(http.StatusOK, web.ResponseTopup{
		Code:   http.StatusOK,
		Status: "success",
		Message:   data,
	})
}

func ResponseBadRequest(ctx *gin.Context, data any)  {
	ctx.JSON(http.StatusBadRequest, web.ResponseError{
		Code:   http.StatusBadRequest,
		Status: "bad request",
		Message: data,
	})
}

func ResponseNotFound(ctx *gin.Context, data any)  {
	ctx.JSON(http.StatusNotFound, web.ResponseError{
		Code:   http.StatusNotFound,
		Status: "not found",
		Message: data,
	})
}