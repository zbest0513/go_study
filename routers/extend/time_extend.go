package extend

import (
	"go_study/global"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"time"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		v4 := uuid.NewV4()
		c.Set("req_uuid",v4.String())
		c.Next()
		status := c.Writer.Status()
		global.GVA_LOG.Println("请求:"+ v4.String() +" url:",c.Request.URL," 返回状态:", status)
		t2 := time.Since(t)
		global.GVA_LOG.Println("请求:"+ v4.String() +" url:",c.Request.URL," 用时:", t2)
	}
}