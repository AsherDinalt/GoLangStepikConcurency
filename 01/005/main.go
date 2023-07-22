package main

import (
	"strings"
)

// nextFunc returns the next word from the generator
type nextFunc func() string

// counter stores the number of digits in each word.
// each key is a word and value is the number of digits in the word.
type counter map[string]int

// pair хранит слово и количество цифр в нем
type pair struct {
	word  string
	count int
}

func main() {
	phrase := "0ne 1wo thr33 4068"
	next := wordGenerator(phrase)
	stats := countDigitsInWords(next)
	printStats(stats)

}

// wordGenerator returns a generator,
// which emits words from a phrase.
func wordGenerator(phrase string) nextFunc {
	words := strings.Fields(phrase)
	idx := 0
	return func() string {
		if idx == len(words) {
			return ""
		}
		word := words[idx]
		idx++
		return word
	}
}

// countDigitsInWords считает количество цифр в словах,
// выбирая очередные слова с помощью next()
func countDigitsInWords(next nextFunc) counter {

	// начало решения

	counted := make(chan pair)
	stats := counter{}
	var pr pair

	go func() {
		// Пройдите по словам,
		// посчитайте количество цифр в каждом,
		// и запишите его в канал counted
		for {
			word := next()
			if word == "" {
				pr.count = 0
				pr.word = ""
				counted <- pr
				break
			}
			count := countDigits(word)
			pr.count = count
			pr.word = word
			counted <- pr
		}

	}()

	for {
		prr := <-counted
		if prr.word == "" {
			break
		}
		stats[prr.word] = prr.count
	}

	// Считайте значения из канала counted
	// и заполните stats.

	// В результате stats должна содержать слова
	// и количество цифр в каждом.

	// конец решения

	return stats
}
