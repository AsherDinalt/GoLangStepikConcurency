package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func main_001() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		say(1, "go is awesome")
	}()
	go func() {
		defer wg.Done()
		say(2, "cats are cute")
	}()

	wg.Wait()
}

func say(id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
}
