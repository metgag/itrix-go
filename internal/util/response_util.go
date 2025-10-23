package util

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/metgag/itrix-challenge/internal/model"
)

func CtxResponse(ctx *gin.Context, statusCode int, response any) {
	ctx.JSON(statusCode, response)
}

func CtxErrResponse(
	ctx *gin.Context, statusCode int, errHead string, err error,
) {
	log.Printf("%s, error: %v\n", errHead, err)
	ctx.JSON(statusCode, model.ErrorResponse{
		Success:    false,
		Error:      errHead,
		StatusCode: statusCode,
	})
}

// func MwareCtxErrResponse(
// 	ctx *gin.Context, statusCode int, errHead string, err error,
// ) {
// 	log.Printf("%s, error: %v\n", errHead, err)
// 	ctx.AbortWithStatusJSON(statusCode, model.ErrorResponse{
// 		Success:    false,
// 		Error:      errHead,
// 		StatusCode: statusCode,
// 	})
// }
