#!/usr/bin/env bash
# Problem 091: Word Frequency (LeetCode #192)
# Difficulty: Medium
# Categories: Bash / Shell
#
# Description:
#   Given words.txt, write a shell command that outputs the word frequency
#   sorted by count descending. Output is one "<word> <count>" per line.
#
# Approach:
#   1. Use `tr -s ' '` to collapse multiple spaces into one (or `xargs -n 1`).
#   2. Use `tr -s ' ' '\n'` to split words on newlines, then sort.
#   3. `uniq -c` to count consecutive duplicates.
#   4. `sort -rn` to sort by count descending.
#   5. `awk '{print $2, $1}'` to swap to "<word> <count>" format.

cat words.txt \
  | tr -s ' ' '\n' \
  | sort \
  | uniq -c \
  | sort -rn \
  | awk '{print $2, $1}'

# ----- Example words.txt -----
# the day is sunny the the
# the sunny is is
#
# Expected output:
#   the 4
#   is 3
#   sunny 2
#   day 1
