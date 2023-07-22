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
	go submitWords(next, pending)

	counted := make(chan pair)
	go countWords(pending, counted)

	return fillStats(counted)

}

// начало решения

// submitWords отправляет слова на подсчет
func submitWords(next nextFunc, pending chan string) {
	for {
		word := next()
		if word == "" {
			pending <- word
			break
		}
		pending <- word
	}
}

// countWords считает цифры в словах
func countWords(pending chan string, counted chan pair) {
	var pr pair
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
}

// fillStats готовит итоговую статистику
func fillStats(counted chan pair) counter {
	stats := counter{}

	for {
		pr := <-counted
		if pr.word == "" {
			break
		}
		stats[pr.word] = pr.count
	}
	return stats
}

// конец решения
