// routes.go

package routes

import (
	control "coinproject/controller"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("./templates/*")
	router.GET("/", control.GetStatus)
	router.GET("/status/:id", control.OriginalPage)

	v1 := router.Group("/v1")
	{
		v1.GET("/url/:id", control.GetUrlDetail)
		v1.GET("/urls/", control.GetUrls)
	}
	return router
}
