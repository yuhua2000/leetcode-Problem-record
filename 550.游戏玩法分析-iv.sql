/* Table: Activity
+--------------+---------+
| Column Name  | Type    |
+--------------+---------+
| player_id    | int     |
| device_id    | int     |
| event_date   | date    |
| games_played | int     |
+--------------+---------+
（player_id，event_date）是此表的主键（具有唯一值的列的组合）。 */
select IFNULL(
        round(
            count(distinct (Result.player_id)) / count(distinct (Activity.player_id)), 2
        ), 0
    ) as fraction
FROM (
        select Activity.player_id as player_id
        FROM (
                SELECT player_id, DATE_ADD(
                        MIN(event_date), INTERVAL 1 DAY
                    ) as second_date
                from Activity
                GROUP BY
                    player_id
            ) as Expected, Activity
        WHERE
            Activity.event_date = Expected.second_date
            AND Activity.player_id = Expected.player_id
    ) as Result, Activity