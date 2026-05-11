//go:build ignore

// Problem 046: Water and Jug Problem (LeetCode #365)
// Difficulty: Medium
// Categories: Math, Number Theory, BFS (alt)
//
// Description:
//   Given two jugs of capacity x and y liters and a target z, determine if
//   you can measure exactly z liters using these operations: fill, empty,
//   pour one into the other (until source empty or target full).
//
// Examples:
//   x=3, y=5, z=4   -> true
//   x=2, y=6, z=5   -> false
//
// Approach: Bezout's Identity
//   Reachable amounts in the union of jugs are exactly the non-negative
//   multiples of gcd(x, y) up to x+y. So z is achievable iff:
//     z == 0 OR (z <= x+y AND z % gcd(x, y) == 0)
//
// Complexity:
//   Time:  O(log(min(x,y)))
//   Space: O(1)

package main

import "fmt"

func canMeasureWater(x, y, z int) bool {
	if z == 0 {
		return true
	}
	if z > x+y {
		return false
	}
	return z%gcd(x, y) == 0
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println(canMeasureWater(3, 5, 4)) // true
	fmt.Println(canMeasureWater(2, 6, 5)) // false
	fmt.Println(canMeasureWater(1, 2, 3)) // true
}
