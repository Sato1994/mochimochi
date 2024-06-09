package model

import "gorm.io/gorm"

type ExternalWebService struct {
	gorm.Model
	Name string `json:"name" gorm:"not null"`
	Url  string `json:"url" gorm:"not null,uniqueIndex"`
}
