package router

import (
	"noobgo/api"

	"github.com/gin-gonic/gin"
)

func InitLotteryRouter(Router *gin.RouterGroup) {
	LotteryRouter := Router.Group("lottery")
	{
		LotteryRouter.POST("", api.LotteryAddAPI)
		LotteryRouter.GET("result", api.LotteryResultsAPI)
	}
}
