--
-- @lc app=leetcode.cn id=177 lang=mysql
--
-- [177] 第N高的薪水
--

-- @lc code=start
CREATE FUNCTION getNthHighestSalary (N INT) RETURNS INT 
  BEGIN 
  DECLARE M INT; 
    SET M = N-1; 
  RETURN (
    # Write your MySQL query statement below.
    select IFNULL(
            (
                select DISTINCT
                    salary
                from Employee
                ORDER BY salary DESC
                LIMIT M, 1
            ), NULL
        ) AS SecondHighestSalary  
);
  END
-- @lc code=end