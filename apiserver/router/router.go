package router

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/zhj0811/dbzl/apiserver/handler"
)

// Router 全局路由
var router *gin.Engine
var onceCreateRouter sync.Once

//GetRouter 获取路由
func GetRouter() *gin.Engine {
	onceCreateRouter.Do(func() {
		router = createRouter()
	})

	return router
}

func createRouter() *gin.Engine {
	r := gin.Default()

	dbzl := r.Group("/v1/dbzl")
	{
		dbzl.POST("/excel/policy", handler.UploadPolicies)
		dbzl.POST("/excel/service", handler.UploadServices)

		dbzl.POST("/policy", handler.UploadPolicy)
		//dbzl.PUT("/policy", handler.ModifyPolicy)
		dbzl.POST("/service", handler.UploadService)
		//dbzl.POST("/invoke/policy", handler.UploadPolicy)
		dbzl.POST("/invoke/policy/:id", handler.InvokePolicy)
		dbzl.POST("/invoke/service/:id", handler.InvokeService)
		//dbzl.POST("/invoke/service", handler.UploadService)
	}
	return r
}
