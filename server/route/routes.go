package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(
	r *gin.Engine,
	db *gorm.DB,
) {
	userRoute := r.Group("/user")
	userRoute.GET("/:id", func(c *gin.Context) {
		GetUser(c, db)
	})
	userRoute.DELETE("/:id", func(c *gin.Context) {
		DeleteUser(c, db)
	})

	usersRoute := r.Group("/users")
	usersRoute.GET("/", func(c *gin.Context) {
		GetUsers(c, db)
	})

	externalWebServiceRoute := r.Group("/external-web-service")
	externalWebServiceRoute.GET("/:id", func(c *gin.Context) {
		GetExternalWebService(c, db)
	})
	externalWebServiceRoute.POST("/", func(c *gin.Context) {
		CreateExternalWebService(c, db)
	})

	externalWebServicesRoute := r.Group("/external-web-services")
	externalWebServicesRoute.GET("/", func(c *gin.Context) {
		GetExternalWebServices(c, db)
	})

	autoActionsRoute := r.Group("/auto-actions")
	autoActionsRoute.GET("/", func(c *gin.Context) {
		GetAutoActions(c, db)
	})

	autoActionRoute := r.Group("/auto-action")
	autoActionRoute.GET("/:url", func(c *gin.Context) {
		GetHtmlFromPath(c)
	})
	autoActionRoute.DELETE("/", func(c *gin.Context) {
		DeleteHtmlFromPath(c, db)
	})
	autoActionRoute.POST("/", func(c *gin.Context) {
		CreateAutoAction(c, db)
	})

}
