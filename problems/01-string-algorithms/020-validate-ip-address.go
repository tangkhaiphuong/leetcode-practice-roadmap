//go:build ignore

// Problem 020: Validate IP Address (LeetCode #468)
// Difficulty: Medium
// Categories: String
//
// Description:
//   Return "IPv4", "IPv6", or "Neither" for the input string.
//   IPv4: x1.x2.x3.x4 where each xi is 0..255 with no leading zeros.
//   IPv6: 8 groups of 1-4 hex digits separated by ':'. Leading zeros allowed
//   per group; case-insensitive. (No '::' shortform per LeetCode.)
//
// Examples:
//   "172.16.254.1"   -> "IPv4"
//   "2001:0db8:85a3:0:0:8A2E:0370:7334" -> "IPv6"
//   "256.256.256.256" -> "Neither"
//
// Approach: Branch on whether input contains '.' or ':'.
//   Validate by splitting and checking each chunk.
//
// Complexity:
//   Time:  O(n)
//   Space: O(n)

package main

import (
	"fmt"
	"strings"
)

func validIPAddress(ip string) string {
	if strings.Contains(ip, ".") && validIPv4(ip) {
		return "IPv4"
	}
	if strings.Contains(ip, ":") && validIPv6(ip) {
		return "IPv6"
	}
	return "Neither"
}

func validIPv4(s string) bool {
	parts := strings.Split(s, ".")
	if len(parts) != 4 {
		return false
	}
	for _, p := range parts {
		if len(p) == 0 || len(p) > 3 {
			return false
		}
		if len(p) > 1 && p[0] == '0' {
			return false
		}
		n := 0
		for _, c := range p {
			if c < '0' || c > '9' {
				return false
			}
			n = n*10 + int(c-'0')
		}
		if n > 255 {
			return false
		}
	}
	return true
}

func validIPv6(s string) bool {
	parts := strings.Split(s, ":")
	if len(parts) != 8 {
		return false
	}
	for _, p := range parts {
		if len(p) == 0 || len(p) > 4 {
			return false
		}
		for _, c := range p {
			if !((c >= '0' && c <= '9') ||
				(c >= 'a' && c <= 'f') ||
				(c >= 'A' && c <= 'F')) {
				return false
			}
		}
	}
	return true
}

func main() {
	for _, s := range []string{
		"172.16.254.1",
		"2001:0db8:85a3:0:0:8A2E:0370:7334",
		"256.256.256.256",
		"2001:0db8:85a3::8A2E:0370:7334",
		"1.0.1.",
		"01.01.01.01",
	} {
		fmt.Printf("%-40s -> %s\n", s, validIPAddress(s))
	}
}
