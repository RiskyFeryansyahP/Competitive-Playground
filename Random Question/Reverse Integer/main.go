package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Reverse Integer Problem")

	// Given a signed 32-bit integer x, return x with its digits reversed.
	// If reversing x causes the value to go outside the signed 32-bit integer range, then return 0.

	testCase := []struct {
		testcase    int
		expectation int
	}{
		{testcase: 123, expectation: 321},
		{testcase: -123, expectation: -321},
		{testcase: 120, expectation: 21},
		{testcase: 1534236469, expectation: 0},
		{testcase: 0, expectation: 0},
	}

	for _, v := range testCase {
		res := reverse2(v.testcase)

		fmt.Printf("testcase: %d, expectation: %d, output: %d \n", v.testcase, v.expectation, res)
	}
}

func reverse(x int) int {
	var isMinus bool

	if x < 0 {
		isMinus = true
		x = x * -1
	}

	s := strconv.Itoa(x)

	size := len(s) - 1

	var res string

	for i := size; i >= 0; i-- {
		res = res + string(s[i])
	}

	result, _ := strconv.Atoi(res)

	if isMinus {
		result = result * -1
	}

	if result < -2147483648 || result > 2147483648 {
		return 0
	}

	return result
}

func reverse2(x int) int {
	var pop int
	var rev int // result reverse

	for x != 0 {
		// pop the value
		pop = x % 10
		x = x / 10

		rev = rev*10 + pop
	}

	if rev > math.MaxInt32 || rev < math.MinInt32 {
		return 0
	}

	return rev
}
