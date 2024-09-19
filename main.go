package main

import (
	"net/http"

	"github.com/edgarmueller/go-api-journal/internal/adapters/controllers"
	"github.com/edgarmueller/go-api-journal/internal/adapters/database"
	"github.com/edgarmueller/go-api-journal/internal/app/services"
	"github.com/edgarmueller/go-api-journal/internal/app/usecases"
	"github.com/gin-gonic/gin"
)

type journalEntry struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Date    string `json:"date"`
	Content string `json:"content"`
}

var journal = []journalEntry{
	{ID: "1", Title: "First Entry", Date: "2024-09-01", Content: "This is the first entry in the journal."},
	{ID: "2", Title: "Second Entry", Date: "2024-09-02", Content: "This is the second entry in the journal."},
	{ID: "3", Title: "Third Entry", Date: "2024-09-03", Content: "This is the third entry in the journal."},
}

func getJournalEntries(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, journal)
}

func main() {
	// TODO
	db := database.Connect("host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable")
	database.Migrate()

	userRepository := database.NewGormUserRepository(db)

	authService := services.NewAuthService()
	authUseCases := usecases.NewAuthUseCases(userRepository, authService)

	router := gin.Default()
	api := router.Group("/api")

	controllers.NewPingController(api)
	controllers.NewUserController(api, authUseCases)
	controllers.NewTokenController(api, authUseCases)

	router.GET("/journal", getJournalEntries)
	router.Run("localhost:9090")
}
