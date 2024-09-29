package repositories

import (
	"time"

	"github.com/edgarmueller/go-api-journal/internal/domain"
)

type JournalRepository interface {
	SaveJournal(journal *domain.WorkJournal) error
	GetDefaultJournalByUserId(userId uint) (*domain.WorkJournal, error)
	GetDefaultJournalByUserIdWithEntry(userID uint, date time.Time) (*domain.WorkJournal, error)
}
