package main

import (
	"imageApi/controllers"
	_ "imageApi/docs"
	"imageApi/models"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title     Images API
// @version         1.0
// @description     An Image API in Go using Gin framework.

// @contact.name   Tim Seitzinger
// @contact.email  tseitzinger@gmail.com

// @host      localhost:8080
// @BasePath  /
func main() {

	r := setupRouter()

	r.Run()

}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	models.ConnectDatabase()

	r.GET("/images", controllers.FindImages)

	r.GET("/images/:image_id", controllers.FindImage)

	r.POST("/images", controllers.CreateImage)

	r.DELETE("/images/:image_id", controllers.DeleteImage)

	r.PATCH("/images/:image_id", controllers.UpdateImage)

	return r
}
