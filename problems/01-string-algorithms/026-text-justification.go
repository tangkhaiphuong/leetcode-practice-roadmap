//go:build ignore

// Problem 026: Text Justification (LeetCode #68)
// Difficulty: Hard
// Categories: String, Array, Simulation
//
// Description:
//   Given an array of words and a maxWidth, format the text so each line has
//   exactly maxWidth chars. Pack as many words as possible per line, then
//   pad with spaces. Distribute extra spaces left-heavy (i.e. left gaps get
//   one more space if not divisible). The last line is left-justified.
//
// Examples:
//   ["This","is","an","example","of","text","justification."], 16
//   -> ["This    is    an", "example  of text", "justification.  "]
//
// Approach: Two-Pass per Line
//   Pass 1: greedily fill words[i..j-1] until next would exceed width.
//   Pass 2: distribute (extraSpaces) over (j-i-1) gaps:
//     base = extra / gaps; rem = extra % gaps.
//     First rem gaps get base+1, rest base.
//   Special: if it's the last line OR only one word, left-justify.
//
// Complexity:
//   Time:  O(total chars)
//   Space: O(total chars) output

package main

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	var res []string
	i := 0
	for i < len(words) {
		// Greedy fit
		j, lineLen := i, len(words[i])
		for j+1 < len(words) && lineLen+1+len(words[j+1]) <= maxWidth {
			j++
			lineLen += 1 + len(words[j])
		}
		// Build line
		gaps := j - i
		isLast := j == len(words)-1
		var line strings.Builder
		if gaps == 0 || isLast {
			// Left-justify
			line.WriteString(words[i])
			for k := i + 1; k <= j; k++ {
				line.WriteByte(' ')
				line.WriteString(words[k])
			}
			line.WriteString(strings.Repeat(" ", maxWidth-line.Len()))
		} else {
			wordsLen := lineLen - gaps // sum of word lengths
			totalSpaces := maxWidth - wordsLen
			base := totalSpaces / gaps
			rem := totalSpaces % gaps
			for k := i; k <= j; k++ {
				line.WriteString(words[k])
				if k < j {
					n := base
					if k-i < rem {
						n++
					}
					line.WriteString(strings.Repeat(" ", n))
				}
			}
		}
		res = append(res, line.String())
		i = j + 1
	}
	return res
}

func main() {
	out := fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16)
	for _, l := range out {
		fmt.Printf("%q\n", l)
	}
}
