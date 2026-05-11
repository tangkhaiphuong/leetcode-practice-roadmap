//go:build ignore

// Problem 077: Pow(x, n) (LeetCode #50)
// Difficulty: Medium
// Categories: Binary Search / Recursion, Math
//
// Description:
//   Implement pow(x, n) — x raised to the integer power n.
//
// Examples:
//   x=2.0, n=10  -> 1024
//   x=2.1, n=3   -> 9.261
//   x=2.0, n=-2  -> 0.25
//
// Constraints:
//   -100.0 < x < 100.0;  -2^31 <= n <= 2^31 - 1.
//
// Approach: Iterative Fast Exponentiation
//   Convert negative n to positive by inverting x. Walk bits of n; result
//   *= base when bit is 1; square base each step.
//
//   Note: n = -2^31 in two's complement can't be negated as int32. Use int64
//   or uint64 for the magnitude.
//
// Complexity:
//   Time:  O(log |n|)
//   Space: O(1)

package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	res := 1.0
	for n > 0 {
		if n&1 == 1 {
			res *= x
		}
		x *= x
		n >>= 1
	}
	return res
}

func main() {
	fmt.Println(myPow(2.0, 10))  // 1024
	fmt.Println(myPow(2.1, 3))   // 9.261
	fmt.Println(myPow(2.0, -2))  // 0.25
}
