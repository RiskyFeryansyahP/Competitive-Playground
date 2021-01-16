package main

import "fmt"

func main() {
	fmt.Println("Median of Two Sorted Arrays")

	m := []int{1}
	n := []int{}

	ans := findMedianSortedArrays(m, n)

	fmt.Println("ans", ans)
}

/* func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    ptrA := 0
    ptrB := 0

    // create new slice
    result := make([]float64, 0)

    for ptrA < len(nums1) {
        if ptrB >= len(nums2) {
            break
        }

        if nums1[ptrA] < nums2[ptrB] {
            result = append(result, float64(nums1[ptrA]))
            ptrA++
        } else {
            result = append(result, float64(nums2[ptrB]))
            ptrB++
        }
    }

    if ptrA < len(nums1) {
        for i := ptrA; i < len(nums1); i++ {
            result = append(result, float64(nums1[i]))
        }
    } else if ptrB < len(nums2) {
        for i := ptrB; i < len(nums2); i++ {
            result = append(result, float64(nums2[i]))
        }
    }

    middle := len(result) / 2
    var answer float64

    if len(result) % 2 == 0 {
        answer = (result[middle] + result[middle-1]) / 2
    } else {
        answer = result[middle]
    }

    return answer
} */

// My Answer 1
/* func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    ptrA := 0
    ptrB := 0

    // create new slice
    result := make([]float64, 0)

    for ptrA < len(nums1) {
        if ptrB >= len(nums2) {
            break
        }

        if nums1[ptrA] < nums2[ptrB] {
            result = append(result, float64(nums1[ptrA]))
            ptrA++
        } else {
            result = append(result, float64(nums2[ptrB]))
            ptrB++
        }
    }

    if ptrA < len(nums1) {
        for i := ptrA; i < len(nums1); i++ {
            result = append(result, float64(nums1[i]))
        }
    } else if ptrB < len(nums2) {
        for i := ptrB; i < len(nums2); i++ {
            result = append(result, float64(nums2[i]))
        }
    }

    middle := len(result) / 2
    var answer float64

    if len(result) % 2 == 0 {
        answer = (result[middle] + result[middle-1]) / 2
    } else {
        answer = result[middle]
    }

    return answer
} */

// My answer 2
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	ptrA := 0
	ptrB := 0

	// create new slice
	result := make([]int, 0)

	for ptrA < len(nums1) {
		if ptrB >= len(nums2) {
			break
		}

		if nums1[ptrA] < nums2[ptrB] {
			result = append(result, nums1[ptrA])
			ptrA++
		} else {
			result = append(result, nums2[ptrB])
			ptrB++
		}
	}

	if ptrA < len(nums1) {
		result = append(result, nums1[ptrA:]...)
	} else if ptrB < len(nums2) {
		result = append(result, nums2[ptrB:]...)
	}

	fmt.Println(result)

	middle := len(result) / 2
	var answer float64

	if len(result)%2 == 0 {
		answer = (float64(result[middle]) + float64(result[middle-1])) / 2
	} else {
		answer = float64(result[middle])
	}

	return answer
}

// Big O(m+n) - 2
/* func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    m, n:= len(nums1), len(nums2)
    nums := make([]int, m+n)
    i, j, k := 0, 0, 0
    for i < m || j < n {
        if i < m && j < n {
            if nums1[i] < nums2[j] {
                nums[k] = nums1[i]
                i++
            } else {
                nums[k] = nums2[j]
                j++
            }
        } else if i < m {
            nums[k] = nums1[i]
            i++
        } else {
            nums[k] = nums2[j]
            j++
        }
        k++
    }

    return float64((nums[(m + n)>>1] + nums[((m + n - 1)>>1)])) / 2.0
} */
