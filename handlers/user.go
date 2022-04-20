package handlers

import (
	"anya-day/helper"
	"anya-day/http/resp"
	sqlModels "anya-day/models/sql"
	webModels "anya-day/models/web"
	repo "anya-day/repository"
	services "anya-day/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

// Register godoc
// @Summary add new user
// @Description  add new user
// @Tags User
// @Param Body body webModels.RegisterUserInput true "User registration data. All required except address line"
// @Produce json
// @Success 201 {object} webModels.WithDataResp{data=webModels.RegisterUserOutput}
// @Router /user [post]
func PostUserHandler(ctx *gin.Context) {
	// get needed db and begin transaction for this endpoint
	tx := ctx.MustGet("db").(*gorm.DB).Begin()
	v := ctx.MustGet("validator").(*validator.Validate)

	if err := tx.Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    err.Error(),
		})
		return
	}

	defer helper.RollbackIfErr(tx)

	// initiate repo with tx
	userRepo := repo.NewUserRepo(tx)
	userCredRepo := repo.NewUserCredRepo(tx)

	service := services.NewUserServices(userRepo, userCredRepo)

	var input webModels.RegisterUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		resp.RespErrWithMsg(ctx, err.Error())
		return
	}

	if err := v.Struct(&input); err != nil {
		resp.RespErrWithMsg(ctx, err.Error())
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
		resp.RespErrWithMsg(ctx, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	tx.Commit()

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