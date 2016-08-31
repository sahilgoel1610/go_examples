package main
import "time"
import "fmt"

func main() {
	
	jobs := make(chan int)
	// result := make(chan int)

	// for w := 1; w <= 3; w++{
	// 	go workers(w, jobs, result)
	// }
	// func_rand(jobs)
	go func_rand(jobs)
	func_rand_reader(jobs)

	time.Sleep(10*time.Second)


	// close(jobs)



}

func workers(id int, jobs <- chan int, result chan <- int ) {
	for j := range jobs{
		fmt.Println("Worker ", id, "performing job",j)
		time.Sleep(1*time.Second)
		result <- j*2
	}
}

func func_rand(jobs chan<- int) {
	
		for j := 1; j <= 9 ; j++{
			jobs <- j
		}
}


func func_rand_reader(jobs <- chan int) {
	for a := 1; a <= 9; a++ {
		fmt.Println( <- jobs)
    }
}