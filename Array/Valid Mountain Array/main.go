package main

import "log"

func main() {
	A := []int{14, 82, 89, 84, 79, 70, 70, 68, 67, 66, 63, 60, 58, 54, 44, 43, 32, 28, 26, 25, 22, 15, 13, 12, 10, 8, 7, 5, 4, 3}
	B := []int{0, 8, 10, 2, 1, 4, 1}
	C := []int{8, 9, 10}
	D := []int{3, 2, 1}
	E := []int{8}

	statusTestA := validMountainArray(A)
	statusTestB := validMountainArray(B)
	statusTestC := validMountainArray(C)
	statusTestD := validMountainArray(D)
	statusTestE := validMountainArray(E)

	log.Printf("Status Test A is: %v, actual is : %v \n", statusTestA, true)
	log.Printf("Status Test B is: %v, actual is : %v \n", statusTestB, false)
	log.Printf("Status Test C is: %v, actual is : %v \n", statusTestC, false)
	log.Printf("Status Test D is: %v, actual is : %v \n", statusTestD, false)
	log.Printf("Status Test E is: %v, actual is : %v \n", statusTestE, false)
}

// Good for readibility
/* func validMountainArray(A []int) bool {
	status := false

	if len(A) < 3 || A[0] > A[1] {
		return status
	}

	for i := 0; i < len(A); i++ {
		if i == len(A)-1 {
			break
		}

		if i < len(A)-1 && A[i] > A[i+1] && A[i-1] < A[i] {
			status = true
			continue
		}

		if i > 0 && A[i] <= A[i+1] && A[i-1] > A[i] {
			status = false
			break
		}
	}

	return status
} */

// Better solution
func validMountainArray(A []int) bool {
	lessThan := false
	greaterThan := false
	status := false

	if len(A) < 3 {
		return status
	}

	for i := 0; i < len(A); i++ {
		if i == len(A)-1 {
			break
		}

		if i < len(A)-1 && A[i] < A[i+1] && greaterThan == false {
			lessThan = true
			continue
		}

		if i < len(A)-1 && A[i] > A[i+1] && lessThan == true {
			greaterThan = true
			status = true
			continue
		}

		lessThan = false
		greaterThan = false
		status = false
		break
	}

	return status
}

// 10 9 8
// func validMountainArray(A []int) bool {
// 	status := false

// 	if len(A) < 3 {
// 		return status
// 	}

// 	for i := 0; i < len(A); i++ {
// 		if i < len(A)-1 && A[i] < A[i+1] {
// 			status = false
// 			continue
// 		}

// 		if i < len(A)-1 && A[i] > A[i+1] && status == true {
// 			status = false
// 			continue
// 		}

// 		status = true
// 	}

// 	return status
// }
