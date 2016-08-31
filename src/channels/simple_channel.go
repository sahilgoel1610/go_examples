package main

import "fmt"

func main() {
	var a int = 5
	go func () {
		fmt.Println(a)
	}()

	fmt.Println(a)
}