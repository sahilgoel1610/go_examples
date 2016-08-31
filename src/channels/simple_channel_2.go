package main

import "fmt"

func main() {
	messages := make(chan int)

	go insert_10(messages)

	for i:= 0; i < 10; i++{
		msg1 := <- messages
		fmt.Println(msg1)
	}
}

func insert_10(channel chan <- int) {
	for i := 0; i < 10; i++{
		channel <- i
	}
}