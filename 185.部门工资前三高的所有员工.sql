--
-- @lc app=leetcode.cn id=185 lang=mysql
--
-- [185] 部门工资前三高的所有员工
--

-- @lc code=start
# Write your MySQL query statement below

SELECT d.Name AS 'Department', e1.name as 'Employee', e1.salary
FROM Employee e1
    JOIN Department d ON e1.DepartmentId = d.Id
WHERE
    3 > (
        SELECT count(DISTINCT e2.salary)
        FROM Employee e2
        WHERE
            e2.salary > e1.salary
            AND e1.DepartmentId = e2.DepartmentId
    );

-- @lc code=end