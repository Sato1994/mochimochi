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
	usersRoute := r.Group("/users")
	externalWebServiceRoute := r.Group("/external-web-service")
	externalWebServicesRoute := r.Group("/external-web-services")

	userRoute.GET("/:id", func(c *gin.Context) {
		GetUser(c, db)
	})
	userRoute.DELETE("/:id", func(c *gin.Context) {
		DeleteUser(c, db)
	})
	usersRoute.GET("/", func(c *gin.Context) {
		GetUsers(c, db)
	})

	externalWebServiceRoute.GET("/:id", func(c *gin.Context) {
		GetExternalWebService(c, db)
	})
	externalWebServiceRoute.POST("/", func(c *gin.Context) {
		CreateExternalWebService(c, db)
	})
	externalWebServicesRoute.GET("/", func(c *gin.Context) {
		GetExternalWebServices(c, db)
	})

}
