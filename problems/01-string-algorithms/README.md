# String & String Algorithms

31 problems covering palindromes, encoding/decoding, parsing, parentheses, regex, KMP, and rolling hash.

## Core patterns

| Pattern | When to use | Problems in this folder |
|---|---|---|
| **Expand around center** | Find / count palindromic substrings | 001, 017 |
| **Two-pointer merge / scan** | Compare strings, edit distance, parsing | 003, 013, 021 |
| **DP on prefixes** | Match / count / segment | 008, 011, 012, 014, 015, 017, 018, 031 |
| **Stack of indices** | Balanced parens, monotonic min/max | 010, 016, 029 |
| **Sliding window + map** | Anagrams, distinct chars, k-window props | 004, 023 |
| **Rolling hash / KMP** | Substring search, longest match | 006, 027 |
| **Greedy with bound** | Smallest/largest concatenations | 009, 016, 024 |
| **State machine / flags** | Number, IP, email validation | 020, 022 |
| **Recursion with cap** | Decoded length explosion | 029 |

## Common interview pitfalls

- **`byte` vs `rune`** — use `byte` for ASCII (LeetCode default), `rune` for Unicode iteration. Iterating `s[i]` gives `byte`; `range s` gives `rune`.
- **String immutability** — `s[i] = 'x'` won't compile. Convert to `[]byte` for in-place edits.
- **Off-by-one in expand-around-center** — remember even-length centers (`expand(i, i+1)`) AND odd-length (`expand(i, i)`).
- **Leading zeros / negative signs** — atoi-style problems silently fail when you forget `+`/`-` handling or skip whitespace.
- **DP base cases** — for matching DP (`dp[0][j]`), think about whether the empty pattern matches the empty string and how `*`/`a*b*` zero-out preceding chars.
- **Allocations** — `strings.Builder` beats `+=` in tight loops; `strings.Fields` splits on any whitespace AND trims (vs `strings.Split` which keeps empties).

## Suggested practice order

1. **Warmup** (1 day each): 001, 003, 004, 007, 012, 013
2. **DP core** (Day 2-3): 008, 011, 014, 017, 018
3. **Stack/Greedy** (Day 4): 010, 015, 016, 029
4. **Hard parsing** (Day 5): 020, 022, 023
5. **Niche but classic** (Day 6+): 005, 006, 024, 025, 026, 027, 030, 031

## Problems in this folder

| # | Problem | Difficulty | Key idea |
|---|---|---|---|
| 001 | Longest Palindromic Substring | Medium | Expand around center |
| 002 | Integer to English Words | Hard | Group-by-thousand recursion |
| 003 | Compare Version Numbers | Medium | Two-pointer parse |
| 004 | Longest Substring Without Repeating Characters | Medium | Sliding window + last-seen map |
| 005 | Find the Closest Palindrome | Hard | Mirror half + 5 candidates |
| 006 | Longest Duplicate Substring | Hard | Binary search + Rabin-Karp |
| 007 | String to Integer (atoi) | Medium | Linear scan + overflow guard |
| 008 | Regular Expression Matching | Hard | 2D DP, `*` consumes 0 or 1+ |
| 009 | Largest Number | Medium | Custom comparator on `a+b` vs `b+a` |
| 010 | Longest Valid Parentheses | Hard | Stack of indices with sentinel |
| 011 | Interleaving String | Medium | 2D DP, rolling 1D row |
| 012 | Decode Ways | Medium | Linear DP, two rolling vars |
| 013 | Reverse Words in a String | Medium | `strings.Fields` + reverse |
| 014 | Wildcard Matching | Hard | 2D DP, `*` = any seq |
| 015 | Valid Parenthesis String | Medium | Track lo/hi range of open count |
| 016 | Remove K Digits | Medium | Monotonic increasing stack |
| 017 | Palindrome Partitioning II | Hard | Palindrome table + cuts DP |
| 018 | Decode Ways II | Hard | DP w/ case enum on `*` |
| 019 | Repeated String Match | Medium | Repeat until length covers |
| 020 | Validate IP Address | Medium | Branch on `.` vs `:` |
| 021 | One Edit Distance | Medium | Length-diff branch + scan |
| 022 | Valid Number | Hard | Boolean flags single pass |
| 023 | Substring with Concatenation of All Words | Hard | Sliding window of words |
| 024 | Strong Password Checker | Hard | Case analysis on length regimes |
| 025 | Fraction to Recurring Decimal | Medium | Long division + remainder map |
| 026 | Text Justification | Hard | Two-pass per line |
| 027 | Shortest Palindrome | Hard | KMP failure function |
| 028 | Making File Names Unique | Medium | Map of next suffix to try |
| 029 | Decoded String at Index | Medium | Reverse walk without materializing |
| 030 | Split Array With Same Average | Hard | Meet in the middle |
| 031 | Count The Repetitions | Hard | Cycle detection on s2-pointer |
