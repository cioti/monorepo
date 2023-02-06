package http

import (
	"github.com/cioti/monorepo/pkg/api"
	"github.com/gin-gonic/gin"
)

type RequestDecoder = func(ctx *gin.Context) (interface{}, api.ApiError)

func NopDecoder(ctx *gin.Context) (interface{}, api.ApiError) {
	return nil, nil
}
