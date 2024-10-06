package controllers

import (
	"grahamkatana/api/events/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    events,
	})
}

func CreateEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.BindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	id := ctx.GetUint("userId")
	event.UserID = id
	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data":    event,
	})
}

func GetEvent(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	event, err := models.GetEventByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    event,
	})
}

func UpdateEvent(ctx *gin.Context) {
	userId := ctx.GetUint("userId")
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	event, err := models.GetEventByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "error",
			"data":    "you are not authorized to update this event",
		})
		return
	}
	var updatedEvent models.Event
	err = ctx.BindJSON(&updatedEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	updatedEvent.ID = event.ID
	err = updatedEvent.UpdateEvent()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    updatedEvent,
	})

}

func DeleteEvent(ctx *gin.Context) {
	userId := ctx.GetUint("userId")
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	event, err := models.GetEventByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "error",
			"data":    "you are not authorized to delete this event",
		})
		return
	}
	err = event.DeleteEvent()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "deleted successfully",
		"data":    nil,
	})
}
