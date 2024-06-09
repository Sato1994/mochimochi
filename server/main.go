// main.go
package main

import (
	"mochimochi/model"
	"mochimochi/route"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=db port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	println("Database connect Success!")
	migrateSchema(db)

	r := gin.Default()
	route.InitRouter(r, db)

	r.Run(":8080")
}

func migrateSchema(db *gorm.DB) {

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Action{})
	db.AutoMigrate(&model.AutoAction{})
	db.AutoMigrate(&model.ExternalWebService{})

}
