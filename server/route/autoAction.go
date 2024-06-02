package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gocolly/colly/v2"
)

func GetHtmlFromPath(con *gin.Context) {
	// https://github.com/gocolly/colly
	co := colly.NewCollector()

	url := "https://github.com/Sato1994"

	co.OnHTML("title", func(e *colly.HTMLElement) {
		print("Title:", e.Text)
	})

	co.OnRequest(func(r *colly.Request) {
		print("Visiting URL:", r.URL.String())
	})

	co.Visit(url)

	con.JSON(http.StatusOK, "ログを見てね")

}
