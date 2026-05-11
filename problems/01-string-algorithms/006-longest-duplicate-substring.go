//go:build ignore

// Problem 006: Longest Duplicate Substring (LeetCode #1044)
// Difficulty: Hard
// Categories: String, Binary Search, Sliding Window, Rolling Hash
//
// Description:
//   Given a string s, find the longest duplicated substring (occurs at least
//   twice; the occurrences may overlap). Return any such substring; "" if
//   none exists.
//
// Examples:
//   "banana"     -> "ana"
//   "abcd"       -> ""
//   "aaaaa"      -> "aaaa"
//
// Constraints:
//   2 <= s.length <= 3*10^4. s consists of lowercase English letters.
//
// Approach: Binary Search on Length + Rabin-Karp Rolling Hash
//   Property: if a duplicate of length L exists, so does one of length L-1.
//   So binary-search the answer length over [1, n-1].
//   For a fixed L, use rolling hash to check whether any substring of length
//   L appears twice. Use mod 2^61 - 1 (Mersenne prime) or a big prime to keep
//   collisions rare; verify hits by comparing actual substrings.
//
// Complexity:
//   Time:  O(n log n) average (n per hash scan, log n binary-search rounds)
//   Space: O(n)

package main

import "fmt"

const mod uint64 = (1 << 61) - 1
const base uint64 = 131

// search returns the start index of any length-L duplicate, or -1.
func search(s string, L int) int {
	if L == 0 {
		return -1
	}
	// power = base^L mod
	var power uint64 = 1
	for i := 0; i < L; i++ {
		power = mulMod(power, base)
	}
	var h uint64
	for i := 0; i < L; i++ {
		h = addMod(mulMod(h, base), uint64(s[i]))
	}
	seen := map[uint64][]int{h: {0}}
	for i := 1; i+L <= len(s); i++ {
		// remove s[i-1], add s[i+L-1]
		h = addMod(mulMod(h, base), uint64(s[i+L-1]))
		// subtract s[i-1] * base^L
		h = subMod(h, mulMod(uint64(s[i-1]), power))
		if starts, ok := seen[h]; ok {
			for _, st := range starts {
				if s[st:st+L] == s[i:i+L] {
					return st
				}
			}
			seen[h] = append(seen[h], i)
		} else {
			seen[h] = []int{i}
		}
	}
	return -1
}

func longestDupSubstring(s string) string {
	lo, hi := 1, len(s)-1
	bestStart, bestLen := 0, 0
	for lo <= hi {
		mid := (lo + hi) / 2
		if idx := search(s, mid); idx != -1 {
			bestStart, bestLen = idx, mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return s[bestStart : bestStart+bestLen]
}

// 64-bit modular ops using __int128-like splits for mod 2^61-1.
func mulMod(a, b uint64) uint64 {
	// Use big-int style split to avoid overflow.
	const mask = (uint64(1) << 32) - 1
	aHi, aLo := a>>32, a&mask
	bHi, bLo := b>>32, b&mask
	hi := aHi * bHi
	lo := aLo * bLo
	mid := aHi*bLo + aLo*bHi
	// (hi << 64) + (mid << 32) + lo, all mod (2^61-1)
	// 2^64 mod (2^61-1) == 8
	// 2^32 mod (2^61-1) == 2^32
	r := addMod(mulRaw(hi, 8), addMod(mulRaw(mid, 1<<32), lo%mod))
	return r
}
func mulRaw(a, b uint64) uint64 {
	// b is small (<= 2^32), so direct multiply mod
	res := uint64(0)
	a %= mod
	for b > 0 {
		if b&1 == 1 {
			res = addMod(res, a)
		}
		a = addMod(a, a)
		b >>= 1
	}
	return res
}
func addMod(a, b uint64) uint64 {
	r := a + b
	if r >= mod {
		r -= mod
	}
	return r
}
func subMod(a, b uint64) uint64 { return addMod(a, mod-b%mod) }

func main() {
	for _, s := range []string{"banana", "abcd", "aaaaa", "abcabcabc"} {
		fmt.Printf("longestDup(%q) = %q\n", s, longestDupSubstring(s))
	}
}
