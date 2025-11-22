select 
    SalesPerson.name 
FROM 
    SalesPerson 
Where
    SalesPerson.sales_id not in
    (
    select 
        SalesPerson.sales_id
    FROM 
        Orders
    JOIN 
        Company on Orders.com_id = Company.com_id
    JOIN 
        SalesPerson on Orders.sales_id = SalesPerson.sales_id
    Where 
        Company.Name= 'Red'
    )
