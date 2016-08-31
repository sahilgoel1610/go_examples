package main

import "fmt"
import "time"

func main() {
	
	fmt.Println("Delayed insertuin into the channel and nested forking of threads")

	messages := make(chan int,10)

	go insert_10(messages)

	fmt.Println("Should print inside routine without waiting for rece")
	

	for i:= 0; i < 10; i++{
		msg1 := <- messages
		time.Sleep(1 * time.Second)		
		fmt.Println(msg1)
	}
}

func insert_10(channel chan <- int) {
	for i := 0; i < 10; i++{
		fmt.Println("inside routine 1",i)
		go insert(i,channel)
	}
}


func insert (num int, channel chan <- int){
	if num%2 == 0{
		time.Sleep(15*time.Second)
	}
	channel <- num
}