package main
import "time"
import "fmt"

func main() {
	
	jobs := make(chan int)
	result := make(chan int)

	for w := 1; w <= 3; w++{
		go workers(w, jobs, result)
	}

	for j := 1; j <= 9 ; j++{
		jobs <- j
	}


	close(jobs)

    for a := 1; a <= 9; a++ {
        <-result
    }

}

func workers(id int, jobs <- chan int, result chan <- int ) {
	for j := range jobs{
		fmt.Println("Worker ", id, "performing job",j)
		time.Sleep(1*time.Second)
		result <- j*2
	}
}