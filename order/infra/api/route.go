package api

import (
	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	router := gin.Default()
	group := router.Group("/coffee-shop")
	group.POST("/order", create)

	return router
}
