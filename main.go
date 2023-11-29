package main

import (
	"InterviewQ/q001"
	"InterviewQ/q002"
	"InterviewQ/q003"
	"InterviewQ/q004"
	"fmt"
	"os"
)

func main() {
	fmt.Println("starting interview question 001 - Demonstrate RUNE")
	q001.DemoStringVRunesWithWriter(os.Stdout)

	// Question was how to swap numbers 321 to 123
	// This question confused me.
	// The term swap was used, which is a reflect Func
	fmt.Println("---  Q2 A  Swapper ---")
	ii := []int{3, 2, 1}
	fmt.Println("----   Q2 ", ii)
	fmt.Println(q002.SwapInts01(ii))

	// But Same question, Assuming x items in the slice of Int
	// I would probably do it this way

	fmt.Println("Q2 B...   but with x items in []int")
	fmt.Println(q002.SwapInts02([]int{7, 4, 12, 3, 2, 1}))

	fmt.Println("Q2 C...   but with x items in []int")
	fmt.Println(q002.SwapInts03([]int{7, 4, 12, 3, 2, 1}))

	// Spawn A go routine.
	// I assumed that this meant a Go Routine.
	// which is just to add "go" to the begining of the function.
	//  but it is more complicated than that.

	// version 1....   just a more complicated goroutine
	fmt.Println("q003========   starting goroutine 3 example with wg, channels")
	q003.DemoGR001()

	// version 2....   just a more complicated goroutine
	fmt.Println("q004========   starting goroutine 4 example with wg, channels")
	q004.DemoGR001()
}
