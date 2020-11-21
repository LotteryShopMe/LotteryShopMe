package router

import (
	"noobgo/api"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("users")
	{
		UserRouter.GET("", api.UserAddAPI)
		UserRouter.GET(":name", api.GETUserAPI)
		UserRouter.POST(":name/add", api.UserAddAmountAPI)
		UserRouter.GET(":name/add", api.UserAddAmountAPI)
		UserRouter.GET(":name/buyflag", api.BuyFlagAPI)
	}
}
