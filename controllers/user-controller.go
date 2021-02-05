package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seigaalghi/firestore-go/models"
	"github.com/seigaalghi/firestore-go/repository"
	"github.com/seigaalghi/firestore-go/utility"
)

// Register is ..
func Register(ctx *gin.Context) {
	var input models.User
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	password, err := utility.HashPassword(input.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data, err := repository.SaveUser(&models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: password,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// Login is ..
func Login(ctx *gin.Context) {
	var input models.LoginUser
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := repository.FindUser(input.Email)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token := utility.GenerateToken(input.Email)

	if utility.CheckPassword(input.Password, user.Password) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": token,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"data": "Invalid email or password",
		})
	}

}
