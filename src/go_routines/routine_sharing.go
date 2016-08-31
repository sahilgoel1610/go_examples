package main

import "fmt"

func f(number int, from string){
	for  i := 0; i < number ; i++{
		fmt.Println(from,":",i)
		number -= 1
	}

}

func main(){
	var a int = 6
	f(a, "normal")
	go f(a, "routine")
	go f(a, "routing")
  	var input string
    fmt.Scanln(&input)

}