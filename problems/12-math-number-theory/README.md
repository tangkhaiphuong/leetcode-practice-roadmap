# Math & Number Theory

13 problems where math (GCD, modular arithmetic, place-value, geometry, primality) is the unlocking insight.

## Core patterns

| Pattern | When to use | Problems |
|---|---|---|
| **Modular arithmetic** | Mod-p residues, large products mod p | 064, 071, 083 |
| **GCD / Bezout's identity** | Reachability via linear combinations | (cross: 046 in 02-array) |
| **Place-value reasoning** | Digit DP, digit counts | 071, 074 |
| **Fast exponentiation** | Pow, modular pow | 077 |
| **Mirror palindrome construction** | Closest palindrome from left half | 005 |
| **Long division with cycle detection** | Recurring decimals | 025 |
| **Next-permutation on digits** | Lex-next number with same digits | 041 |
| **Prime palindrome generation** | Odd-length palindromes from left half + primality | 084 |
| **Sum-of-two-squares search** | Two pointers on `[0, sqrt(c)]` | 072 |
| **Geometry (atan2, slopes, distance)** | Visibility, self-crossing | 047, 055 |
| **Group-by-thousands recursion** | Number-to-words | 002 |
| **Strata + arithmetic series** | Diminishing sums | 083 |

## Common interview pitfalls

- **Mod negative result** — Go's `%` keeps the sign of the dividend. To get a non-negative residue, do `((a % m) + m) % m`.
- **Integer overflow** — Go's `int` is 64-bit but multiplication can still overflow. Cast to `int64` (or `math/big`) when in doubt.
- **`math.Sqrt` precision** — for integer answer, verify by squaring after rounding: `r := int(math.Sqrt(float64(n))); if r*r > n { r-- }`.
- **Even-length palindromes & primality** — every even-length palindrome >11 is divisible by 11, so generate ODD-length palindromes only (084).
- **Float angles & wraparound** — for visibility-cone problems, work in radians and add `2π` to a duplicated copy.
- **Closest palindrome — five candidates** — mirror as-is, mirror±1, plus 10^(L-1)-1 and 10^L+1 to cover length boundaries.
- **Long-division position tracking** — when the same remainder recurs, that's the start of the cycle; wrap with parens.
- **Place-value loop bounds** — for "count of 1s in 0..n", iterate `p = 1, 10, 100, ...` until `p > n`.

## Suggested practice order

1. **Day 1**: 077 (Pow), 072 (Sum of Squares) — fundamental tools
2. **Day 2**: 002, 005 — string/number constructions
3. **Day 3**: 025, 074 — long division & place value
4. **Day 4**: 041, 064 — next perm & modular prefix
5. **Day 5**: 071 (Digit DP) — place-value DP
6. **Day 6**: 084 (Prime Palindrome) — generation + primality
7. **Day 7**: 047, 055 — geometry
8. **Day 8**: 083 — arithmetic series under modulo

## Problems in this folder

| # | Problem | Difficulty | Key idea |
|---|---|---|---|
| 002 | Integer to English Words | Hard | Group-by-thousand recursion |
| 005 | Find the Closest Palindrome | Hard | Mirror half + 5 candidates |
| 025 | Fraction to Recurring Decimal | Medium | Long division + remainder map |
| 041 | Next Greater Element III | Medium | Next permutation on digits |
| 047 | Self Crossing | Hard | Three geometric cases |
| 055 | Maximum Number of Visible Points | Hard | Polar angles + sliding window |
| 064 | Make Sum Divisible by P | Medium | Prefix sum mod p + first-index map |
| 071 | Number of Digit One | Hard | Place-value digit DP |
| 072 | Sum of Square Numbers | Medium | Two pointers on `[0, sqrt(c)]` |
| 074 | Nth Digit | Medium | Bracket by digit-length bands |
| 077 | Pow(x, n) | Medium | Iterative fast exponentiation |
| 083 | Sell Diminishing-Valued Colored Balls | Medium | Strata sum mod p |
| 084 | Prime Palindrome | Medium | Odd-length palindrome gen + primality |
