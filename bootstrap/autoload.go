package bootstrap

import (
	"github.com/gin-gonic/gin"
	//"my-gin/router"
	"my-gin/config"
	"my-gin/router"
	"my-gin/service"
	//"os"
	"net/http"
	"my-gin/app/Middleware"
)

// 加载基础服务
func InitServer() {	
	// runing mode
	gin.SetMode(config.App("app_mode"))
	// gin实例
	r := gin.Default()
	// 异常捕获
	r.Use(Middleware.Recover)
	// 初始化路由服务
	InitRouter(r)
	// 初始化数据库连接
	con := service.GetIns().InitConnect()
	if !con {
		println("mysql connection error!")
	}
	// run server
	r.Run(config.App("host") + ":" + config.App("port"))
}

// 初始化路由信息
func InitRouter(r *gin.Engine) {
	// 视图变量渲染方式
	r.Delims(config.App("delim_left"), config.App("delim_right"))
	// 加载视图
	r.LoadHTMLGlob(config.App("html_path"))
	// 加载静态文件
	r.StaticFS("/static", http.Dir("static"))
	// 加载路由配置
	router.Web(r)
	router.Api(r)
}