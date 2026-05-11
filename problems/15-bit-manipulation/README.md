# Bit Manipulation / Bitmask

3 problems where bit tricks unlock the solution: division without `/`, bit-pattern greedy, and bitmask subset enumeration.

## Core patterns

| Pattern | When to use | Problems |
|---|---|---|
| **Set / clear / toggle / test** | `x |= 1<<k`, `x &^= 1<<k`, `x ^= 1<<k`, `(x>>k)&1` | — |
| **Lowest set bit** | `x & -x` (or `x & (^x + 1)`) | — |
| **Clear lowest set bit** | `x &= x - 1` — useful for popcount loops | — |
| **Bitmask DP / subset enumeration** | `for sub := mask; sub > 0; sub = (sub-1) & mask` | 030 |
| **Doubling / fast exponent** | Walk bits of n; Pow, modular exponentiation | (cross: 077) |
| **Division without `/`** | Repeated doubling subtraction | 089 |
| **Bit-pattern greedy** | Halve, +1, -1 — pick based on trailing bits | 069 |

## Go-specific notes

- **Signed shifts** — Go's `>>` on `int` is arithmetic (sign-extending). Use `uint` if you want logical right shift.
- **Width matters** — for 32-bit clamps, use `int32` types or `math.MaxInt32` / `math.MinInt32` constants. Mind that `-math.MinInt32` overflows int32.
- **`bits.OnesCount`, `bits.LeadingZeros`** — `math/bits` package gives you popcount, leading/trailing zeros, etc.

## Common interview pitfalls

- **`-math.MinInt32` is invalid in int32** — promote to `int64` first.
- **`x++` on uint64 near max** — overflows silently; check before incrementing.
- **Confusing `&` with `&&`** — `&` is bitwise; `&&` is logical short-circuit.
- **Hex vs decimal in masks** — `0xff`, `0x0f` are common; `^uint(0)` gives all-ones for the platform width.
- **Integer Replacement** — for `x % 4 == 3` (`...11`), incrementing creates more trailing zeros (faster path); for `x % 4 == 1` (`...01`), decrementing is one step. Special-case `x == 3` (go to 2 not 4).
- **Meet-in-the-middle subset enumeration** — careful to avoid choosing the empty subset from both halves (gives empty A, which violates "non-empty").

## Suggested practice order

1. **Day 1**: 089 (Divide Two Integers) — repeated doubling
2. **Day 2**: 069 (Integer Replacement) — bit-pattern greedy
3. **Day 3**: 030 (Split Array Same Average) — meet-in-the-middle with bitmask subsets

## Problems in this folder

| # | Problem | Difficulty | Key idea |
|---|---|---|---|
| 030 | Split Array With Same Average | Hard | Meet in the middle (bitmask subsets) |
| 069 | Integer Replacement | Medium | Greedy on trailing bits |
| 089 | Divide Two Integers | Medium | Repeated doubling subtraction |
