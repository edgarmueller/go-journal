package controllers

import (
	"time"

	"github.com/edgarmueller/go-api-journal/internal/adapters/controllers/dto"
	"github.com/edgarmueller/go-api-journal/internal/adapters/middlewares"
	"github.com/edgarmueller/go-api-journal/internal/app/usecases"
	"github.com/gin-gonic/gin"
)

type JournalAPIController struct {
	journal *usecases.JournalUseCases
}

func NewJournalAPIController(router *gin.RouterGroup, journal *usecases.JournalUseCases) {
	controller := &JournalAPIController{
		journal: journal,
	}
	router.PUT("/journal/:date", middlewares.Auth(false), controller.UpsertEntry)
	router.GET("/journal", middlewares.Auth(false), controller.GetEntries)
}

func (u *JournalAPIController) GetEntries(ctx *gin.Context) {
	userID, exists := ctx.Get("UserUUID")

	if !exists {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}

	journal, err := u.journal.GetEntries(userID.(string))

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(200, gin.H{"entries": dto.ToJournalResponse(journal.Entries)})
}

func (u *JournalAPIController) UpsertEntry(ctx *gin.Context) {
	var upsertEntry usecases.UpsertEntry

	if err := ctx.ShouldBindJSON(&upsertEntry); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	userID, _ := ctx.Get("UserUUID")
	dateStr := ctx.Param("date")
	parsedDate, err := time.Parse("2006-01-02", dateStr)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	_, err = u.journal.UpsertEntry(parsedDate, userID.(string), upsertEntry)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(200, gin.H{"entry": upsertEntry})
}
