package model

import "gorm.io/gorm"

type Action struct {
	gorm.Model
	AutoAction     AutoAction `json:"autoAction"`
	AutoActionID   uint       `json:"autoActionID" gorm:"not null"`
	Order          int        `json:"order" gorm:"not null"`
	ActionType     ActionType `json:"actionType" gorm:"not null"`
	NotdeClassName string     `json:"noteClassName" gorm:"not null"`
}
