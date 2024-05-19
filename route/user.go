package route

import (
	"mochimochi/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUser(c *gin.Context, db *gorm.DB) {
	var users model.User
	findUser := db.Order("name").Find(&users)

	c.JSON(http.StatusOK, findUser)

}
