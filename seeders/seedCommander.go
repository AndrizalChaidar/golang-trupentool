package seeders

import (
	"encoding/json"
	"fmt"
	"golang-trupentool/initializers"
	"golang-trupentool/models"
	"io"
	"log"
	"os"
)

type Commander struct {
	Name          string `json:"name"`
	Nation        string `json:"nation"`
	Age           int    `json:"age"`
	MilitaryForce int    `json:"militaryForce"`
}

func Main() {
	commanderFile, err := os.Open("commanders.json")
	if err != nil {
		log.Fatal("Error opening commanders.json")
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened commanders.json")
	defer commanderFile.Close()
	byteValue, _ := io.ReadAll(commanderFile)
	var commanders []models.Commander
	json.Unmarshal(byteValue, &commanders)
	initializers.DB.CreateInBatches(commanders, len(commanders))
}
