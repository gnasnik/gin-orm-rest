package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gnasnik/gin-orm-rest/core/dao"
	"github.com/gnasnik/gin-orm-rest/errors"
	"net/http"
)

func GetSchedulersHandler(c *gin.Context) {
	schedulers, err := dao.GetSchedulers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, respError(errors.ErrNotFound))
		return
	}

	c.JSON(http.StatusOK, respJSON(JsonObject{
		"schedulers": schedulers,
	}))
	return
}
