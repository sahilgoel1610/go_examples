package main

import "time"
import "fmt"
import "runtime"

func main() {
		
	runtime.GOMAXPROCS(1)
	
	go func(){
		for i := 0 ; i < 100 ; i++{
			fmt.Println("Hello")
			time.Sleep(100*time.Millisecond)
		}
	} ()

	go func () {
		for i := 0; i < 100 ; i++{
			fmt.Println("GO")
			time.Sleep(100*time.Millisecond)
		}	
	}()

	time.Sleep(20*time.Second)

}