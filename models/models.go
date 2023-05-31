package models

import "gorm.io/gorm"

type Move struct {
	gorm.Model
	PlayerID         string `json:"PlayerID" gorm:"text;not null;default:null"`
	LocationNextTile string `json:"LocationNextTile" gorm:"text;not null;default:null"`
}
