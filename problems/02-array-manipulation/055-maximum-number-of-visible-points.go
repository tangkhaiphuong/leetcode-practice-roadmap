//go:build ignore

// Problem 055: Maximum Number of Visible Points (LeetCode #1610)
// Difficulty: Hard
// Categories: Array, Sliding Window, Math, Geometry
//
// Description:
//   You're at location (px, py) with field-of-view angle (in degrees). You
//   can rotate freely. Given a list of points, return the max number of
//   points you can see at once. Points at the same location as you are
//   always counted.
//
// Approach: Polar Angles + Sliding Window
//   Compute atan2(dy, dx) (in radians) for every non-overlapping point. Sort
//   them. To handle wrap-around, append each angle + 2π. Slide a window where
//   angle[r] - angle[l] <= angle (converted to radians). Track max window
//   size; add count of overlapping (same-location) points at the end.
//
// Complexity:
//   Time:  O(n log n)
//   Space: O(n)

package main

import (
	"fmt"
	"math"
	"sort"
)

func visiblePoints(points [][]int, angle int, location []int) int {
	angles := []float64{}
	overlap := 0
	for _, p := range points {
		dx := p[0] - location[0]
		dy := p[1] - location[1]
		if dx == 0 && dy == 0 {
			overlap++
			continue
		}
		angles = append(angles, math.Atan2(float64(dy), float64(dx)))
	}
	sort.Float64s(angles)
	n := len(angles)
	for i := 0; i < n; i++ {
		angles = append(angles, angles[i]+2*math.Pi)
	}
	cap := float64(angle) * math.Pi / 180.0
	best, l := 0, 0
	for r := 0; r < len(angles); r++ {
		for angles[r]-angles[l] > cap+1e-9 {
			l++
		}
		if r-l+1 > best {
			best = r - l + 1
		}
	}
	return best + overlap
}

func main() {
	fmt.Println(visiblePoints([][]int{{2, 1}, {2, 2}, {3, 3}}, 90, []int{1, 1}))         // 3
	fmt.Println(visiblePoints([][]int{{2, 1}, {2, 2}, {3, 4}, {1, 1}}, 90, []int{1, 1})) // 4
	fmt.Println(visiblePoints([][]int{{1, 0}, {2, 1}}, 13, []int{1, 1}))                 // 1
}
