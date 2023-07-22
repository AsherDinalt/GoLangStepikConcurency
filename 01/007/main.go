package main

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

// countDigitsInWords считает количество цифр в словах,
// выбирая очередные слова с помощью next()
func countDigitsInWords(next nextFunc) counter {
	pending := make(chan string)
	counted := make(chan pair)

	// начало решения

	stats := counter{}
	var pr pair

	go func() {
		// Пройдите по словам,
		// посчитайте количество цифр в каждом,
		// и запишите его в канал counted
		for {
			word := next()
			if word == "" {
				pending <- word
				break
			}
			pending <- word
		}

	}()

	go func() {
		// Считайте слова из канала pending,
		// посчитайте количество цифр в каждом,
		// и запишите его в канал counted
		for {
			word := <-pending
			if word == "" {
				pr.word = ""
				pr.count = 0
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
