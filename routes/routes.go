package routes

import (
	"log"
	"net/http"

	"anya-day/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Attach(r *gin.Engine, key string, data any) {
	r.Use(func(ctx *gin.Context) {
		ctx.Set(key, data)
	})
}

func InitAPI(r *gin.Engine) *gin.Engine {
	log.Println("Initiating route API...")
	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	api := r.Group("/api/v1")

	api.POST("/user", handlers.PostUserHandler)

	// swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
