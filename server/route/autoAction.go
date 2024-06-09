package route

import (
	"mochimochi/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
	"gorm.io/gorm"
)

func GetAutoActions(con *gin.Context, db *gorm.DB) {
	autoActions := []model.AutoAction{}
	db.Find(&autoActions)
	con.JSON(http.StatusOK, autoActions)
}

func GetHtmlFromPath(con *gin.Context) {
	// https://github.com/gocolly/colly
	c := colly.NewCollector()

	url := "https://" + con.Param("url")

	c.OnHTML("title", func(e *colly.HTMLElement) {
	})
	c.OnRequest(func(r *colly.Request) {
		println("Visiting URL:", r.URL.String())
	})

	err := c.Visit(url)

	if err != nil {
		con.JSON(http.StatusBadRequest, "エラー"+err.Error())
		return
	}
	con.JSON(http.StatusOK, "遷移成功")
}

func DeleteHtmlFromPath(con *gin.Context, db *gorm.DB) {
	removeId := con.Param("id")
	db.Delete(removeId)
	con.JSON(http.StatusOK, "削除")
}

func CreateAutoAction(c *gin.Context, db *gorm.DB) {
	autoAction := model.AutoAction{}
	println(autoAction)
	err := c.BindJSON(&autoAction)
	if err != nil {
		c.JSON(http.StatusBadRequest, "エラー"+err.Error())
		return
	}
	db.Create(&autoAction)
	c.JSON(http.StatusOK, autoAction)

}
