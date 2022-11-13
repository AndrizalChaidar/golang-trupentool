package views

import (
	"golang-trupentool/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type resCommander struct {
	models.Commander
	No          int
	TotalTroops int
	Selected    bool
}

func Commanders(c *gin.Context, commanders *[]models.Commander) {
	var arrCommander []resCommander
	for i, commander := range *commanders {
		newResCommander := resCommander{
			Commander:   commander,
			No:          i + 1,
			TotalTroops: len(commander.Troops),
		}
		arrCommander = append(arrCommander, newResCommander)
	}
	c.HTML(http.StatusOK, "commanders.html", gin.H{"commanders": arrCommander})
}
