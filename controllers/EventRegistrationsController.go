package controllers

import (
	"grahamkatana/api/events/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BookEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseUint(ctx.Param("eventId"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	var eventRegistration models.Registration
	err = ctx.BindJSON(&eventRegistration)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	userId := ctx.GetUint("userId")
	eventRegistration.UserID = int(userId)
	eventRegistration.EventID = int(eventId)
	err = eventRegistration.Save()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    eventRegistration,
	})

}

func CancelEvent(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	err = models.DeleteRegistration(int(id), int(ctx.GetUint("userId")))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    nil,
	})

}
