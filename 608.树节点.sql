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
