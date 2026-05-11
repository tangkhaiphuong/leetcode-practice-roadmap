//go:build ignore

// Problem 065: Line Reflection (LeetCode #356)
// Difficulty: Medium
// Categories: Hash Table, Math, Geometry
//
// Description:
//   Given n points, return true iff there exists a vertical line such that
//   every point's mirror across that line is also in the input.
//
// Examples:
//   [[1,1],[-1,1]]                -> true
//   [[1,1],[-1,-1]]               -> false
//   [[1,2],[2,2],[1,4],[2,4]]     -> true (line x=1.5)
//
// Approach: Median X via Min/Max + Hash Set
//   The candidate line of symmetry has x = (minX + maxX) / 2 (working with
//   2x to avoid fractions: target = minX + maxX, mirror of (x,y) is
//   (target-x, y)).
//   For each point, check if its mirror is present.
//
// Complexity:
//   Time:  O(n)
//   Space: O(n)

package main

import (
	"fmt"
	"math"
)

func isReflected(points [][]int) bool {
	if len(points) == 0 {
		return true
	}
	minX, maxX := math.MaxInt32, math.MinInt32
	set := map[[2]int]bool{}
	for _, p := range points {
		if p[0] < minX {
			minX = p[0]
		}
		if p[0] > maxX {
			maxX = p[0]
		}
		set[[2]int{p[0], p[1]}] = true
	}
	target := minX + maxX
	for p := range set {
		mirror := [2]int{target - p[0], p[1]}
		if !set[mirror] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isReflected([][]int{{1, 1}, {-1, 1}}))               // true
	fmt.Println(isReflected([][]int{{1, 1}, {-1, -1}}))              // false
	fmt.Println(isReflected([][]int{{1, 2}, {2, 2}, {1, 4}, {2, 4}})) // true
}
