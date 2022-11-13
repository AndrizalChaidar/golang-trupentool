package controllers

import (
	"golang-trupentool/initializers"
	"golang-trupentool/models"
	"golang-trupentool/views"

	"github.com/gin-gonic/gin"
)

func GetCommanders(c *gin.Context) {
	commanders := RetrieveAllCommanders()
	views.Commanders(c, commanders)
}

func RetrieveAllCommanders() *[]models.Commander {
	var commanders []models.Commander
	initializers.DB.Find(&commanders)
	return &commanders
}
