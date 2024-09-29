package dto

import (
	"time"

	"github.com/edgarmueller/go-api-journal/internal/domain"
)

type JournalEntryResponse struct {
	ID           uint      `json:"id"`
	Date         time.Time `json:"date"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Tasks        []string  `json:"tasks"`
	WorkingHours float64   `json:"workingHours"`
}

func ToJournalResponse(entries []domain.JournalEntry) []JournalEntryResponse {
	var response []JournalEntryResponse
	for _, entry := range entries {
		response = append(response, JournalEntryResponse{
			ID:           entry.ID,
			Date:         entry.Date,
			CreatedAt:    entry.CreatedAt,
			UpdatedAt:    entry.UpdatedAt,
			Tasks:        entry.Tasks,
			WorkingHours: entry.WorkingHours,
		})
	}
	return response
}
