package handlers

import (
	"context"

	"github.com/asaskevich/EventBus"
	"github.com/edgarmueller/go-api-journal/internal/app/usecases"
	"github.com/edgarmueller/go-api-journal/internal/domain"
	"google.golang.org/appengine/v2/log"
)

type UserCreatedHandler struct {
	eventBus EventBus.Bus
	journal  *usecases.JournalUseCases
}

func NewUserCreatedHandler(journal *usecases.JournalUseCases, eventBus EventBus.Bus) {
	handler := UserCreatedHandler{journal: journal}
	eventBus.Subscribe("user.created", handler.Handle)
}

func (u *UserCreatedHandler) Handle(user domain.User) {
	err := u.journal.CreateDefaultJournalForUser(user)
	if err != nil {
		log.Errorf(context.Background(), "Error creating default journal for user %v: %v", user.ID, err)
	}
}
