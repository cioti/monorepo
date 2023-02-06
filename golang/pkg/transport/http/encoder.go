package http

import "github.com/gin-gonic/gin"

type ResponseEncoder = func(ctx *gin.Context, statusCode int, response interface{})

func EncodeJSONResponse(ctx *gin.Context, statusCode int, response interface{}) {
	ctx.JSON(statusCode, response)
}
