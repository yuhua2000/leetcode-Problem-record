--
-- @lc app=leetcode.cn id=184 lang=mysql
--
-- [184] 部门工资最高的员工
--

-- @lc code=start
# Write your MySQL query statement below
SELECT Department.name AS 'Department', Employee.name AS 'Employee', Employee.salary AS 'Salary'
FROM Employee
    JOIN Department ON Employee.DepartmentId = Department.id
WHERE (
        Employee.DepartmentId,
        Employee.salary
    ) in (
        select departmentId, max(salary)
        FROM Employee
        GROUP BY
            departmentId
    );
-- @lc code=end