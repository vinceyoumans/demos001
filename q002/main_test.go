package q002

import (
	"reflect"
	"testing"
)

func TestSwapInts01(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{3, 2, 1}

	result := SwapInts01(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestSwapInts02(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expected := []int{6, 5, 4, 3, 2, 1}

	result := SwapInts02(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestSwapInts03(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{3, 2, 1}

	result := SwapInts03(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
