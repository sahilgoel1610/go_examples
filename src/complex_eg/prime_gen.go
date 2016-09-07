package "main"

import "fmt"


func main() {
	
	for {

	}
}

func check_prime(num int,filt <-chan int ) {
	for{
		i := <- filt
		if i%num != 0
	}
}