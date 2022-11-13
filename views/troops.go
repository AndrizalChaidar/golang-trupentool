package views

import (
	"fmt"
	"golang-trupentool/models"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type resTroop struct {
	models.Troop
	No       int
	Selected bool
}

func GetTroops(c *gin.Context, troops *[]models.Troop, commanders *[]models.Commander) {
	var arrTroops []resTroop
	var arrCommander []resCommander
	var wg sync.WaitGroup
	qId := c.Query("CommanderId")
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i, troop := range *troops {
			newresTroop := resTroop{
				Troop: troop,
				No:    i + 1,
			}
			arrTroops = append(arrTroops, newresTroop)
		}
	}()
	go func() {
		defer wg.Done()
		for _, commander := range *commanders {
			id := fmt.Sprintf("%d", commander.ID)
			newResCommander := resCommander{
				Commander: commander,
				Selected:  id == qId,
			}
			arrCommander = append(arrCommander, newResCommander)
		}
	}()
	wg.Wait()
	c.HTML(http.StatusOK, "troops.html", gin.H{"troops": arrTroops, "commanders": arrCommander, "commanderId": qId})
}

func GetTroopsTrain(c *gin.Context, commanders *[]models.Commander) {
	c.HTML(http.StatusOK, "addTroops.html", gin.H{"commanders": *commanders})
}

func PostTroopsTrain(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/commanders")
}
