package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}

func sortedSquares(A []int) []int {
	output := make([]int, len(A))

	for k, v := range A {
		output[k] = v * v
	}

	sort.Ints(A)

	fmt.Println("Value A", A)

	sort.Ints(output)

	return output
}

/*
* example 2
func sortedSquares(arr []int) []int {
	l, r := 0, len(arr)-1
	res := make([]int, len(arr))
	resp := len(arr) - 1
	for l <= r {
		if sqrt(arr[l]) > sqrt(arr[r]) {
			res[resp] = sqrt(arr[l])
			resp--
			l++
		} else {
			res[resp] = sqrt(arr[r])
			resp--
			r--
		}
	}
	return res
}
*/
