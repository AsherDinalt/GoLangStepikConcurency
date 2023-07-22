package main

import (
	"fmt"
	"time"
)

func main_03() {
	messages := make(chan string)

	go func() {
		fmt.Println("B: Sending message...")
		time.Sleep(500 * time.Millisecond)
		messages <- "ping"              // (1)
		fmt.Println("B: Message sent!") // (2)
	}()

	fmt.Println("A: Doing some work...")
	//time.Sleep(500 * time.Millisecond)
	fmt.Println("A: Ready to receive a message...")

	<-messages //  (3) тоже блокируется

	fmt.Println("A: Messege received!")
	time.Sleep(100 * time.Millisecond)
}

func main_1() {
	// Канал создается через `make(chan тип)`
	// и может передавать только значения указанного типа:
	messages := make(chan string)

	// Чтобы отправить значение в канал,
	// используют синтаксис `канал <-`
	// Отправим «пинг»:
	go func() { messages <- "ping" }()

	// Чтобы получить значение из канала,
	// используют синтаксис `<-канал`
	// Получим «пинг» и напечатаем его:
	msg := <-messages
	fmt.Println(msg)
}
