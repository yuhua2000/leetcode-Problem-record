--
-- @lc app=leetcode.cn id=178 lang=mysql
--
-- [178] 分数排名
--

-- @lc code=start
# Write your MySQL query statement below
SELECT S.score, DENSE_RANK() OVER (
        ORDER BY S.score desc
    ) AS 'rank'
from Scores S;

SELECT S1.score, (
        SELECT COUNT(DISTINCT S2.score)
        FROM Scores S2
        WHERE
            S2.score >= S1.score
    ) as 'rank'
FROM Scores S1
ORDER BY S1.score DESC;
-- @lc code=end