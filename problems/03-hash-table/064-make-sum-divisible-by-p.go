//go:build ignore

// Problem 064: Make Sum Divisible by P (LeetCode #1590)
// Difficulty: Medium
// Categories: Hash Table, Prefix Sum, Math
//
// Description:
//   Given nums and integer p, remove the smallest contiguous subarray such
//   that the remaining sum is divisible by p. Return its length, or -1 if
//   not possible. (You can't remove the entire array.)
//
// Examples:
//   nums=[3,1,4,2], p=6     -> 1   (remove [4])
//   nums=[6,3,5,2], p=9     -> 2   (remove [5,2] or [6,3])
//   nums=[1,2,3], p=3       -> 0
//
// Approach: Prefix Sum Mod p, Find Closest Earlier Index
//   Let S = sum(nums) mod p. If S == 0, return 0.
//   We need a subarray with sum ≡ S (mod p). For each prefix index r, want
//   the latest index l with P[l] ≡ (P[r] - S) mod p. Track most-recent index
//   per residue.
//
// Complexity:
//   Time:  O(n)
//   Space: O(p) bounded by O(n)

package main

import "fmt"

func minSubarray(nums []int, p int) int {
	total := 0
	for _, v := range nums {
		total = (total + v) % p
	}
	target := total
	if target == 0 {
		return 0
	}
	last := map[int]int{0: -1}
	prefix := 0
	best := len(nums)
	for i, v := range nums {
		prefix = (prefix + v) % p
		want := (prefix - target + p) % p
		if idx, ok := last[want]; ok {
			if i-idx < best {
				best = i - idx
			}
		}
		last[prefix] = i
	}
	if best == len(nums) {
		return -1
	}
	return best
}

func main() {
	fmt.Println(minSubarray([]int{3, 1, 4, 2}, 6)) // 1
	fmt.Println(minSubarray([]int{6, 3, 5, 2}, 9)) // 2
	fmt.Println(minSubarray([]int{1, 2, 3}, 3))    // 0
}
