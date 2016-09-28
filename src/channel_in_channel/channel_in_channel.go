package main

import "time"
import "fmt"



func doStuff(t time.Duration, ch <-chan chan time.Duration) {
	ac := <- ch
	time.Sleep(t)
	ac <- t	
}

func main() {

	sendCh := make(chan chan time.Duration)
	
	for i := 0; i < 10; i++{
		go doStuff(time.Duration(i+1) * time.Second, sendCh)
	}

	recvCh := make(chan time.Duration)

	for i:= 0; i < 10 ; i++{
		sendCh <- recvCh
	}

	fmt.Println("hey")

	for i := 0; i < 10 ; i++{
		ack := <- recvCh
		fmt.Println("Ack received of time", ack)
	}



}