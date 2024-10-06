package controllers

import (
	"grahamkatana/api/events/models"
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
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
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
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
