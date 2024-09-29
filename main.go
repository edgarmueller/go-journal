package main

import (
	"os"

	"github.com/asaskevich/EventBus"
	controllers "github.com/edgarmueller/go-api-journal/internal/adapters/controllers/api"
	presentation "github.com/edgarmueller/go-api-journal/internal/adapters/controllers/presentation"
	"github.com/edgarmueller/go-api-journal/internal/adapters/database"
	"github.com/edgarmueller/go-api-journal/internal/app/handlers"
	"github.com/edgarmueller/go-api-journal/internal/app/services"
	"github.com/edgarmueller/go-api-journal/internal/app/usecases"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Connect(os.Getenv("DATABASE_URL"))
	database.Migrate()

	bus := EventBus.New()

	// Repos --
	userRepository := database.NewGormUserRepository(db, bus)
	journalRepository := database.NewGormJournalRepository(db)

	// Use cases --
	authService := services.NewAuthService()
	authUseCases := usecases.NewAuthUseCases(userRepository, authService)
	userUseCases := usecases.NewUserUseCases(userRepository, authService)
	journalUseCases := usecases.NewJournalUseCases(journalRepository, userRepository)

	// Event handlers --
	handlers.NewUserCreatedHandler(journalUseCases, bus)

	router := gin.Default()
	api := router.Group("/api")

	// API --
	controllers.NewUserAPIController(api, authUseCases, userUseCases)
	controllers.NewTokenAPIController(api, authUseCases)
	controllers.NewJournalAPIController(api, journalUseCases)

	// Presentation --
	ginHtmlRenderer := router.HTMLRender
	router.HTMLRender = &presentation.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	presentation.NewJournalController(router, journalUseCases)
	presentation.NewAuthController(router, authUseCases, userUseCases)

	router.Run("localhost:9090")
}
