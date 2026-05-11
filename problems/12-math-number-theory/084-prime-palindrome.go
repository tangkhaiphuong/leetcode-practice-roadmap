//go:build ignore

// Problem 084: Prime Palindrome (LeetCode #866)
// Difficulty: Medium
// Categories: Math, Number Theory
//
// Description:
//   Find the smallest prime that is also a palindrome and >= n.
//
// Examples:
//   n=6     -> 7
//   n=8     -> 11
//   n=13    -> 101
//
// Approach: Generate Palindromes by Their Left Half
//   Key fact: every even-length palindrome > 11 is divisible by 11, hence
//   composite. So we only need to consider odd-length palindromes (and
//   special-case 11).
//   Build palindromes from left-half h = 1,2,3,...; mirror to form an
//   odd-length palindrome of length 2*len(h)-1. Check primality with trial
//   division up to sqrt.
//
// Complexity:
//   Time:  O(palindrome_count * sqrt(P))
//   Space: O(1)

package main

import (
	"fmt"
	"strconv"
)

func primePalindrome(n int) int {
	if n <= 2 {
		return 2
	}
	if n <= 3 {
		return 3
	}
	if n <= 5 {
		return 5
	}
	if n <= 7 {
		return 7
	}
	if n <= 11 {
		return 11
	}
	for h := 1; ; h++ {
		s := strconv.Itoa(h)
		// Construct odd-length palindrome.
		full := s
		for i := len(s) - 2; i >= 0; i-- {
			full += string(s[i])
		}
		v, _ := strconv.Atoi(full)
		if v < n {
			continue
		}
		if isPrime(v) {
			return v
		}
	}
}

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	if x%2 == 0 {
		return x == 2
	}
	for d := 3; d*d <= x; d += 2 {
		if x%d == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primePalindrome(6))  // 7
	fmt.Println(primePalindrome(8))  // 11
	fmt.Println(primePalindrome(13)) // 101
}
