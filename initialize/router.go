package initialize

import (
	"net/http"
	"noobgo/global"
	"noobgo/router"

	"github.com/gin-gonic/gin"
)

// 初始化总路由

func Routers() *gin.Engine {
	var r = gin.Default()

	r.LoadHTMLGlob("./dist/*.html")      // 添加入口index.html
	r.Static("/static", "./dist/static") // 添加资源路径
	global.NOOBGO_LOG.Debug("register swagger handler")

	// 方便统一添加路由组前缀 多服务器上线使用
	apiGroup := r.Group("/api")

	router.InitUserRouter(apiGroup) // 注册用户路由
	router.InitLotteryRouter(apiGroup)

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	global.NOOBGO_LOG.Info("router register success")

	return r
}
