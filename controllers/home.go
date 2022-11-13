package controllers

import (
	"golang-trupentool/views"

	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	views.Home(c)
}
