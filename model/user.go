package model

type User struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

type ExternalWebService struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"not null"`
	Url  string `json:"url" gorm:"not null"`
}
