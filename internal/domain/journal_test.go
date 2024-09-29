package domain

import (
	"testing"
	"time"
)

func TestAddEntry(t *testing.T) {
	j := WorkJournal{}
	j.AddEntry(AddEntry{
		Date:         time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		WorkingHours: 8,
		Tasks:        []string{"Task 1"},
	})

	if len(j.Entries) != 1 {
		t.Errorf("Expected 1 entry, got %d", len(j.Entries))
	}
}
