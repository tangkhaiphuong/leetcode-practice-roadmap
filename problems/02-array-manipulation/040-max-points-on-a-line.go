//go:build ignore

// Problem 040: Max Points on a Line (LeetCode #149)
// Difficulty: Hard
// Categories: Array, Hash Table, Geometry, Math
//
// Description:
//   Given n points on a 2D plane, return the max number of points that lie
//   on the same straight line.
//
// Examples:
//   [[1,1],[2,2],[3,3]]            -> 3
//   [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]] -> 4
//
// Approach: For Each Anchor, Hash Slopes (using GCD pair as key)
//   For each point i, count duplicates and slopes to every other point j.
//   Slope = reduced (dy/dx) using gcd; normalize sign so (-dy,-dx) maps to
//   the same key. Vertical line: key (1, 0). Track max + duplicates.
//
//   Using float slopes loses precision; gcd-reduced pair is safest.
//
// Complexity:
//   Time:  O(n^2)
//   Space: O(n) per anchor

package main

import "fmt"

func maxPoints(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}
	best := 1
	for i := 0; i < n; i++ {
		slopes := map[[2]int]int{}
		dupes, localMax := 1, 0
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			dx := points[j][0] - points[i][0]
			dy := points[j][1] - points[i][1]
			if dx == 0 && dy == 0 {
				dupes++
				continue
			}
			g := gcd(absInt(dx), absInt(dy))
			dx /= g
			dy /= g
			// Normalize sign: ensure dx>0; if dx==0 (vertical), dy=1.
			if dx < 0 || (dx == 0 && dy < 0) {
				dx, dy = -dx, -dy
			}
			slopes[[2]int{dx, dy}]++
			if slopes[[2]int{dx, dy}] > localMax {
				localMax = slopes[[2]int{dx, dy}]
			}
		}
		if localMax+dupes > best {
			best = localMax + dupes
		}
	}
	return best
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a == 0 {
		return 1
	}
	return a
}
func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(maxPoints([][]int{{1, 1}, {2, 2}, {3, 3}}))                     // 3
	fmt.Println(maxPoints([][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}})) // 4
}
