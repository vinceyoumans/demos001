package q003

import (
	"fmt"
	"sync"
)

func producer(numbers chan<- int, count int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= count; i++ {
		numbers <- i
	}
	close(numbers)
}

func consumer(id int, numbers <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range numbers {
		fmt.Printf("Consumer %d received: %d\n", id, num)
	}
}

func DemoGR001() {
	const numConsumers = 3
	const numNumbers = 10

	var wg sync.WaitGroup
	numbers := make(chan int, numNumbers)

	// Start the producer
	wg.Add(1)
	go producer(numbers, numNumbers, &wg)

	// Start the consumers
	for i := 1; i <= numConsumers; i++ {
		wg.Add(1)
		go consumer(i, numbers, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All goroutines have completed.")
}
