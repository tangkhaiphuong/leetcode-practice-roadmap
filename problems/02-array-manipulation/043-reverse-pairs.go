//go:build ignore

// Problem 043: Reverse Pairs (LeetCode #493)
// Difficulty: Hard
// Categories: Array, Binary Search, Divide & Conquer, Merge Sort
//
// Description:
//   Given nums, return the count of pairs (i, j) such that i < j and
//   nums[i] > 2 * nums[j].
//
// Examples:
//   [1,3,2,3,1] -> 2
//   [2,4,3,5,1] -> 3
//
// Approach: Merge Sort Counting
//   While merging two sorted halves, count for each i in left how many j in
//   right satisfy left[i] > 2*right[j]. Two pointers across right makes the
//   count O(n) per merge layer.
//
// Complexity:
//   Time:  O(n log n)
//   Space: O(n)

package main

import "fmt"

func reversePairs(nums []int) int {
	tmp := make([]int, len(nums))
	return mergeSort(nums, tmp, 0, len(nums)-1)
}

func mergeSort(a, tmp []int, lo, hi int) int {
	if lo >= hi {
		return 0
	}
	mid := (lo + hi) / 2
	count := mergeSort(a, tmp, lo, mid) + mergeSort(a, tmp, mid+1, hi)
	// Count cross pairs.
	j := mid + 1
	for i := lo; i <= mid; i++ {
		for j <= hi && a[i] > 2*a[j] {
			j++
		}
		count += j - (mid + 1)
	}
	// Merge.
	i, k, p := lo, mid+1, lo
	for i <= mid && k <= hi {
		if a[i] <= a[k] {
			tmp[p] = a[i]
			i++
		} else {
			tmp[p] = a[k]
			k++
		}
		p++
	}
	for i <= mid {
		tmp[p] = a[i]
		i++
		p++
	}
	for k <= hi {
		tmp[p] = a[k]
		k++
		p++
	}
	copy(a[lo:hi+1], tmp[lo:hi+1])
	return count
}

func main() {
	fmt.Println(reversePairs([]int{1, 3, 2, 3, 1})) // 2
	fmt.Println(reversePairs([]int{2, 4, 3, 5, 1})) // 3
}
