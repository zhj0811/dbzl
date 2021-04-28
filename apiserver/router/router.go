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

	r.POST("/login", handler.Login)       //登录
	r.POST("/register", handler.Register) //机构注册
	//r.Use(handler.TokenAuthMiddleware())
	dbzl := r.Group("/v1/dbzl")
	{
		//dbzl.POST("/excel/policy", handler.UploadPolicies)
		//dbzl.POST("/excel/service", handler.UploadServices)
		//
		//dbzl.POST("/policy", handler.UploadPolicy)
		//dbzl.GET("/policy/:id", handler.GetPolicyByNumber)
		//dbzl.GET("/policy", handler.GetPolicies)
		////dbzl.PUT("/policy", handler.ModifyPolicy)
		//dbzl.POST("/service", handler.UploadService)
		//dbzl.GET("/service/:id", handler.GetServiceById)
		//dbzl.GET("/service", handler.GetServices)
		dbzl.POST("/invoke/policy", handler.UploadInvokePolicy)
		dbzl.GET("/query/policy/:id", handler.QueryPolicy)
		//dbzl.POST("/invoke/policy/:id", handler.InvokePolicy)
		//dbzl.POST("/invoke/service/:id", handler.InvokeService)
		dbzl.POST("/invoke/service", handler.UploadInvokeService)
		dbzl.GET("/query/service/:id", handler.QueryService)
		dbzl.POST("/invoke/company", handler.UploadCompany)
		dbzl.GET("/query/company/:id", handler.QueryCompany)
		//dbzl.GET("/query/:id", handler.QueryPolicyByNumber) //查询链上信息
		////template := dbzl.Group("/template")
		//dbzl.StaticFS("/template", http.Dir("./template"))
	}
	return r
}
