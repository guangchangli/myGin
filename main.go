package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	//1.创建路由
	//默认使用两个中间件 Logger Recovery
	r := gin.Default()
	//r:gin.New()
	//2.绑定路由规则，执行函数
	/**
	context param 方法获取 api 参数
	*/
	r.GET("/gin", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin")
	})
	r.GET("/put", func(context *gin.Context) {
		log.Println("gin start")
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "gin json out",
		})
	})
	r.GET("/apiParam/:name/*action", func(context *gin.Context) {
		param := context.Param("name")
		log.Println("url param is ", param)
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "hello " + context.Param("name") + context.Param("action"),
		})
	})
	r.GET("/urlParam", func(context *gin.Context) {
		//不存在返回默认值
		query := context.DefaultQuery("name", "gin")
		context.String(http.StatusOK, fmt.Sprintf("urlParam is %s", query))
	})
	r.POST("/gin/form", func(context *gin.Context) {
		//表达参数设置默认值
		defaultPostForm := context.DefaultPostForm("type", "alert")
		form := context.PostForm("username")
		postForm := context.PostForm("password")
		hobby := context.PostFormArray("hobby")

		context.String(http.StatusOK, "type is %s ,username is %s ,password is %s,hobby is %s", defaultPostForm, form, postForm, hobby)
	})

	r.POST("/gin/upload", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		log.Println(file.Filename)
		//传到项目根目录
		context.SaveUploadedFile(file, file.Filename)
		context.String(http.StatusOK, fmt.Sprintf("%s upload ", file.Filename))

	})
	//限制表单上传大小 默认32MB
	r.MaxMultipartMemory = 8 << 20
	r.POST("/gin/uploads", func(context *gin.Context) {
		form, err := context.MultipartForm()
		if err != nil {
			context.String(http.StatusBadRequest, fmt.Sprintf("get error %s", err.Error()))
		}
		//获取多个图片
		files := form.File["file"]
		//遍历图片
		for _, file := range files {
			//逐个存
			if err := context.SaveUploadedFile(file, file.Filename); err != nil {
				context.String(http.StatusBadRequest, fmt.Sprintf("upload file error %s", err.Error()))
				return
			}
			context.SaveUploadedFile(file, file.Filename)
			context.String(http.StatusOK, fmt.Sprintf("upload file ok %d files", len(files)))
		}
	})

	//router group
	g1 := r.Group("/g1")
	{
		g1.GET("/login", login)
	}
	g2 := r.Group("/g2")
	{
		g2.POST("/submit",submit)
	}

	//3.监听端口
	r.Run(":8000")
	//r.Run("localhost:8080")
}
func login(c *gin.Context){
	name := c.DefaultQuery("name", "what's your name")
	c.String(http.StatusOK,fmt.Sprintf("hello %s\n",name))
}
func submit(c *gin.Context){
	name:= c.DefaultQuery("name", "jasmine")
	c.String(http.StatusOK,fmt.Sprintf("hello %s\n",name))
}
