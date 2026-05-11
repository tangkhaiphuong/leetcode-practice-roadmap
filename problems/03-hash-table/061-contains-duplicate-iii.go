//go:build ignore

// Problem 061: Contains Duplicate III (LeetCode #220)
// Difficulty: Hard
// Categories: Hash Table, Sliding Window, Bucket
//
// Description:
//   Given nums and integers indexDiff k, valueDiff t, return true iff there
//   exist i != j with |i - j| <= k and |nums[i] - nums[j]| <= t.
//
// Examples:
//   nums=[1,2,3,1], k=3, t=0       -> true
//   nums=[1,5,9,1,5,9], k=2, t=3   -> false
//
// Approach: Bucketing
//   Use bucket size = t+1. For nums[i], its bucket is nums[i]/(t+1) (with
//   adjustment for negatives). Two numbers in the same bucket -> diff <= t.
//   Also check the two adjacent buckets: any number there might still be
//   within t. Maintain only buckets in the last k+1 indices.
//
//   Special-case t == 0: only same-bucket counts (no adjacency math).
//
// Complexity:
//   Time:  O(n)
//   Space: O(min(n, k))

package main

import "fmt"

func containsNearbyAlmostDuplicate(nums []int, k, t int) bool {
	if t < 0 {
		return false
	}
	w := t + 1
	getBucket := func(x int) int {
		if x >= 0 {
			return x / w
		}
		return -((-x-1)/w + 1)
	}
	buckets := map[int]int{}
	for i, x := range nums {
		b := getBucket(x)
		if _, ok := buckets[b]; ok {
			return true
		}
		if v, ok := buckets[b-1]; ok && abs(x-v) <= t {
			return true
		}
		if v, ok := buckets[b+1]; ok && abs(x-v) <= t {
			return true
		}
		buckets[b] = x
		if i >= k {
			delete(buckets, getBucket(nums[i-k]))
		}
	}
	return false
}

func abs(x int) int { if x < 0 { return -x }; return x }

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))      // true
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3)) // false
}
