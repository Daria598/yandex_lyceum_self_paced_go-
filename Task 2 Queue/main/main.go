package main

import (
	"fmt"
)

func main() {

	var line [5]string = [5]string{"-", "-", "-", "-", "-"}

	var length int

	for {

		var name string
		var number int
		fmt.Scanln(&name, &number)

		if name == "очередь" {
			for i := 0; i < len(line); i++ {
				fmt.Printf("%d. %s\n", i+1, line[i])
			}
		} else if name == "количество" {
			free := len(line) - length
			fmt.Printf(
				fmt.Sprintf("Осталось свободных мест: %d\nВсего человек в очереди: %d\n", free, length),
			)
		} else if name == "конец" {
			for i := 0; i < len(line); i++ {
				fmt.Printf("%d. %s\n", i+1, line[i])
			}
			break
		} else if name == "" {
			continue
		} else if number <= 0 || number > 5 {
			fmt.Printf(
				fmt.Sprintf("Запись на место номер %d невозможна: некорректный ввод\n", number),
			)
		} else if length == 5 && number >= 1 && number <= 5 {
			fmt.Printf(
				fmt.Sprintf("Запись на место номер %d невозможна: очередь переполнена\n", number),
			)
		} else if line[number-1] != "-" {
			fmt.Printf(
				fmt.Sprintf("Запись на место номер %d невозможна: место уже занято\n", number),
			)
		} else if length < len(line) && number != 0 {
			line[number-1] = name
			length++
		}
	}
}


