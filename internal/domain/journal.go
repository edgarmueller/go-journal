package domain

import "time"

type WorkJournal struct {
	ID      string
	Title   string
	Entries []WorkJournalEntry
}

type WorkJournalEntry struct {
	ID           string    `json:"id"`
	Date         time.Time `json:"date"`
	WorkingHours float64   `json:"workingHours"`
	Tasks        []string  `json:"tasks"`
}

func (j *WorkJournal) AddEntry(entry WorkJournalEntry) {
	j.Entries = append(j.Entries, entry)
}
