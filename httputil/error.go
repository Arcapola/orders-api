// Package httputil contains HTTP related utilities.
package httputil

import "github.com/gin-gonic/gin"

// NewError returns a JSON-compatible HTTP error.
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// HTTPError is a JSON-compatible HTTP error representation.
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
