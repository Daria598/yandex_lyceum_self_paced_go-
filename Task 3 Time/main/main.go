package main

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

func currentDayOfTheWeek() string {

	weekday := time.Now().Weekday() - 2

	a := "Понедельник"
	b := "Вторник"
	c := "Среда"
	d := "Четверг"
	e := "Пятница"
	f := "Суббота"
	g := "Воскресенье"

	if weekday.String() == "Monday" {
		return a
	} else if weekday.String() == "Tuesday" {
		return b
	} else if weekday.String() == "Wednesday" {
		return c
	} else if weekday.String() == "Thursday" {
		return d
	} else if weekday.String() == "Friday" {
		return e
	} else if weekday.String() == "Saturday" {
		return f
	} else {
		return g
	}
}

func dayOrNight() string {
	now := time.Now()
	DayOrNight := now.Hour()
	if DayOrNight >= 10 && DayOrNight <= 22 {
		a := "День"
		return a
	} else {
		b := "Ночь"
		return b
	}
}

func nextFriday() int {
	weekday := time.Now().Weekday() - 2
	if weekday.String() == "Monday" {
		return 4
	} else if weekday.String() == "Tuesday" {
		return 3
	} else if weekday.String() == "Wednesday" {
		return 2
	} else if weekday.String() == "Thursday" {
		return 1
	} else if weekday.String() == "Friday" {
		return 0
	} else if weekday.String() == "Saturday" {
		return 6
	} else {
		return 5
	}
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	if strings.ToLower(answer) == strings.ToLower(currentDayOfTheWeek()) {
		return true
	} else {
		return false
	}
}

var (
	Mistake = errors.New("исправь свой ответ, а лучше ложись поспать")
)

func CheckNowDayOrNight(answer string) (bool, error) {
	now := time.Now()
	DayOrNight := now.Hour()
	if utf8.RuneCountInString(answer) != 4 {
		return false, Mistake // нужна пустая строка вместо false
	} else if strings.ToLower(answer) == "день" && DayOrNight >= 10 && DayOrNight <= 22 || strings.ToLower(answer) == "ночь" && DayOrNight > 22 || DayOrNight < 10 {
		return true, nil
	} else {
		return false, nil
	}
}
