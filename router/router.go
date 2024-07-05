package router

import (
	"gin-rest-api/controller"
	"gin-rest-api/controller/v1"
	_ "gin-rest-api/docs"
	"gin-rest-api/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.GET("/", controller.Index)

	//获取token
	server.POST("/api/v1/auth", v1.GetAuth)
	apiV1 := server.Group("/api/v1")
	apiV1.Use(middleware.JWT())
	{
		//获取标签列表
		apiV1.GET("/tags", v1.GetTags)
		//新建标签
		apiV1.POST("/tags", v1.AddTag)
		//获取指定标签
		apiV1.GET("/tags/:id", v1.GetTag)
		//更新指定标签
		apiV1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiV1.DELETE("/tags/:id", v1.DeleteTag)

	}
}
