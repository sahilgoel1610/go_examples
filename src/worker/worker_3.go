package main
import "time"
import "fmt"

func main() {
	
	jobs := make(chan int,10)
	result := make(chan int)

	// initilizing a pool of readers. 3 readrs
	for w := 1; w <= 3; w++{
		go workers(w, jobs, result)
	}
	// go func_rand_reader(jobs)
	// go result_reader(result)
	insert(jobs)

	fmt.Println("Closing jobs") 
	close(jobs)

	// time.Sleep(10*time.Second)

	    // Finally we collect all the results of the work.
	result_reader(result)

}

func workers(id int, jobs <- chan int, result chan <- int ) {
	for j := range jobs{
		fmt.Println("Worker ", id, "performing job",j)
		time.Sleep(1*time.Second)
		result <- j*2
	}
}

func insert(jobs chan<- int) {
	
		for j := 1; j <= 9 ; j++{
			jobs <- j
			fmt.Println(j,"job inserted")
		}
}


func func_rand_reader(jobs <- chan int) {
	for a := 1; a <= 9; a++ {
		fmt.Println( <- jobs)
    }
}

func result_reader(result <-chan int) {
	 for a := 1; a <= 9; a++ {
        <-result
    }
}