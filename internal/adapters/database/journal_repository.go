package database

import (
	"time"

	"github.com/edgarmueller/go-api-journal/internal/domain"
	"github.com/edgarmueller/go-api-journal/internal/domain/repositories"
	"gorm.io/gorm"
)

type GormJournalRepository struct {
	db *gorm.DB
}

func NewGormJournalRepository(db *gorm.DB) repositories.JournalRepository {
	return &GormJournalRepository{db: db}
}

func (g *GormJournalRepository) SaveJournalEntry(entry domain.JournalEntry) error {
	if entry.ID != 0 {
		err := g.db.Save(&entry)
		if err := err.Error; err != nil {
			return err
		}
		return nil
	}
	e := g.db.Create(&entry)
	if err := e.Error; err != nil {
		return err
	}
	return nil
}

func (g *GormJournalRepository) GetDefaultJournalByUserId(userId uint) (*domain.WorkJournal, error) {
	var journal domain.WorkJournal
	result := g.db.Model(&domain.WorkJournal{}).Preload("Entries").First(&journal, "owner_id = ? AND is_default = ?", userId, true)
	if result.Error != nil {
		return &journal, result.Error
	}
	journal.SortEntriesByDate()
	return &journal, nil
}

func (g *GormJournalRepository) GetDefaultJournalByUserIdWithEntry(userID uint, date time.Time) (*domain.WorkJournal, error) {
	var journal domain.WorkJournal
	result := g.db.Model(&domain.WorkJournal{}).Preload("Entries", "date = ?", date).First(&journal, "owner_id = ? AND is_default = ?", userID, true)
	if result.Error != nil {
		return &journal, result.Error
	}
	return &journal, nil
}

func (g *GormJournalRepository) SaveJournal(journal *domain.WorkJournal) error {
	if journal.CreatedAt != (time.Time{}) {
		for _, entry := range journal.Entries {
			err := g.SaveJournalEntry(entry)

			if err != nil {
				return err
			}
		}
		err := g.db.Omit("Entries").Updates(journal)
		if err := err.Error; err != nil {
			return err
		}
	} else {
		e := g.db.Create(&journal)
		if err := e.Error; err != nil {
			return err
		}
		for _, entry := range journal.Entries {
			entry.WorkJournalID = journal.ID
			err := g.SaveJournalEntry(entry)

			if err != nil {
				return err
			}
		}
	}
	return nil
}
