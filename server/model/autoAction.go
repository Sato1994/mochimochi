package model

import "gorm.io/gorm"

type AutoAction struct {
	gorm.Model
	Actions []Action `json:"actions"`
}
