package main

import (
	"fmt"
	"strconv"
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
	var tickets []Ticket
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		parts := strings.Split(line, "_")
		if len(parts) != 4 {
			continue
		}

		if !strings.HasPrefix(parts[0], "TICKET") {
			continue
		}

		ticketNumStr := strings.TrimPrefix(parts[0], "TICKET-")
		if _, err := strconv.Atoi(ticketNumStr); err != nil {
			continue
		}

		userName := parts[1]
		if userName == "" {
			continue
		}

		if parts[2] != "Готово" && parts[2] != "В работе" && parts[2] != "Не будет сделано" {
			continue
		}

		date, err := time.Parse("2006-01-02", parts[3])
		if err != nil {
			continue
		}

		ticket := Ticket{
			Ticket: parts[0],
			User:   parts[1],
			Status: parts[2],
			Date:   date,
		}

		if user != nil && ticket.User != *user {
			continue
		}

		if status != nil && ticket.Status != *status {
			continue
		}

		tickets = append(tickets, ticket)
		// if ticket.User == "Паша Попов" {
		// 	fmt.Println(ticket)
		// }
	}

	return tickets

}

func main() {
	chatHistory := `
TICKET-12345_Паша Попов_Готово_2024-01-01
TICKET-12346_Иван Иванов_В работе_2024-01-02
TICKET-12347_Анна Смирнова_Не будет сделано_2024-01-03
TICKET-12348_Паша Попов_В работе_2024-01-04
TICKET-53543_Иван Иванов_Готово_2024-01-03
TICKET-12354_Паша Попов_Самый лучший день_2024-01-04
TICKET-fhffdf_Иван Иванов_Готово_2024-01-08
`

	user := "Паша Попов"
	tasks := GetTasks(chatHistory, &user, nil)
	fmt.Println("Задачи пользователя Паша Попов:")
	for _, t := range tasks {
		fmt.Printf("%+v\n", t)
	}

	user1 := "Иван Иванов"
	status1 := "В работе"
	tasks1 := GetTasks(chatHistory, &user1, &status1)
	fmt.Println("\nЗадачи пользователя Иван Иванов:")
	fmt.Println(tasks1)
	// for _, t := range tasks1 {
	// 	fmt.Printf("%+v\n", t)
	// }

	status := "В работе"
	tasksByStatus := GetTasks(chatHistory, nil, &status)
	fmt.Println("\nЗадачи со статусом В работе:")
	for _, t := range tasksByStatus {
		fmt.Printf("%+v\n", t)
	}
}
