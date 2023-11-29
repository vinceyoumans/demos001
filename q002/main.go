package q002

import (
	"fmt"
	"reflect"
)

//  I missunderstood the question.
// It appears you meant swap a slice of Integers

func SwapInts01(i []int) []int {
	fmt.Println("========  SwapInt01 ")
	//  This assumes there are only 3 Items in the Slice,
	// Which was not mentioned by Bhanu.  This confused me.
	// But assuming 3 Items in the slice
	fmt.Printf("Slice001 = %v\n", i)
	swap := reflect.Swapper(i)

	swap(0, 2)
	fmt.Printf("Slice002 = %v\n", i)

	return i

}

func SwapInts02(i []int) []int {
	fmt.Println("========  SwapInt02 ")
	fmt.Printf("Slice001 = %v\n", i)
	ct := len(i) / 2

	swap := reflect.Swapper(i)

	for x := 0; x < ct; x++ {
		ctx := len(i) - 1 - x
		swap(x, ctx)
	}

	fmt.Printf("Slice001 = %v\n", i)
	return i
}

// this is the wrong way because creating a new Slice of int.
func SwapInts03(i []int) []int {
	fmt.Println("========  SwapInt03 ")
	fmt.Println("Q2 - SwapInts03")
	fmt.Printf("Slice001 = %v\n", i)

	ln := len(i)
	ii := []int{}
	for x := ln - 1; x >= 0; x-- {
		ii = append(ii, i[x])
	}
	fmt.Printf("Slice002 = %v\n", ii)
	return ii
}
