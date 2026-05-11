//go:build ignore

// Problem 030: Split Array With Same Average (LeetCode #805)
// Difficulty: Hard
// Categories: String/Array, DP, Greedy, Bitmask, Meet-in-the-Middle
//
// Description:
//   Given an integer array nums, partition into two non-empty subsets A, B
//   such that average(A) == average(B). Return true iff possible.
//
// Examples:
//   [1,2,3,4,5,6,7,8] -> true   ([1,4,5,8] and [2,3,6,7] both avg 4.5)
//   [3,1]             -> false
//
// Constraints:
//   1 <= n <= 30, 0 <= nums[i] <= 10^4.
//
// Approach: Meet in the Middle (n up to 30)
//   We need a subset of size k (1..n-1) summing to k * S / n where S = sum.
//   Equivalently, target = k*S must be divisible by n.
//   Split nums into two halves of size ~15. Enumerate all subset sums of
//   each half, indexed by subset size. For each size k1 in left, look up if
//   right has size (k - k1) summing to (target - leftSum).
//
//   The pruning by integer divisibility (k*S % n == 0) cuts most invalid k.
//
// Complexity:
//   Time:  O(2^(n/2) * n)
//   Space: O(2^(n/2) * n)

package main

import "fmt"

func splitArraySameAverage(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	S := 0
	for _, v := range nums {
		S += v
	}
	// Find any k in [1..n/2] s.t. k*S divisible by n.
	possible := false
	for k := 1; k <= n/2; k++ {
		if k*S%n == 0 {
			possible = true
			break
		}
	}
	if !possible {
		return false
	}

	half := n / 2
	left, right := nums[:half], nums[half:]
	leftSums := subsetSumsBySize(left)   // map[size]map[sum]bool
	rightSums := subsetSumsBySize(right) // same

	// We need k items summing to k*S/n where (k*S)%n==0, k in [1..n-1].
	for k := 1; k < n; k++ {
		if k*S%n != 0 {
			continue
		}
		target := k * S / n
		// Choose k1 from left (0..min(k,len(left))), k2 = k-k1 from right.
		for k1 := 0; k1 <= k && k1 <= len(left); k1++ {
			k2 := k - k1
			if k2 < 0 || k2 > len(right) {
				continue
			}
			// Empty-on-both-sides means subset is empty — skip.
			lset, lok := leftSums[k1]
			rset, rok := rightSums[k2]
			if k1 == 0 && k2 == 0 {
				continue
			}
			if k1 == 0 && rok {
				if rset[target] {
					// Don't allow taking ALL of right when left is empty AND len(right)==k.
					if !(k1 == 0 && k2 == n) {
						return true
					}
				}
				continue
			}
			if k2 == 0 && lok {
				if lset[target] {
					if !(k2 == 0 && k1 == n) {
						return true
					}
				}
				continue
			}
			if !lok || !rok {
				continue
			}
			for ls := range lset {
				if rset[target-ls] {
					return true
				}
			}
		}
	}
	return false
}

func subsetSumsBySize(a []int) map[int]map[int]bool {
	m := map[int]map[int]bool{}
	n := len(a)
	for mask := 0; mask < (1 << n); mask++ {
		size, sum := 0, 0
		for i := 0; i < n; i++ {
			if mask>>i&1 == 1 {
				size++
				sum += a[i]
			}
		}
		if m[size] == nil {
			m[size] = map[int]bool{}
		}
		m[size][sum] = true
	}
	return m
}

func main() {
	fmt.Println(splitArraySameAverage([]int{1, 2, 3, 4, 5, 6, 7, 8})) // true
	fmt.Println(splitArraySameAverage([]int{3, 1}))                   // false
	fmt.Println(splitArraySameAverage([]int{0, 0}))                   // false (not non-empty disjoint)
}
