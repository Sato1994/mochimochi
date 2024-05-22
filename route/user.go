package route

import (
	"mochimochi/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUser(c *gin.Context, db *gorm.DB) {
	user := model.User{}
	id := c.Param("id")
	db.First(&user, id)
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	db.Delete(&model.User{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func GetUsers(c *gin.Context, db *gorm.DB) {
	users := []model.User{}
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}
