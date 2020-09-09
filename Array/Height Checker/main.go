package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Height Checker")

	dataTest := []struct {
		heights []int
		output  int
	}{
		{heights: []int{1, 1, 4, 2, 1, 3}, output: 3},
		{heights: []int{5, 1, 2, 3, 4}, output: 5},
		{heights: []int{1, 2, 3, 4, 5}, output: 0},
	}

	for _, v := range dataTest {
		fmt.Println("Input:", v.heights)

		output := heightChecker(v.heights)

		fmt.Println("Output:", output)

		fmt.Println("Expected Output:", v.output)

		fmt.Println("==========================")
	}
}

func heightChecker(heights []int) int {
	sortedHeights := make([]int, len(heights))
	wrongPosition := 0

	copy(sortedHeights, heights)

	sort.Ints(sortedHeights)

	for i := 0; i < len(heights); i++ {
		if heights[i] != sortedHeights[i] {
			wrongPosition++
		}
	}

	return wrongPosition
}
