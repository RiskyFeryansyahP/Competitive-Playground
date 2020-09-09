package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Find All Numbers Disappeared in an Array Solution")

	dataTest := []struct {
		nums   []int
		output []int
	}{
		{nums: []int{4, 3, 2, 7, 8, 2, 3, 1}, output: []int{5, 6}},
		{nums: []int{2, 5, 3, 1, 1}, output: []int{4}},
	}

	for _, v := range dataTest {
		fmt.Println("Input:", v.nums)

		output := findDisappearedNumbers(v.nums)

		fmt.Println("Output:", output)

		fmt.Println("Expected output:", v.output)

		fmt.Println("==============================")
	}
}

func findDisappearedNumbers(nums []int) []int {
	result := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		currentNumber := int(math.Abs(float64(nums[i])) - 1)

		nums[currentNumber] = int(math.Abs(float64(nums[currentNumber])) * -1)
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}

	return result
}
