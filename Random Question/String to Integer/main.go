package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("String to Integer")

	res := myAtoi("44465 hello world") // expected return 44465
	fmt.Println(res)
}

func myAtoi(s string) int {
	var result int

	if len(s) < 1 {
		return 0
	}

	if s[0] == ' ' {
		for k, v := range s {
			if v != ' ' {
				s = s[k:]
				break
			}
		}
	}

	s0 := s[0]

	if s[0] == '-' || s[0] == '+' {
		s = s[1:]

		if len(s) < 1 {
			return 0
		}
	}

	for _, v := range []byte(s) {
		if v == 32 {
			break
		}

		v -= '0'

		if v > 9 {
			break
		}

		result = result*10 + int(v)

		if result > math.MaxInt32 || result < math.MinInt32 {
			break
		}
	}

	if s0 == '-' {
		result = -result
	}

	if result > math.MaxInt32 {
		return int(1<<31 - 1)
	} else if result < math.MinInt32 {
		return int(-1 << 31)
	}

	return result
}

// i should get the first chars of string to detect either '-' or '+', if '-' then should negative
// i should loop string and get char of string and wrap string in array of byte
// get value each byte and convert to int and sum with result
// result = result * 10, why we multiple by 10 because sometimes the result is > 10
