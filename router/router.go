package router

import (
	"collect/controller/user"
	"collect/middleware"

	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init() {
	router := gin.Default()
	// CrossDomain跨域处理，options请求处理
	router.Use(middleware.CrossDomain())
	v1 := router.Group("/v1")
	{
		v1.POST("/userSave", user.SaveUser)
		v1.GET("/userGet", user.Get)
		v1.GET("/search", user.Search)
	}
	router.Run(":8000")
}
