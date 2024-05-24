package route

import (
	"mochimochi/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetExternalWebService(c *gin.Context, db *gorm.DB) {
	externalWebService := model.ExternalWebService{}
	id := c.Param("id")
	db.First(&externalWebService, id)
	c.JSON(http.StatusOK, externalWebService)
}

func CreateExternalWebService(c *gin.Context, db *gorm.DB) {
	externalWebService := model.ExternalWebService{}
	c.BindJSON(&externalWebService)
	db.Create(&externalWebService)
	c.JSON(http.StatusOK, externalWebService)
}

func GetExternalWebServices(c *gin.Context, db *gorm.DB) {
	externalWebServices := []model.ExternalWebService{}
	db.Find(&externalWebServices)
	c.JSON(http.StatusOK, externalWebServices)
}
