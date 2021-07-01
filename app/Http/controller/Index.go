package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"my-gin/app/Http/model"
)

// index
type IndexController struct {}

var (
	Index = &IndexController {}
)

// 首页
func (i *IndexController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// welcome
func (i *IndexController) Welcome(c *gin.Context) {
	res := model.User.FindUser()
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": res,
	})
}