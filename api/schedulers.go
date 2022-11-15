package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gnasnik/titan-operator/core/dao"
	"github.com/gnasnik/titan-operator/errors"
	"net/http"
)

func GetSchedulersHandler(c *gin.Context) {
	schedulers, err := dao.GetSchedulers()
	if err != nil {
		c.JSON(http.StatusOK, respError(errors.ErrNotFound))
		return
	}

	c.JSON(http.StatusOK, respJSON(JsonObject{
		"schedulers": schedulers,
	}))
	return
}
