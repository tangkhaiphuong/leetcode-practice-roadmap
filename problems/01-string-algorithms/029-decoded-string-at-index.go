//go:build ignore

// Problem 029: Decoded String at Index (LeetCode #880)
// Difficulty: Medium
// Categories: String, Stack, Math
//
// Description:
//   An encoded string s is constructed by reading chars left to right: a
//   letter appends itself; a digit d repeats the current decoded string d
//   times. Given s and k, return the k-th letter (1-indexed) of the fully
//   decoded string.
//
// Examples:
//   "leet2code3", k=10 -> "o"  (decoded: "leetleetcodeleetleetcodeleetleetcode")
//   "ha22",        k=5 -> "h"
//   "a2345678999999999999999", k=1 -> "a"
//
// Approach: Reverse Walk Without Materializing
//   Forward, compute the total size after each char (caps at k). Then walk
//   backwards: if current char is a digit d, we know decoded_len/d is the
//   length of one repetition; reduce k = k % (decoded_len/d) and shrink
//   decoded_len. If letter and k==0 (or k==decoded_len), that's the answer.
//
//   This avoids exponential blowup when sizes can reach 2*10^9.
//
// Complexity:
//   Time:  O(n)
//   Space: O(1)

package main

import "fmt"

func decodeAtIndex(s string, k int) string {
	size := int64(0)
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			size *= int64(s[i] - '0')
		} else {
			size++
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		k = int(int64(k) % size)
		if k == 0 && (s[i] < '0' || s[i] > '9') {
			return string(s[i])
		}
		if s[i] >= '0' && s[i] <= '9' {
			size /= int64(s[i] - '0')
		} else {
			size--
		}
	}
	return ""
}

func main() {
	fmt.Println(decodeAtIndex("leet2code3", 10))
	fmt.Println(decodeAtIndex("ha22", 5))
	fmt.Println(decodeAtIndex("a2345678999999999999999", 1))
}
