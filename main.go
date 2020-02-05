package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	//1.创建路由
	r:=gin.Default()
	//2.绑定路由规则，执行函数
	r.GET("/gin", func(context *gin.Context) {
		context.String(http.StatusOK,"hello gin")
	})
	//3.监听端口
	r.Run(":8000")
}
