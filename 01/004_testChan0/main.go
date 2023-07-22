package main

import (
	"fmt"
	"strings"
	"unicode"
)

type counter map[string]int

func main() {
	phrase := "0ne 1wo thr33 4068"
	stats := countDigitsInWords(phrase)
	printStats(stats)
}

// countDigitsInWords считает количество цифр в словах фразы
func countDigitsInWords(phrase string) counter {
	words := strings.Fields(phrase)
	counted := make(chan int)
	stats := counter{}
	// начало решения
	go func() {
		// Пройдите по словам,
		// посчитайте количество цифр в каждом,
		// и запишите его в канал counted
		for _, word := range words {
			c := countDigits(word)
			counted <- c
		}
	}()

	for _, word := range words {
		stats[word] = <-counted
	}

	// Считайте значения из канала counted
	// и заполните stats.

	// В результате stats должна содержать слова
	// и количество цифр в каждом.

	// конец решения

	return stats
}

// countDigits возвращает количество цифр в строке
func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// printStats печатает слова и количество цифр в каждом
func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}
