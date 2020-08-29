package main

import "log"

func main() {
	log.Println("Check if N and Its Double Exist!")

	arr := []int{0, 1}

	result := checkIfExist(arr)

	log.Println("Result:", result)
}

func checkIfExist(arr []int) bool {
	output := false

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i == j && arr[j] == 0 {
				continue
			}

			if arr[j]*2 == arr[i] {
				output = true
				break
			}
		}
	}

	return output
}
