package api

import (
	"github.com/gin-gonic/gin"
	err "github.com/gnasnik/gin-orm-rest/errors"
	"github.com/pkg/errors"
)

type JsonObject map[string]interface{}

func respJSON(v interface{}) gin.H {
	return gin.H{
		"success": true,
		"data":    v,
	}
}

func respError(e error) gin.H {
	var apiError err.ApiError
	if !errors.As(e, &apiError) {
		apiError = err.ErrUnknown
	}

	return gin.H{
		"success": false,
		"code":    apiError.Code(),
		"message": apiError.Error(),
	}
}
