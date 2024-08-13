/*
表: Employee
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| name        | varchar |
| department  | varchar |
| managerId   | int     |
+-------------+---------+
id 是此表的主键（具有唯一值的列）。
该表的每一行表示雇员的名字、他们的部门和他们的经理的id。
如果managerId为空，则该员工没有经理。
没有员工会成为自己的管理者。
*/
select name
from (
        select E1.id, E1.name, count(*) as num
        FROM Employee E1
            JOIN Employee E2 ON E1.id = E2.managerId
        group by
            E1.id
    ) as A
Where
    num >= 5;