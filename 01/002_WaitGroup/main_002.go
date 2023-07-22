package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
	"unicode"
)

type counter map[string]int

func main_002() {
	phrase := "0ne 1wo thr33 4068"
	counts := countDigitsInWords(phrase)
	printStats(counts)
}

func countDigitsInWords(phrase string) counter {
	words := strings.Fields(phrase)

	var wg sync.WaitGroup
	syncStats := sync.Map{}

	// начало решения
	for _, word := range words {
		wg.Add(1)
		go func(ww string, wgg *sync.WaitGroup) {
			defer wg.Done()
			c := countDigits(ww)
			syncStats.Store(ww, c)
		}(word, &wg)
	}
	wg.Wait()
	// конец решения
	return asStats(&syncStats)
}

// asStats converts stats from sync.Map to ordinary map
func asStats(m *sync.Map) counter {
	stats := counter{}
	m.Range(func(word, count any) bool {
		stats[word.(string)] = count.(int)
		return true
	})
	return stats
}

// printStats prints words and their digit counts
func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

func test_001() {
	values := []int{1, 2, 3, 4, 5}
	for _, val := range values {
		go func(ee int) {
			fmt.Printf("%d ", ee)
		}(val)
	}

	time.Sleep(time.Second)
}
