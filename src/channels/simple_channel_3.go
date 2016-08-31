package main

import "fmt"
import "time"

func main() {
	messages := make(chan int)

	go insert_10(messages)

	msg1 := <- messages
	fmt.Println(msg1)
	time.Sleep(10 * time.Second)
}

func insert_10(channel chan <- int) {
	for i := 0; i < 10; i++{
		channel <- i
		fmt.Println("inside_routine",i)
	}
}