SELECT 
    r1.requester_id AS id,
    COUNT(*) AS num
FROM
    (
        SELECT requester_id 
        FROM RequestAccepted
        UNION ALL
        SELECT accepter_id 
        FROM RequestAccepted
    ) AS r1
GROUP BY 
    r1.requester_id
ORDER BY 
    num DESC
LIMIT 1;
