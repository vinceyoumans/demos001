package q003

import (
	"sync"
	"testing"
)

func TestProducerConsumer(t *testing.T) {
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

	// No need to assert anything. If there are no panics, the test is considered successful.
}

func TestDemoGR001(t *testing.T) {
	DemoGR001()
	// No need to assert anything. If there are no panics, the test is considered successful.
}
