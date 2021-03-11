package controller

import (
	"github.com/gin-gonic/gin"
	"go_study/stress/model"
	"net/http"
	"strconv"
	"strings"
)

func urlAccessTestHandle (c *gin.Context) {
	url := c.Query("url")
	if !strings.HasPrefix(url,"http://") && !strings.HasPrefix(url,"https://") {
		c.JSON(http.StatusInternalServerError,gin.H{"msg":"url must begin with http:// or https://"})
		return
	}
	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"msg":"count must be number"})
		return
	}
	runner := model.Runner {Target: url, ThreadCount: count}
	c.JSON(http.StatusOK,runner.Run())
	return
}