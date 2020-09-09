package main

import "fmt"

func main() {
	fmt.Println("Third Maximum Number")

	dataTest := []struct {
		nums   []int
		output int
	}{
		{nums: []int{2, 2, 3, 1}, output: 1},
		{nums: []int{3, 3, 4, 3, 4, 3, 0, 3, 3}, output: 0},
	}

	for _, v := range dataTest {
		fmt.Println("Input:", v.nums)

		output := thirdMax(v.nums)

		fmt.Println("Output:", output)

		fmt.Println("Expected Output:", v.output)

		fmt.Println("==========================")
	}
}

// 3, 3, 4, 3, 4, 3, 0, 3, 3

func thirdMax(nums []int) int {
	max := -1 // why -1 because int in golang is strong typed doesnt have null
	secondMax := -1
	thirdMax := -1

	for _, num := range nums {
		if num == max || num == secondMax || num == thirdMax {
			continue
		}

		if max == -1 || num > max {
			thirdMax = secondMax
			secondMax = max
			max = num
		} else if secondMax == -1 || num > secondMax {
			thirdMax = secondMax
			secondMax = num
		} else if thirdMax == -1 || num > thirdMax {
			thirdMax = num
		}
	}

	if thirdMax == -1 {
		return max
	}

	return thirdMax
}
