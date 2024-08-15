/*
表：Employee 
+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| empId       | int     |
| name        | varchar |
| supervisor  | int     |
| salary      | int     |
+-------------+---------+
empId 是该表中具有唯一值的列。
该表的每一行都表示员工的姓名和 id，以及他们的工资和经理的 id。
表：Bonus
+-------------+------+
| Column Name | Type |
+-------------+------+
| empId       | int  |
| bonus       | int  |
+-------------+------+
empId 是该表具有唯一值的列。
empId 是 Employee 表中 empId 的外键(reference 列)。
该表的每一行都包含一个员工的 id 和他们各自的奖金。
编写解决方案，报告每个奖金 少于 1000 的员工的姓名和奖金数额。
以 任意顺序 返回结果表。
结果格式如下所示。
*/

select name, bonus
from Employee
    left join Bonus on Employee.empId = Bonus.empId
where
    bonus is null
    or bonus < 1000;