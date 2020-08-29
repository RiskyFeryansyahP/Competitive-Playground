package main

func main() {
	duplicateZeros([]int{0, 4, 1, 0, 0, 8, 0, 0, 3})
}

func duplicateZeros(arr []int) {
	skipNext := false
	for i, num := range arr {
		if skipNext {
			// We are looking at a zero we just added. Skip it!
			skipNext = false
			continue
		}
		if num == 0 {
			// Add extra 0 to slice
			copy(arr[i+1:], arr[i:])
			// Set a flag to ignore the zero we just added on the next iteration
			skipNext = true
		}
	}
}
