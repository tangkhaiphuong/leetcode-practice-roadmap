//go:build ignore

// Problem 076: Longest Common Subpath (LeetCode #1923)
// Difficulty: Hard
// Categories: Binary Search, Rolling Hash
//
// Description:
//   Given n cities and m friends. Each friend has a path (an array of city
//   ids). Return the longest contiguous subpath common to ALL friends' paths.
//
// Examples:
//   n=5, paths=[[0,1,2,3,4],[2,3,4],[4,0,1,2,3]] -> 2  (e.g. [2,3])
//
// Approach: Binary Search Length L + Rolling Hash Intersection
//   Binary search L in [0, min(path_lengths)]. For a given L:
//     - For each path, compute the set of length-L substring hashes.
//     - Intersect across all paths.
//     - If non-empty, L is feasible; try larger.
//   Use polynomial hashing modulo a large prime; verify hits with actual
//   subpath equality on collision (paranoid version) — for competitive code
//   double-hash often suffices.
//
// Complexity:
//   Time:  O(sum(|path|) * log(min |path|))
//   Space: O(sum(|path|))

package main

import "fmt"

const mod76 = uint64(1_000_000_007)
const base76 = uint64(100003)

func hashSet(path []int, L int) map[uint64]bool {
	if L > len(path) {
		return nil
	}
	var pw uint64 = 1
	for i := 0; i < L; i++ {
		pw = pw * base76 % mod76
	}
	var h uint64
	for i := 0; i < L; i++ {
		h = (h*base76 + uint64(path[i])) % mod76
	}
	set := map[uint64]bool{h: true}
	for i := L; i < len(path); i++ {
		h = (h*base76 + uint64(path[i])) % mod76
		h = (h + mod76 - uint64(path[i-L])*pw%mod76) % mod76
		set[h] = true
	}
	return set
}

func feasible(paths [][]int, L int) bool {
	if L == 0 {
		return true
	}
	common := hashSet(paths[0], L)
	if common == nil {
		return false
	}
	for _, p := range paths[1:] {
		s := hashSet(p, L)
		if s == nil {
			return false
		}
		for k := range common {
			if !s[k] {
				delete(common, k)
			}
		}
		if len(common) == 0 {
			return false
		}
	}
	return len(common) > 0
}

func longestCommonSubpath(n int, paths [][]int) int {
	lo, hi := 0, len(paths[0])
	for _, p := range paths {
		if len(p) < hi {
			hi = len(p)
		}
	}
	ans := 0
	for lo <= hi {
		mid := (lo + hi) / 2
		if feasible(paths, mid) {
			ans = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return ans
}

func main() {
	fmt.Println(longestCommonSubpath(5, [][]int{{0, 1, 2, 3, 4}, {2, 3, 4}, {4, 0, 1, 2, 3}})) // 2
	fmt.Println(longestCommonSubpath(3, [][]int{{0}, {1}, {2}}))                               // 0
}
