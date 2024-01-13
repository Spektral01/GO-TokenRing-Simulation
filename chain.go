package main

import (
	"fmt"
	"time"
)

type Token struct {
	data      string
	recipient int
	ttl       int
}

func Handling(id int, inChannel <-chan Token, outChannel chan<- Token) {
	for {
		token := <-inChannel
		if token.recipient == id {
			fmt.Printf("Канал %d получил сообщение: %s\n", id, token.data)
		} else {
			if token.ttl > 0 {
				fmt.Printf("TTL: %d \n", token.ttl)
				token.ttl--
				outChannel <- token
			} else {
				fmt.Printf("TTL истекло")
			}
		}
	}
}

func main() {
	var numChannel int
	fmt.Printf("Введите количество каналов: ")
	fmt.Scanln(&numChannel)

	channels := make([]chan Token, numChannel)
	for i := range channels {
		channels[i] = make(chan Token)
	}

	for i := 0; i < numChannel; i++ {
		go Handling(i, channels[i], channels[(i+1)%numChannel])
	}

	fmt.Printf("Введите стартовый канал: ")
	var start int
	fmt.Scanln(&start)

	fmt.Printf("Введите канал доставки (от 0 до %d): ", numChannel-1)
	var recipient int
	fmt.Scanln(&recipient)

	fmt.Printf("Введите TTL: ")
	var ttl int
	fmt.Scanln(&ttl)

	channels[start] <- Token{data: "Hello world!", recipient: recipient, ttl: ttl}

	time.Sleep(time.Millisecond * time.Duration(numChannel))
}
