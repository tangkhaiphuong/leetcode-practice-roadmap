# Database / SQL

2 problems testing SQL fluency and shell pipelines. These are quick wins on LeetCode but interviewers expect *idiomatic* answers, not just correct ones.

## Core patterns

| Pattern | When to use | Tool |
|---|---|---|
| **Window function** (`ROW_NUMBER`, `DENSE_RANK`) | "Nth highest", "top per group" | SQL |
| **`LIMIT ... OFFSET`** | Single-row pagination, "Nth value" | SQL |
| **Subquery / CTE** | Reuse intermediate result, recursion | SQL |
| **`GROUP BY` + `HAVING`** | Filter aggregates | SQL |
| **`tr | sort | uniq -c`** | Tokenize + count + sort | bash |

## SQL idioms

```sql
-- Nth highest with DENSE_RANK (handles ties correctly)
SELECT salary
FROM (SELECT salary, DENSE_RANK() OVER (ORDER BY salary DESC) AS r FROM Employee) t
WHERE r = N;

-- Nth highest with LIMIT/OFFSET (returns NULL automatically if absent)
SELECT DISTINCT salary
FROM Employee
ORDER BY salary DESC
LIMIT 1 OFFSET (N-1);
```

## Common interview pitfalls

- **`LIMIT` doesn't accept expressions in MySQL** — declare a variable first: `DECLARE M INT; SET M = N - 1;` then `LIMIT 1 OFFSET M`.
- **`DENSE_RANK` vs `RANK` vs `ROW_NUMBER`** — DENSE_RANK gives 1,2,2,3; RANK gives 1,2,2,4; ROW_NUMBER gives 1,2,3,4. Pick correctly for "Nth highest distinct salary".
- **NULL behavior** — `SELECT MAX(salary) FROM ...` returns NULL on empty. `LIMIT/OFFSET` returns no row, which the function returns as NULL. Both are usually what LeetCode wants.
- **Bash word splitting** — `tr -s ' '` collapses runs of spaces; pair with `tr -s ' ' '\n'` to convert to newlines for `sort | uniq -c`.
- **`uniq` requires sorted input** — `sort | uniq -c` is the standard idiom; `uniq` alone only collapses *adjacent* duplicates.

## Suggested practice order

1. **Day 1**: 090 (Nth Highest Salary) — try with DENSE_RANK and with LIMIT/OFFSET both, understand the difference
2. **Day 2**: 091 (Word Frequency) — practice the bash counting pipeline
3. **Bonus**: try LeetCode SQL problems 175-185 for breadth

## Problems in this folder

| # | Problem | Difficulty | Tool |
|---|---|---|---|
| 090 | Nth Highest Salary | Medium | SQL function with LIMIT OFFSET |
| 091 | Word Frequency | Medium | bash: tr | sort | uniq -c | sort -rn |
