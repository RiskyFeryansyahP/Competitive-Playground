package main

import "fmt"

func main() {
	fmt.Println("Two Pointer Technique")

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 4 // find if there exists any pair of elements such that their sum is equal to X

	startIndex := 0
	endIndex := len(array) - 1
	counter := 0

	for i := 0; i < len(array); i++ {
		if array[startIndex]+array[endIndex] == target {
			fmt.Println("num 1:", array[startIndex])
			fmt.Println("num 2:", array[endIndex])
			counter++
		}

		if array[startIndex]+array[endIndex] > target {
			endIndex--
		} else {
			startIndex++
		}
	}

	fmt.Println("the sum of the two elements that add up to the target is", counter)
}
