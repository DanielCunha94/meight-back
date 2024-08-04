package errors

import (
	"github.com/gin-gonic/gin"
	"log"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if ctx.Errors != nil {
			err := ctx.Errors.Last()

			appError := err.Err.(*AppError)
			log.Printf("Error: Code %d, Message: %s", appError.Code, appError.Message)

			var errorMessage string
			switch appError.Code {
			case ErrBadRequest:
				errorMessage = "Bad request"
			case ErrNotFound:
				errorMessage = "Resource not found"
			case ErrConflict:
				errorMessage = "Conflict error"
			default:
				errorMessage = "Internal server error"
			}
			ctx.AbortWithStatusJSON(appError.Code, gin.H{
				"error": errorMessage,
			})
		}

	}
}
