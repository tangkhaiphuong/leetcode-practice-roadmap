//go:build ignore

// Problem 041: Next Greater Element III (LeetCode #556)
// Difficulty: Medium
// Categories: Array (digits), Math
//
// Description:
//   Given a positive int n, return the smallest int that uses the same digits
//   and is strictly greater than n. Return -1 if it doesn't exist or doesn't
//   fit in a 32-bit signed integer.
//
// Examples:
//   n=12   -> 21
//   n=21   -> -1
//   n=12443322 -> 13222344
//
// Approach: Same as Next Permutation on Digits
//   1. Find pivot i: rightmost index where digits[i] < digits[i+1].
//   2. If none -> -1.
//   3. Find rightmost j > i with digits[j] > digits[i]. Swap.
//   4. Reverse digits[i+1:].
//   5. Parse, check int32 overflow.
//
// Complexity:
//   Time:  O(d) where d = digits in n
//   Space: O(d)

package main

import (
	"fmt"
	"math"
	"strconv"
)

func nextGreaterElement(n int) int {
	digits := []byte(strconv.Itoa(n))
	i := len(digits) - 2
	for i >= 0 && digits[i] >= digits[i+1] {
		i--
	}
	if i < 0 {
		return -1
	}
	j := len(digits) - 1
	for digits[j] <= digits[i] {
		j--
	}
	digits[i], digits[j] = digits[j], digits[i]
	for l, r := i+1, len(digits)-1; l < r; l, r = l+1, r-1 {
		digits[l], digits[r] = digits[r], digits[l]
	}
	v, _ := strconv.ParseInt(string(digits), 10, 64)
	if v > math.MaxInt32 {
		return -1
	}
	return int(v)
}

func main() {
	fmt.Println(nextGreaterElement(12))         // 21
	fmt.Println(nextGreaterElement(21))         // -1
	fmt.Println(nextGreaterElement(12443322))   // 13222344
	fmt.Println(nextGreaterElement(2147483486)) // -1 (overflow)
}
