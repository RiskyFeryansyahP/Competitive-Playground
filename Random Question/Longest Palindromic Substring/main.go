package main

import "fmt"

func main() {
	testCase := []struct {
		input       string
		expectation string
	}{
		{input: "babad", expectation: "bab"},
		{input: "a", expectation: "a"},
		{input: "cbbd", expectation: "bb"},
		{input: "ac", expectation: "a"},
	}

	for _, test := range testCase {
		output := longestPalindrome(test.input)

		fmt.Printf("input: %s, expectation: %s, output: %s \n", test.input, test.expectation, output)
	}
}

func longestPalindrome(s string) string {
	right, left, maxLen := 0, 0, 0

	for i := 0; i < len(s); i++ {
		// check for odd len
		oddLen := expandMid(s, i, i)
		// check for event len
		evenLen := expandMid(s, i, i+1)

		curLen := max(oddLen, evenLen)
		if curLen > maxLen {
			left = i - (curLen-1)/2
			right = i + curLen/2
			maxLen = curLen
		}
	}

	return s[left : right+1]

}

func expandMid(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left = left - 1
		right = right + 1
	}

	return right - left - 1
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
