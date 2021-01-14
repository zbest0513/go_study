package main

import (
	"github.com/gin-gonic/gin"
	"go_study/global"
	"go_study/routers"
	"go_study/week7/controller"
	"io"
	"log"
	"os"
)

func main() {
	// 加载多个APP的路由配置
	routers.Include(controller.Routers)
	f, _ := os.Create("gin.log")
	global.GVA_LOG = log.New(f, "[info]", log.Ldate|log.Ltime|log.Lshortfile)

	// 初始化路由
	r := routers.Init()
	defer f.Close()
	gin.DefaultWriter = io.MultiWriter(f)
	if err := r.Run(":8085"); err != nil {
		global.GVA_LOG.Printf("startup service failed, err:%v\n", err)
	}
}