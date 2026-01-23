package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"strings"
	"time"
)

type Ticket struct {
	Ticket string    `json:"ticket"`
	User   string    `json:"user"`
	Status string    `json:"status"`
	Date   time.Time `json:"date"`
}

func GetTasks(
	ctx context.Context,
	r io.Reader,
	w io.Writer,
	user *string,
	status *string,
	timeout time.Duration,
) error {
	byteData, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	data := string(byteData)

	tickets := parseTickets(data, user, status)

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	js, err := json.MarshalIndent(tickets, "", "  ")
	if err != nil {
		return err
	}

	if _, err := w.Write(js); err != nil {
		return err
	}
	if _, err := w.Write([]byte("\n")); err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(timeout):
		return errors.New("timeout")
	default:
		return nil
	}
}

func parseTickets(data string, user *string, status *string) []Ticket {
	var tickets []Ticket

	lines := strings.Split(data, "\n")
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

		ticket := Ticket{
			Ticket: ticketID,
			User:   userName,
			Status: taskStatus,
			Date:   date,
		}
		tickets = append(tickets, ticket)
	}

	return tickets
}
