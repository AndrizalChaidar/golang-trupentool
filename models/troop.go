package models

import "gorm.io/gorm"

type Troop struct {
	gorm.Model
	Name        string
	Tribe       string
	Type        string
	AttackPower int
	CommanderID uint
	Commander   Commander `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (troop *Troop) BeforeSave(tx *gorm.DB) (err error) {
	if troop.Type == "Infantry" {
		switch troop.Tribe {
		case "Gallia":
			troop.AttackPower = 65
		case "Teuton":
			troop.AttackPower = 60
		case "Roman":
			troop.AttackPower = 70
		}
	} else if troop.Type == "Cavalry" {
		switch troop.Tribe {
		case "Gallia":
			troop.AttackPower = 140
		case "Teuton":
			troop.AttackPower = 150
		case "Roman":
			troop.AttackPower = 180
		}
	}
	return
}
