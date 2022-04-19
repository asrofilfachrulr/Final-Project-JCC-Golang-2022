package handlers

import (
	sqlModels "anya-day/models/sql"
	webModels "anya-day/models/web"
	repo "anya-day/repository"
	services "anya-day/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary add new user
// @Description  add new user
// @Tags User
// @Param Body body webModels.RegisterUserInput true "User registration data. All required except address line"
// @Produce json
// @Success 201 {object} webModels.WithDataResp{data=webModels.RegisterUserOutput}
// @Router /user [post]
func PostUser(ctx *gin.Context) {
	// get needed repos for this endpoint
	userRepo := ctx.MustGet("user_repo").(*repo.UserRepository)
	userCredRepo := ctx.MustGet("user_cred_repo").(*repo.UserCredentialRepository)

	service := services.NewUserServices(userRepo, userCredRepo)

	var input webModels.RegisterUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    err.Error(),
		})
		return
	}

	user := &sqlModels.User{
		Fullname:    input.Fullname,
		Username:    input.Username,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		AddressLine: input.AddressLine,
	}

	userCred := &sqlModels.UserCredential{
		Password: input.Password,
	}

	if err := service.CreateUser(user, userCred); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, webModels.WithDataResp{
		Status: "success",
		Msg:    "user created",
		Data: webModels.RegisterUserOutput{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	})
}
