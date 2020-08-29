package main

import "log"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	merge(nums1, 3, nums2, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	first := m - 1
	second := n - 1

	for i := len(nums1) - 1; i >= 0; i-- {
		if n < 0 {
			break
		}

		if second < 0 {
			break
		}

		if first >= 0 && nums1[first] > nums2[second] {
			nums1[i] = nums1[first]
			first--
		} else {
			nums1[i] = nums2[second]
			second--
		}
	}

	log.Println(nums1)
}
