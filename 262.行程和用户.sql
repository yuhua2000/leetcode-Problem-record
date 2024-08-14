--
-- @lc app=leetcode.cn id=262 lang=mysql
--
-- [262] 行程和用户
--

-- @lc code=start
# Write your MySQL query statement below
SELECT request_at as 'Day', ROUND(
        1 - (
            (
                select count(*)
                FROM Trips T2
                WHERE
                    T2.request_at = T1.request_at
                    AND T2.client_id in (
                        select users_id
                        FROM Users
                        WHERE
                            banned = 'No'
                    )
                    AND T2.driver_id in (
                        select users_id
                        FROM Users
                        WHERE
                            banned = 'No'
                    )
                    AND T2.status = 'completed'
            ) / (
                select count(*)
                FROM Trips T2
                WHERE
                    T2.request_at = T1.request_at
                    AND T2.client_id in (
                        select users_id
                        FROM Users
                        WHERE
                            banned = 'No'
                    )
                    AND T2.driver_id in (
                        select users_id
                        FROM Users
                        WHERE
                            banned = 'No'
                    )
            )
        ), 2
    ) as 'Cancellation Rate'
FROM Trips T1
Where (
        select count(*)
        FROM Trips T2
        WHERE
            T2.request_at = T1.request_at
            AND T2.client_id in (
                select users_id
                FROM Users
                WHERE
                    banned = 'No'
            )
            AND T2.driver_id in (
                select users_id
                FROM Users
                WHERE
                    banned = 'No'
            )
    ) != 0
GROUP BY
    request_at
HAVING
    request_at in (
        '2013-10-01',
        '2013-10-02',
        '2013-10-03'
    )


    SELECT T.request_at AS `Day`, 
	ROUND(
			SUM(
				IF(T.STATUS = 'completed',0,1)
			)
			/ 
			COUNT(T.STATUS),
			2
	) AS `Cancellation Rate`
FROM Trips AS T
JOIN Users AS U1 ON (T.client_id = U1.users_id AND U1.banned ='No')
JOIN Users AS U2 ON (T.driver_id = U2.users_id AND U2.banned ='No')
WHERE T.request_at BETWEEN '2013-10-01' AND '2013-10-03'
GROUP BY T.request_at

-- @lc code=end