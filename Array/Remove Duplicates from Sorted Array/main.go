package main

import "log"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 1}

	length := removeDuplicates(nums)

	for i := 0; i < length; i++ {
		log.Printf("nums %+v = %+v", i, nums[i])
	}
}

func removeDuplicates(nums []int) int {
	length := 1

	index := 0

	if len(nums) == 1 {
		return length
	}

	for i := 0; i < len(nums); i++ {
		if nums[index] == nums[i] {
			continue
		}

		nums[length] = nums[i]
		index = i
		length++
	}

	return length
}
