/*
SQL Schema
Pandas Schema
表：Tree

+-------------+------+
| Column Name | Type |
+-------------+------+
| id          | int  |
| p_id        | int  |
+-------------+------+
id 是该表中具有唯一值的列。
该表的每行包含树中节点的 id 及其父节点的 id 信息。
给定的结构总是一个有效的树。
 

树中的每个节点可以是以下三种类型之一：

"Leaf"：节点是叶子节点。
"Root"：节点是树的根节点。
"lnner"：节点既不是叶子节点也不是根节点。
编写一个解决方案来报告树中每个节点的类型。

以 任意顺序 返回结果表。

结果格式如下所示。

 

示例 1：


输入：
Tree table:
+----+------+
| id | p_id |
+----+------+
| 1  | null |
| 2  | 1    |
| 3  | 1    |
| 4  | 2    |
| 5  | 2    |
+----+------+
输出：
+----+-------+
| id | type  |
+----+-------+
| 1  | Root  |
| 2  | Inner |
| 3  | Leaf  |
| 4  | Leaf  |
| 5  | Leaf  |
+----+-------+
解释：
节点 1 是根节点，因为它的父节点为空，并且它有子节点 2 和 3。
节点 2 是一个内部节点，因为它有父节点 1 和子节点 4 和 5。
节点 3、4 和 5 是叶子节点，因为它们有父节点而没有子节点。
示例 2：


输入：
Tree table:
+----+------+
| id | p_id |
+----+------+
| 1  | null |
+----+------+
输出：
+----+-------+
| id | type  |
+----+-------+
| 1  | Root  |
+----+-------+
解释：如果树中只有一个节点，则只需要输出其根属性。
*/



WITH
Roots AS (
    SELECT id, 'Root' AS type
    FROM Tree
    WHERE p_id IS NULL
),
Leaves AS (
    SELECT id, 'Leaf' AS type
    FROM Tree
    WHERE id NOT IN (SELECT p_id FROM Tree WHERE p_id IS NOT NULL UNION SELECT id FROM Roots)
),
Inners AS (
    SELECT id, 'Inner' AS type
    FROM Tree
    WHERE id NOT IN (
        SELECT id FROM Roots
        UNION
        SELECT id FROM Leaves
    )
)
SELECT * FROM Roots
UNION ALL
SELECT * FROM Leaves
UNION ALL
SELECT * FROM Inners;

--  case

SELECT
    t.id,
    CASE
        WHEN t.p_id IS NULL THEN 'Root'
        WHEN NOT EXISTS (SELECT 1 FROM Tree c WHERE c.p_id = t.id) THEN 'Leaf'
        ELSE 'Inner'
    END AS type
FROM Tree t;
