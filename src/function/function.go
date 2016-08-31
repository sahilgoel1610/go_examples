package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}


func main(){
	a := 1
	b := 2
	res := plus(a,b)
	fmt.Println("Result is ",res)
}