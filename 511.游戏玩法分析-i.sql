--
-- @lc app=leetcode.cn id=511 lang=mysql
--
-- [511] 游戏玩法分析 I
--

/* Activity 表：
+-----------+-----------+------------+--------------+
| player_id | device_id | event_date | games_played |
+-----------+-----------+------------+--------------+
| 1         | 2         | 2016-03-01 | 5            |
| 1         | 2         | 2016-05-02 | 6            |
| 2         | 3         | 2017-06-25 | 1            |
| 3         | 1         | 2016-03-02 | 0            |
| 3         | 4         | 2018-07-03 | 5            |
+-----------+-----------+------------+--------------+ */
-- @lc code=start
# Write your MySQL query statement below
/* SELECT player_id, min(event_date) AS 'first_login' FROM Activity GROUP BY player_id; */

SELECT X.player_id, X.event_date AS 'first_login' FROM (
    SELECT A.player_id,A.event_date,RANK() OVER(
        PARTITION BY
         A.player_id
          ORDER BY
          A.event_date


)as rnk FROM Activity A ) X 

WHERE x.rnk = 1;

-- @lc code=end