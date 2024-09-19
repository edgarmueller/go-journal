package app

import "github.com/edgarmueller/go-api-journal/internal/app/command"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	// CreateJournal command.CreateJournalHandler
	// DeleteJournal command.DeleteJournalHandler
	// EditJournal   command.EditJournalHandler

	CreateJournalEntry command.CreateJournalEntryHandler
	// DeleteJournalEntry command.DeleteJournalEntryHandler
	// EditJournalEntry   command.EditJournalEntryHandler
}

type Queries struct {
	// HourAvailability      query.HourAvailabilityHandler
	// TrainerAvailableHours query.AvailableHoursHandler
}
