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

	for i := 0; i < len(nums); i++ {
		if len(nums) == 1 {
			break
		}

		if nums[index] == nums[i] {
			continue
		}

		nums[length] = nums[i]
		index = i
		length++
	}

	return length
}
