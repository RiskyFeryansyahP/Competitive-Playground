package main

import "fmt"

func main() {
	fmt.Println("Sort Array By Parity")

	dataTest := []struct {
		A      []int
		answer []int
	}{
		{A: []int{3, 1, 2, 4}, answer: []int{2, 4, 1, 3}},
		{A: []int{5, 8, 10, 11, 13, 27}, answer: []int{8, 10, 27, 13, 11, 5}},
	}

	for _, v := range dataTest {
		fmt.Println("Input:", v.A)

		fmt.Println("Expected ouput:", sortArrayByParity(v.A))

		fmt.Println("Answer:", v.answer)

		fmt.Println("==========")
	}
}

func sortArrayByParity(A []int) []int {
	oddsNumber := 0
	evenNumber := len(A) - 1

	newArray := make([]int, len(A))

	for _, v := range A {
		if v%2 == 0 {
			newArray[oddsNumber] = v
			oddsNumber++
		} else {
			newArray[evenNumber] = v
			evenNumber--
		}
	}

	return newArray
}
