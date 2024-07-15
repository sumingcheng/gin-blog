package router

import (
	_ "blog/docs"
	"blog/handler"
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.Engine) {
	apiRouter := router.Group("/api")
	// 接口
	apiRouter.POST("/login", handler.Login)
	apiRouter.POST("/token", handler.GetAuthToken)

	apiRouter.POST("/blog/belong", handler.BlogBelong)
	apiRouter.GET("/blog/list/:uid", middleware.Auth(), handler.BlogList)
	apiRouter.GET("/blog/:bid", handler.BlogDetail)
	apiRouter.POST("/blog/update", middleware.Auth(), handler.BlogUpdate)
}
