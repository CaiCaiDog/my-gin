package Middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckLogin (c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/")
}