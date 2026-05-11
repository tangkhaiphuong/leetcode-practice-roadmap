-- Problem 090: Nth Highest Salary (LeetCode #177)
-- Difficulty: Medium
-- Categories: Database / SQL
--
-- Description:
--   Write a SQL function to return the n-th highest salary from the Employee
--   table. Return null if there's no such salary.
--
-- Schema:
--   Employee(id INT PRIMARY KEY, salary INT)
--
-- Approach:
--   Use SELECT DISTINCT salary ORDER BY DESC LIMIT 1 OFFSET (N-1).
--   Wrap in a function with a SET statement so the offset is a constant.
--
-- Notes:
--   Many SQL dialects don't allow N-1 directly inside LIMIT; assign to a
--   variable first.

CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  DECLARE M INT;
  SET M = N - 1;
  RETURN (
    SELECT DISTINCT salary
    FROM Employee
    ORDER BY salary DESC
    LIMIT 1 OFFSET M
  );
END

-- ----- Example -----
-- Employee:
--   id | salary
--   ---+-------
--    1 |  100
--    2 |  200
--    3 |  300
-- getNthHighestSalary(2) -> 200
-- getNthHighestSalary(4) -> NULL
