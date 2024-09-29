package presentation

import (
	"net/http"
	"strconv"
	"time"

	presentation "github.com/edgarmueller/go-api-journal/internal/adapters/controllers/presentation/templates"
	"github.com/edgarmueller/go-api-journal/internal/adapters/middlewares"
	"github.com/edgarmueller/go-api-journal/internal/app/usecases"
	"github.com/gin-gonic/gin"
)

type JournalController struct {
	journal *usecases.JournalUseCases
}

func NewJournalController(router *gin.Engine, journal *usecases.JournalUseCases) {
	controller := &JournalController{
		journal: journal,
	}

	router.GET("/journal", middlewares.Auth(true), controller.GetEntries)
	router.POST("/journal", middlewares.Auth(true), controller.SubmitAddEntry)
	router.PUT("/journal", middlewares.Auth(true), controller.SubmitEditEntry)
}

func (c *JournalController) GetEntries(ctx *gin.Context) {
	userId, exists := ctx.Get("UserUUID")

	if !exists {
		r := New(ctx.Request.Context(), http.StatusOK, presentation.Error("401", "Unauthorized"))
		ctx.Render(http.StatusOK, r)
		return
	}

	journal, err := c.journal.GetEntries(userId.(string))

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	r := New(ctx.Request.Context(), http.StatusOK, presentation.Journal(journal.Entries))
	ctx.Render(http.StatusOK, r)
}

func (c *JournalController) SubmitAddEntry(ctx *gin.Context) {
	dateStr := ctx.PostForm("date")
	task := ctx.PostForm("tasks")
	workingHours := ctx.PostForm("workingHours")
	userId, exists := ctx.Get("UserUUID")

	if !exists {
		r := New(ctx.Request.Context(), http.StatusOK, presentation.Error("500", "Unauthorized"))
		ctx.Render(http.StatusOK, r)
		return
	}

	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	if !exists {
		r := New(ctx.Request.Context(), http.StatusOK, presentation.Error("401", "Unauthorized"))
		ctx.Render(http.StatusOK, r)
		return
	}

	hours, err := strconv.ParseFloat(workingHours, 64)

	if err != nil {
		r := New(ctx.Request.Context(), http.StatusOK, presentation.Error("400", "Invalid working hours"))
		ctx.Render(http.StatusOK, r)
		return
	}

	_, err = c.journal.UpsertEntry(parsedDate, userId.(string), usecases.UpsertEntry{
		Tasks:        []string{task},
		WorkingHours: hours,
	})

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	journal, err := c.journal.GetEntries(userId.(string))

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	journal.SortEntriesByDate()
	r := New(ctx.Request.Context(), http.StatusOK, presentation.Journal(journal.Entries))
	ctx.Render(http.StatusOK, r)
}

func (c *JournalController) SubmitEditEntry(ctx *gin.Context) {
	dateStr := ctx.PostForm("date")
	task := ctx.PostForm("tasks")
	workingHours := ctx.PostForm("workingHours")

	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		r := New(ctx.Request.Context(), http.StatusOK, presentation.Error("400", "Invalid date"))
		ctx.Render(http.StatusOK, r)
		return
	}
	hours, err := strconv.ParseFloat(workingHours, 64)
	if err != nil {
		r := New(ctx.Request.Context(), http.StatusOK, presentation.Error("400", "Invalid working hours"))
		ctx.Render(http.StatusOK, r)
		return
	}

	userId, exists := ctx.Get("UserUUID")
	if !exists {
		r := New(ctx.Request.Context(), http.StatusOK, presentation.Error("401", "Unauthorized"))
		ctx.Render(http.StatusOK, r)
		return
	}

	e, err := c.journal.UpsertEntry(parsedDate, userId.(string), usecases.UpsertEntry{
		Tasks:        []string{task},
		WorkingHours: hours,
	})

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	r := New(ctx.Request.Context(), http.StatusOK, presentation.Entry(e, false))
	ctx.Render(http.StatusOK, r)
}
