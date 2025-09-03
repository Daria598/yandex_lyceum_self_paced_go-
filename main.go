package main

import (
	"errors"
	"fmt"
	"sort"
)

type Worker struct {
	Name       string
	Position   string
	Salary     uint
	Experience uint
}

type Company struct {
	Workers []Worker
}

type CompanyInterface interface {
	AddWorkerInfo(name, position string, salary, experience uint) error
	SortWorkers() ([]string, error)
}

var positionRank = map[string]int{
	"директор":       5,
	"зам. директора": 4,
	"начальник цеха": 3,
	"мастер":         2,
	"рабочий":        1,
}

func (c *Company) AddWorkerInfo(name, position string, salary, experience uint) error {
	if position != "директор" && position != "зам. директора" && position != "начальник цеха" && position != "мастер" && position != "рабочий" {
		return errors.New("несуществующая должность")
	}
	c.Workers = append(c.Workers, Worker{name, position, salary, experience})
	return nil

}

func (c *Company) SortWorkers([]string, error) {
	sort.SliceStable(c.Workers, func(i, j int) bool {
		incomeI := c.Workers[i].Salary * c.Workers[i].Experience
		incomeJ := c.Workers[j].Salary * c.Workers[j].Experience
		if incomeI != incomeJ {
			return incomeI > incomeJ // по убыванию дохода
		}
		// если доход равен, сортируем по должности по убыванию
		return positionRank[c.Workers[i].Position] > positionRank[c.Workers[j].Position]
	})

	var result []string
	for _, w := range c.Workers {
		income := w.Salary * w.Experience
		result = append(result, fmt.Sprintf("%s — %d — %s", w.Name, income, w.Position))
		return result, nil
	}

}

// type Interface interface {
// 	Len() int
// 	Less(i, j int) bool
// 	Swap(i, j int)
// }

// func (w Worker) Len() int {
// 	return len(w.Names)
// }

// func (w Worker) Swap(i, j int) {
// 	w.Names[i], w.Names[j] = w.Names[j], w.Names[i]
// 	w.Salaries[i], w.Salaries[j] = w.Salaries[j], w.Salaries[i]
// 	w.Positions[i], w.Positions[j] = w.Positions[j], w.Positions[i]
// }

// func (w Worker) Less(i, j int) bool {
// 	return w.Salaries[i] < w.Salaries[j]
// }
