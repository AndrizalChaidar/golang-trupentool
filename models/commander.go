package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Commander struct {
	gorm.Model
	Name          string  `json:"name"`
	Nation        string  `json:"nation"`
	Age           int     `json:"age"`
	MilitaryForce int     `json:"militaryForce"`
	Troops        []Troop `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

func (commander *Commander) GetTitle() string {
	if commander.MilitaryForce >= 500 {
		return fmt.Sprintf("General	 %s", commander.Name)
	}
	if commander.MilitaryForce >= 100 {
		return fmt.Sprintf("Major %s", commander.Name)
	}
	return fmt.Sprintf("Sergeant %s", commander.Name)
}
