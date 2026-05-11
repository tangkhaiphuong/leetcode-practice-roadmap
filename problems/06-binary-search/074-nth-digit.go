//go:build ignore

// Problem 074: Nth Digit (LeetCode #400)
// Difficulty: Medium
// Categories: Binary Search, Math
//
// Description:
//   Given an integer n, return the n-th digit of the infinite sequence
//   "123456789101112...".
//
// Examples:
//   n=3   -> 3
//   n=11  -> 0   ("...91011": 9,1,0,1,1 -> position 11 is '0')
//
// Approach: Bracket by digit-length
//   Numbers with k digits range over [10^(k-1), 10^k - 1], count =
//   9*10^(k-1), each contributing k digits, so total digits in this band =
//   9*10^(k-1)*k.
//   Subtract bands until n falls within one. Then number = 10^(k-1) +
//   (n-1)/k, digit = string(number)[(n-1)%k].
//
// Complexity:
//   Time:  O(log n)
//   Space: O(1)

package main

import (
	"fmt"
	"strconv"
)

func findNthDigit(n int) int {
	k, base, count := 1, 1, 9
	for n > k*count {
		n -= k * count
		k++
		base *= 10
		count *= 10
	}
	num := base + (n-1)/k
	digit := strconv.Itoa(num)[(n-1)%k]
	return int(digit - '0')
}

func main() {
	fmt.Println(findNthDigit(3))  // 3
	fmt.Println(findNthDigit(11)) // 0
	fmt.Println(findNthDigit(1000000000))
}
