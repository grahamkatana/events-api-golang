package controllers

import (
	"grahamkatana/api/events/models"
	"grahamkatana/api/events/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	err = models.CheckPassword(user.Email, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	data, err := models.GetUserByEmail(user.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	token, err := utils.GenerateJwtToken(data.Email, data.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": gin.H{
			"token": token,
			"id":    data.ID,
			"name":  data.Name,
			"email": data.Email,
		},
	})

}

func Register(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	token, err := utils.GenerateJwtToken(user.Email, user.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data": gin.H{
			"token": token,
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
