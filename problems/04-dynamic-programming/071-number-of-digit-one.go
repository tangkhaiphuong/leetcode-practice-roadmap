//go:build ignore

// Problem 071: Number of Digit One (LeetCode #233)
// Difficulty: Hard
// Categories: DP (Digit DP), Math
//
// Description:
//   Given an integer n, count the total number of digit '1' appearing in all
//   non-negative integers <= n.
//
// Examples:
//   n=13  -> 6     (1,10,11,12,13: 1+1+2+1+1)
//   n=0   -> 0
//
// Approach: Place-Value Math
//   For each digit position with place value p (1, 10, 100, ...):
//     high = n / (p*10)
//     cur  = (n / p) % 10
//     low  = n % p
//     count += high * p
//     if cur == 0:        +0
//     elif cur == 1:      + low + 1
//     else (cur >= 2):    + p
//
//   Iterate p until p > n.
//
// Complexity:
//   Time:  O(log10 n)
//   Space: O(1)

package main

import "fmt"

func countDigitOne(n int) int {
	if n <= 0 {
		return 0
	}
	count := 0
	p := 1
	for p <= n {
		high := n / (p * 10)
		cur := (n / p) % 10
		low := n % p
		count += high * p
		switch {
		case cur == 1:
			count += low + 1
		case cur > 1:
			count += p
		}
		p *= 10
	}
	return count
}

func main() {
	fmt.Println(countDigitOne(13)) // 6
	fmt.Println(countDigitOne(0))  // 0
	fmt.Println(countDigitOne(20)) // 12
	fmt.Println(countDigitOne(100))// 21
}
