package main

import "fmt"
import "time"

func main() {
	messages := make(chan int)

	go insert_10(messages)

	for i:= 0; i < 12; i++{
		msg1 := <- messages
		time.Sleep(2 * time.Second)		
		fmt.Println(msg1)
	}
}

func insert_10(channel chan <- int) {
	for i := 0; i < 10; i++{
		channel <- i
		fmt.Println("inside_routine",i)
	}
}