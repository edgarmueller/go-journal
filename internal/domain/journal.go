package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"sort"
	"time"

	"gorm.io/gorm"
)

type WorkJournal struct {
	gorm.Model
	ID        uint `gorm:"unique;primaryKey;autoIncrement"`
	Title     string
	OwnerId   uint `gorm:"unique"`
	IsDefault bool
	Entries   []JournalEntry
}

type JournalEntry struct {
	gorm.Model
	ID            uint `gorm:"unique;primaryKey;autoIncrement"`
	WorkJournalID uint
	Tasks         JSONB     `gorm:"type:jsonb;default:'[]';not null"`
	Date          time.Time `gorm:"type:date"`
	WorkingHours  float64
	OwnerId       uint
}

type AddEntry struct {
	Date         time.Time
	WorkingHours float64
	OwnerId      uint
	Tasks        JSONB
}

type EditEntry struct {
	Date         time.Time
	WorkingHours float64
	Tasks        JSONB
}

func (w *WorkJournal) AddEntry(entry AddEntry) JournalEntry {
	w.Entries = append(w.Entries, createJournalEntry(
		w,
		entry.Date,
		entry.WorkingHours,
		w.OwnerId,
		entry.Tasks[0],
	))
	return w.Entries[len(w.Entries)-1]
}

func (w *WorkJournal) EditEntry(entry EditEntry) JournalEntry {
	for i, e := range w.Entries {
		if e.Date.Day() == entry.Date.Day() && e.Date.Month() == entry.Date.Month() && e.Date.Year() == entry.Date.Year() {
			w.Entries[i].WorkingHours = entry.WorkingHours
			w.Entries[i].Tasks = entry.Tasks
			return w.Entries[i]
		}
	}
	return JournalEntry{}
}

func (wj *WorkJournal) HasEntryForDate(date time.Time) bool {
	for _, entry := range wj.Entries {
		if entry.Date.Day() == date.Day() && entry.Date.Month() == date.Month() && entry.Date.Year() == date.Year() {
			return true
		}
	}

	return false
}

func (wj *WorkJournal) SortEntriesByDate() {
	sort.Slice(wj.Entries, func(i, j int) bool {
		return wj.Entries[i].Date.After(wj.Entries[j].Date)
	})
}

func CreateDefaultJournalForUser(user User) WorkJournal {
	return WorkJournal{
		OwnerId:   user.ID,
		Title:     "Default Journal",
		IsDefault: true,
	}
}

func CreateJournalForUser(user User) WorkJournal {
	return WorkJournal{
		OwnerId:   user.ID,
		Title:     "Default Journal",
		IsDefault: true,
	}
}

func createJournalEntry(journal *WorkJournal, date time.Time, workingHours float64, ownerId uint, task string) JournalEntry {
	return JournalEntry{
		Date:          date,
		WorkingHours:  workingHours,
		OwnerId:       ownerId,
		WorkJournalID: journal.ID,
		Tasks:         []string{task},
	}
}

// JSON serialization
type JSONB []string

func (a JSONB) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
