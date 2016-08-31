package main

import "fmt"
import "time"

func main() {
	messages := make(chan int)

	go insert_10(messages)

	fmt.Println("Should wait for 5 seconds before first digits output")
	time.Sleep(5*time.Second)

	for i:= 0; i < 10; i++{
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