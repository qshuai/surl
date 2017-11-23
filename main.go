package main

import (
	_ "surl/app"
	_ "surl/storage"

	"surl/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/store", controllers.StoreUrl)
	r.GET("/:surl", controllers.GetUrl)
	r.Run(":8080")
}
