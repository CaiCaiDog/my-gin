package router

// import (
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// 	"my-gin/app/controller"
// )

// var (
// 	r *gin.Engine
// )

// // 初始化路由
// func InitRouter(router *gin.Engine) {
// 	r = router
// 	// 视图变量渲染方式
// 	r.Delims("{{{", "}}}")
// 	// 加载视图
// 	r.LoadHTMLGlob("app/view/*")
// 	// 加载路由
// 	Router()
// 	// 加载静态文件
// 	r.StaticFS("/static", http.Dir("static"))
// }

// // 路由配置
// func Router() {
// 	// !认证
// 	r.GET("/", controller.Index.Index)

// 	// 认证
// 	ar := r.Group("/")
// 	ar.Use(Check)
// 	{

// 	}
// }

// // 拦截器
// func Check(c *gin.Context) {
// 	if false {
// 		// redirect to login
// 		c.Request.URL.Path = "/login"
// 		r.HandleContext(c)
// 	}
// }