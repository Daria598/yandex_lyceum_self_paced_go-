package main

import (
	"strings"
	"time"
)

type Ticket struct {
	Ticket string
	User   string
	Status string
	Date   time.Time
}

func GetTasks(text string, user *string, status *string) []Ticket {
	var result []Ticket

	if text == "" {
		return result
	}

	lines := strings.Split(text, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if !strings.HasPrefix(line, "TICKET-") {
			continue
		}

		parts := strings.Split(line, "_")
		if len(parts) != 4 {
			continue
		}

		ticketID := parts[0]
		userName := parts[1]
		taskStatus := parts[2]
		dateStr := parts[3]

		switch taskStatus {
		case "Готово", "В работе", "Не будет сделано":

		default:

			continue
		}

		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			continue
		}

		if user != nil && *user != userName {
			continue
		}

		if status != nil && *status != taskStatus {
			continue
		}

		result = append(result, Ticket{
			Ticket: ticketID,
			User:   userName,
			Status: taskStatus,
			Date:   date,
		})
	}

	return result
}
