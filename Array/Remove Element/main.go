package main

import "fmt"

func main() {
	nums := []int{0, 4, 4, 0, 4, 4, 4, 0, 2}
	val := 4

	newLength := removeElement(nums, val)

	for i := 0; i < newLength; i++ {
		fmt.Println(nums[i])
	}
}

func removeElement(nums []int, val int) int {
	i := 0
	length := len(nums)

	for i < length {
		if nums[i] == val {
			nums[i] = nums[length-1]
			// reduce array size by one
			length--
		} else {
			i++
		}
	}

	return length
}

func removeElement2(nums []int, val int) int {
	// 0, 4, 4, 0, 4, 4, 4, 0, 2
	// 4

	length := 0

	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[length] = nums[j]
			length++
		}
	}

	return length
}
