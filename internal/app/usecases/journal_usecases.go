package usecases

import (
	"time"

	"github.com/edgarmueller/go-api-journal/internal/domain"
	"github.com/edgarmueller/go-api-journal/internal/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JournalUseCases struct {
	journalRepository repositories.JournalRepository
	userRepository    repositories.UserRepository
}

type UpsertEntry struct {
	Tasks        []string `json:"tasks"`
	WorkingHours float64  `json:"workingHours"`
}

func NewJournalUseCases(journalRepository repositories.JournalRepository, userRepository repositories.UserRepository) *JournalUseCases {
	return &JournalUseCases{
		journalRepository: journalRepository,
		userRepository:    userRepository,
	}
}

func (u *JournalUseCases) GetEntries(userId string) (*domain.WorkJournal, error) {
	uuid, err := uuid.Parse(userId)

	if err != nil {
		return &domain.WorkJournal{}, err
	}

	user, err := u.userRepository.GetUserByUUID(uuid)

	if err != nil {
		return &domain.WorkJournal{}, err
	}

	journal, err := u.journalRepository.GetDefaultJournalByUserId(user.ID)
	if err != nil {
		return &domain.WorkJournal{}, err
	}

	return journal, nil
}

func (u *JournalUseCases) UpsertEntry(date time.Time, userId string, upsertEntry UpsertEntry) (*domain.JournalEntry, error) {
	uuid, err := uuid.Parse(userId)

	if err != nil {
		return &domain.JournalEntry{}, err
	}

	user, err := u.userRepository.GetUserByUUID(uuid)

	if err != nil {
		return &domain.JournalEntry{}, err
	}

	journal, err := u.journalRepository.GetDefaultJournalByUserIdWithEntry(user.ID, date)
	if !journal.HasEntryForDate(date) {
		// for now we only have a single (default) journal
		e := journal.AddEntry(domain.AddEntry{
			Date:         date,
			WorkingHours: upsertEntry.WorkingHours,
			Tasks:        upsertEntry.Tasks,
		})

		err = u.journalRepository.SaveJournal(journal)

		if err != nil {
			return &domain.JournalEntry{}, err
		}

		return &e, nil
	} else if err != nil {
		return &domain.JournalEntry{}, err
	}

	e := journal.EditEntry(domain.EditEntry(domain.EditEntry{
		Date:         date,
		WorkingHours: upsertEntry.WorkingHours,
		Tasks:        upsertEntry.Tasks,
	}))
	err = u.journalRepository.SaveJournal(journal)
	if err != nil {
		return &domain.JournalEntry{}, err
	}
	return &e, nil
}

func (u *JournalUseCases) CreateDefaultJournalForUser(user domain.User) error {
	_, err := u.journalRepository.GetDefaultJournalByUserId(user.ID)

	if err == gorm.ErrRecordNotFound {
		journal := domain.CreateDefaultJournalForUser(user)
		return u.journalRepository.SaveJournal(&journal)
	}

	return nil
}
