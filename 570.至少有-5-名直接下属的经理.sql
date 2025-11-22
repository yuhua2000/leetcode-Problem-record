
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