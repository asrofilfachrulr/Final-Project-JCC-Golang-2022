package routes

import (
	"net/http"

	"anya-day/handlers"
	repo "anya-day/repository"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// db type inverted to any there is multiple db type
func AttachRepo(r *gin.Engine, db *gorm.DB) {
	r.Use(func(ctx *gin.Context) {
		ctx.Set("user_repo", repo.NewUserRepo(db))
		ctx.Set("user_cred_repo", repo.NewUserCredRepo(db))
	})
}

func InitAPI(r *gin.Engine) *gin.Engine {
	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	api := r.Group("/api/v1")

	api.POST("/user", handlers.PostUser)

	// swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
