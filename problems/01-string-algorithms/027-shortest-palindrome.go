//go:build ignore

// Problem 027: Shortest Palindrome (LeetCode #214)
// Difficulty: Hard
// Categories: String, Rolling Hash, KMP
//
// Description:
//   Given s, return the shortest palindrome you can find by adding chars in
//   front of s.
//
// Examples:
//   "aacecaaa" -> "aaacecaaa"
//   "abcd"     -> "dcbabcd"
//   ""         -> ""
//
// Approach: KMP Failure Function on s + "#" + reverse(s)
//   We want the longest prefix of s that is itself a palindrome. Equivalently,
//   the longest prefix of s that equals a suffix of reverse(s).
//   Build pattern = s + "#" + reverse(s) and compute the KMP failure array;
//   fail[len-1] is the answer length L. Then prepend reverse(s[L:]) to s.
//
//   The "#" separator prevents the prefix from overlapping the boundary.
//
// Complexity:
//   Time:  O(n)
//   Space: O(n)

package main

import "fmt"

func shortestPalindrome(s string) string {
	if s == "" {
		return s
	}
	rev := reverse(s)
	pat := s + "#" + rev
	fail := make([]int, len(pat))
	for i := 1; i < len(pat); i++ {
		j := fail[i-1]
		for j > 0 && pat[i] != pat[j] {
			j = fail[j-1]
		}
		if pat[i] == pat[j] {
			j++
		}
		fail[i] = j
	}
	L := fail[len(pat)-1] // longest palindromic prefix length
	return rev[:len(s)-L] + s
}

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func main() {
	for _, s := range []string{"aacecaaa", "abcd", "", "a", "aa"} {
		fmt.Printf("shortestPalindrome(%q) = %q\n", s, shortestPalindrome(s))
	}
}
