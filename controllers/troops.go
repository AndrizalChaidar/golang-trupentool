package controllers

import (
	"fmt"
	"golang-trupentool/helpers"
	"golang-trupentool/initializers"
	"golang-trupentool/models"
	"golang-trupentool/views"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTroops(c *gin.Context) {
	query := c.Query("CommanderId")
	fmt.Println(query)
	var troops *[]models.Troop
	var commanders *[]models.Commander
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		commanders = RetrieveAllCommanders()
	}()
	go func() {
		defer wg.Done()
		if len(query) > 0 {
			troops = RetrieveTroopsByCommanderId(query)
		} else {
			troops = RetrieveAllTroops()
		}

	}()
	wg.Wait()
	views.GetTroops(c, troops, commanders)
}

func GetTroopsTrain(c *gin.Context) {
	var commanders *[]models.Commander
	commanders = RetrieveAllCommanders()
	views.GetTroopsTrain(c, commanders)
}

func PostTroopTrain(c *gin.Context) {
	var commander models.Commander
	commanderId := helpers.StringToUint(c.PostForm("CommanderId"))
	newTroop := &models.Troop{
		Name:        c.PostForm("name"),
		Tribe:       c.PostForm("tribe"),
		CommanderID: commanderId,
		Type:        c.PostForm("type"),
	}
	if err := initializers.DB.First(&commander, commanderId).Error; err != nil {
		log.Fatal(err)
	}
	var addForce int

	initializers.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(newTroop).Error; err != nil {
			return err
		}
		if newTroop.Tribe == commander.Nation {
			addForce = int((float32(newTroop.AttackPower) + float32(commander.MilitaryForce)) * 1.75)
		} else {
			addForce = newTroop.AttackPower + commander.MilitaryForce
		}
		if err := tx.Model(&commander).Updates(models.Commander{MilitaryForce: addForce}).Error; err != nil {
			return err
		}
		return nil
	})
	views.PostTroopsTrain(c)
}

func RetrieveAllTroops() *[]models.Troop {
	var troops []models.Troop
	initializers.DB.Preload("Commander").Find(&troops)
	return &troops
}

func RetrieveTroopsByCommanderId(strCommanderId string) *[]models.Troop {
	var troops []models.Troop
	commanderId := helpers.StringToUint(strCommanderId)
	initializers.DB.Preload("Commander").Where(models.Troop{CommanderID: commanderId}).Find(&troops)
	return &troops
}
