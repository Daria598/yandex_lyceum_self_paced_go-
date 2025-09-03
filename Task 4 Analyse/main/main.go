package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func getTopWords(wordMap map[string]int, n int) []string {

	keys := make([]int, 0, len(wordMap))
	values := make([]string, 0, len(wordMap))

	for v, k := range wordMap {
		keys = append(keys, k)
		values = append(values, v)
	}

	sort.Slice(keys, func(i, j int) bool { return keys[i] > keys[j] })

	wordsSlice := make([]string, 0)
	for i := 0; i < n; i++ {
		for words, value := range wordMap {
			if value == keys[i] {
				wordsSlice = append(wordsSlice, words)

			}
		}
	}

	return wordsSlice
}

func AnalyzeText(text string) {
	textNew := strings.ToLower(text)
	re := regexp.MustCompile(`[^a-zа-яё ]+`)
	textClean := re.ReplaceAllString(textNew, "")

	words := strings.Fields(textClean)
	wordsNumber := len(words)

	wordMap := make(map[string]int)
	for _, word := range words {
		_, exists := wordMap[word]
		if exists {
			wordMap[word] += 1
		} else {
			wordMap[word] = 1
		}
	}
	uniqueWords := len(wordMap)

	var maxNumber int
	for _, n := range wordMap {
		maxNumber = n
		break
	}
	for _, n := range wordMap {
		if n > maxNumber {
			maxNumber = n
		}
	}

	var mostPopularWord string
	for key, n := range wordMap {
		if n == maxNumber {
			mostPopularWord = key
			break
		}

	}
	// func getTopWords(wordMap map[string]int, n int) []string

	keys := make([]int, 0, len(wordMap))
	values := make([]string, 0, len(wordMap))

	for v, k := range wordMap {
		keys = append(keys, k)
		values = append(values, v)
	}

	sort.Slice(keys, func(i, j int) bool { return keys[i] > keys[j] })

	var word1 string
	var number1 int
	for words, value := range wordMap {
		if value == keys[0] {
			word1 = words
			number1 = keys[0]
			// fmt.Println(word1)
		}
	}

	var word2 string
	var number2 int
	for words, value := range wordMap {
		if value == keys[1] {
			word2 = words
			number2 = keys[1]
			// fmt.Println(word2)
		}
	}

	var word3 string
	var number3 int
	for words, value := range wordMap {
		if value == keys[2] {
			word3 = words
			number3 = keys[2]
			// fmt.Println(word3)
		}
	}

	var word4 string
	var number4 int
	for words, value := range wordMap {
		if value == keys[3] {
			word4 = words
			number4 = keys[3]
			// fmt.Println(word4)
		}
	}

	var word5 string
	var number5 int
	for words, value := range wordMap {
		if value == keys[4] {
			word5 = words
			number5 = keys[4]
			// fmt.Println(word5)
		}
	}

	fmt.Printf("Количество слов: %d\nКоличество уникальных слов: %d\nСамое часто встречающееся слово: \"%s\" (встречается %d раз)\nТоп-5 самых часто встречающихся слов:\n\"%s\": %d раз\n\"%s\": %d раз\n\"%s\": %d раз\n\"%s\": %d раз\n\"%s\": %d раз\n", wordsNumber, uniqueWords, mostPopularWord, maxNumber, word1, number1, word2, number2, word3, number3, word4, number4, word5, number5)
}
