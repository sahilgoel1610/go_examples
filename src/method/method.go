package main

import "fmt"

type rect struct {
	width, height int
}


func (r rect) area() int {
	return r.width * r.height
}

type square struct {
	side int
}


func (s square) area() int {
	return s.side * s.side
}


func main() {
	
	r := rect{width: 4, height: 5}
	fmt.Println("area", r.area())
	s := square{side: 10}
	fmt.Println("area", s.area())	
}
