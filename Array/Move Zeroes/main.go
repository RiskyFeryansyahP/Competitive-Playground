package main

import "fmt"

func main() {
	fmt.Println("Move Zeroes")

	dataTest := []struct {
		nums []int
	}{
		{nums: []int{0, 1, 0, 3, 12}}, // 1 12 3 0 0
		{nums: []int{0, 1, 0, 3, 0}},
	}

	for _, v := range dataTest {
		fmt.Println("Input:", v.nums)

		moveZeroes(v.nums)

		fmt.Println("Output:", v.nums)
	}
}

func moveZeroes(nums []int) {
	zeroNumber := 0

	for _, v := range nums {
		if v != 0 {
			nums[zeroNumber] = v
			zeroNumber++
		}
	}

	for i := zeroNumber; i < len(nums); i++ {
		nums[i] = 0
	}
}
