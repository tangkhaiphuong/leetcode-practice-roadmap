//go:build ignore

// Problem 054: Frequency of the Most Frequent Element (LeetCode #1838)
// Difficulty: Medium
// Categories: Array, Sliding Window, Binary Search, Sorting, Greedy
//
// Description:
//   Given nums and k, you can increment any element by 1, at most k total
//   times. Return the maximum possible frequency of any element.
//
// Examples:
//   nums=[1,2,4], k=5     -> 3 (add 5 to make all 4)
//   nums=[1,4,8,13], k=5  -> 2
//
// Approach: Sort + Sliding Window
//   Sort. To raise all elements in nums[l..r] to nums[r], cost =
//   nums[r]*(r-l+1) - sum(nums[l..r]). Slide r right; while cost > k, advance
//   l. Track max window size.
//
// Complexity:
//   Time:  O(n log n)
//   Space: O(1)

package main

import (
	"fmt"
	"sort"
)

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	best, l, sum := 0, 0, 0
	for r := 0; r < len(nums); r++ {
		sum += nums[r]
		for nums[r]*(r-l+1)-sum > k {
			sum -= nums[l]
			l++
		}
		if r-l+1 > best {
			best = r - l + 1
		}
	}
	return best
}

func main() {
	fmt.Println(maxFrequency([]int{1, 2, 4}, 5))      // 3
	fmt.Println(maxFrequency([]int{1, 4, 8, 13}, 5))  // 2
	fmt.Println(maxFrequency([]int{3, 9, 6}, 2))      // 1
}
