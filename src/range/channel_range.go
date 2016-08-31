package main
import "fmt"
import "time"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	go close_channel(queue)

	for elem := range queue{
		fmt.Println(elem)
	}


}


func close_channel(channel chan <- string) {
	time.Sleep(4*time.Second)
	channel <- "three"
	close(channel)
}