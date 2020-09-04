package main

import "fmt"

func main() {
	fmt.Println("Replace Elements with Greatest Element on Right Side")

	dataTest := []struct {
		arr []int
	}{
		{arr: []int{17, 18, 5, 4, 6, 1}},
		{arr: []int{10, 5, 4, 7, 9, 4, 7, 5, 8}},
		{arr: []int{5, 4, 7, 9, 1, 4, 5, 3, 6, 6, 1}},
		{arr: []int{1, 2}},
	}

	for _, v := range dataTest {
		fmt.Printf("Input: %+v \n", v.arr)

		replaceElements(v.arr)

		fmt.Printf("Output: %+v \n", v.arr)
	}
}

func replaceElements(arr []int) []int {
	bigNumber := arr[len(arr)-1]
	var tmpBigNumber int

	for i := len(arr) - 1; i >= 0; i-- {
		if i == len(arr)-1 {
			arr[i] = -1
			continue
		}

		if arr[i] > bigNumber {
			tmpBigNumber = arr[i]
			arr[i] = bigNumber
			bigNumber = tmpBigNumber
		} else {
			arr[i] = bigNumber
		}
	}

	return arr
}

func replaceElements2(arr []int) []int {
	max := -1
	for i := len(arr) - 1; i >= 0; i-- {
		tmp := arr[i]
		arr[i] = max
		if tmp > max {
			max = tmp
		}
	}
	return arr
}
