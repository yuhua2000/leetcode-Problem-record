--
-- @lc app=leetcode.cn id=176 lang=mysql
--
-- [176] 第二高的薪水
--

-- @lc code=start
# Write your MySQL query statement below
select IFNULL(
        (
            select DISTINCT
                salary
            from Employee
            ORDER BY salary DESC
            LIMIT 1, 1
        ), NULL
    ) AS SecondHighestSalary
-- @lc code=end