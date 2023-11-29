package q004

import (
	"fmt"
	"sync"
	"time"
)

//	I took this to mean how to spawn a routine.
//
// The term "spawned" confused me, because
// it could mean a lot of things depending on context and language.
// I assumed it was to "spawn" a goroutine,

func DemoGR001() {
	const numJobs = 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Start three worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs to the workers
	go func() {
		for i := 1; i <= numJobs; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// Wait for all workers to finish
	go func() {
		wg.Wait() // Blocks until the WaitGroup counter goes to zero
		close(results)
	}()

	// Collect results from the workers
	for result := range results {
		fmt.Printf("Received result: %d\n", result)
	}

	fmt.Println("All jobs have been completed.")
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrements the counter when the goroutine completes

	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // Simulating some work
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2
	}
}
