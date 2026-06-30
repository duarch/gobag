Cheat Sheet — HackerRank SQL (Intermediate)

A certificação SQL (Intermediate) do HackerRank normalmente vai muito além de SELECT, WHERE e GROUP BY. Ela costuma avaliar se você consegue resolver problemas de análise de dados usando SQL.

Os assuntos que mais aparecem são:

1. JOINs
2. GROUP BY + HAVING
3. Funções de agregação
4. Window Functions
5. CTE (WITH)
6. Subqueries
7. CASE
8. Funções de data
9. Ranking
10. Problemas envolvendo “Top N”
11. Diferença entre ROW_NUMBER, RANK e DENSE_RANK
12. EXISTS
13. SELF JOIN
14. UNION / UNION ALL
15. NULL
16. Strings
17. Lógica de negócios

⸻

1. INNER JOIN

Retorna apenas registros existentes nas duas tabelas.

SELECT
    c.customer_name,
    o.order_id
FROM Customers c
INNER JOIN Orders o
ON c.customer_id = o.customer_id;

⸻

2. LEFT JOIN

Retorna todos da esquerda.

SELECT
    c.customer_name,
    o.order_id
FROM Customers c
LEFT JOIN Orders o
ON c.customer_id = o.customer_id;

Clientes sem pedidos:

SELECT *
FROM Customers c
LEFT JOIN Orders o
ON c.customer_id = o.customer_id
WHERE o.order_id IS NULL;

⸻

3. RIGHT JOIN

SELECT *
FROM Orders o
RIGHT JOIN Customers c
ON o.customer_id = c.customer_id;

⸻

4. FULL OUTER JOIN

Nem todo banco suporta.

SELECT *
FROM A
FULL OUTER JOIN B
ON A.id = B.id;

⸻

5. GROUP BY

SELECT
    department,
    COUNT(*)
FROM Employees
GROUP BY department;

⸻

6. HAVING

Filtra grupos.

SELECT
    department,
    COUNT(*)
FROM Employees
GROUP BY department
HAVING COUNT(*) > 5;

⸻

7. COUNT

COUNT(*)
COUNT(column)
COUNT(DISTINCT customer_id)

⸻

8. SUM

SELECT
SUM(price)
FROM Products;

⸻

9. AVG

SELECT AVG(salary)
FROM Employees;

⸻

10. MIN e MAX

SELECT
MIN(price),
MAX(price)
FROM Products;

⸻

11. CASE

Muito frequente.

SELECT
employee,
CASE
    WHEN salary > 10000 THEN 'Senior'
    WHEN salary > 5000 THEN 'Pleno'
    ELSE 'Junior'
END
FROM Employees;

⸻

12. COALESCE

Troca NULL.

SELECT
COALESCE(phone,'Sem telefone')
FROM Customers;

⸻

13. IFNULL

(MySQL)

IFNULL(phone,'Sem telefone')

⸻

14. NULLIF

NULLIF(score,0)

⸻

15. DISTINCT

SELECT DISTINCT department
FROM Employees;

⸻

16. ORDER BY

ORDER BY salary DESC

⸻

17. LIMIT

LIMIT 10;

⸻

18. OFFSET

LIMIT 10 OFFSET 20;

⸻

19. UNION

Remove duplicados.

SELECT city FROM A
UNION
SELECT city FROM B;

⸻

20. UNION ALL

Mantém duplicados.

SELECT city FROM A
UNION ALL
SELECT city FROM B;

⸻

21. EXISTS

Muito usado.

SELECT *
FROM Customers c
WHERE EXISTS (
    SELECT 1
    FROM Orders o
    WHERE o.customer_id=c.customer_id
);

⸻

22. IN

WHERE id IN (1,2,3)

Subquery:

WHERE customer_id IN
(
SELECT customer_id
FROM Orders
)

⸻

23. NOT EXISTS

SELECT *
FROM Customers c
WHERE NOT EXISTS
(
SELECT 1
FROM Orders o
WHERE o.customer_id=c.customer_id
);

⸻

24. Subquery Escalar

SELECT
employee,
(
SELECT AVG(salary)
FROM Employees
)
FROM Employees;

⸻

25. CTE (WITH)

Muito comum.

WITH Sales AS
(
SELECT
customer_id,
SUM(total) total_sales
FROM Orders
GROUP BY customer_id
)
SELECT *
FROM Sales;

⸻

26. CTE Recursiva

Quase nunca cai.

WITH RECURSIVE Numbers AS
(
SELECT 1
UNION ALL
SELECT n+1
FROM Numbers
WHERE n<10
)
SELECT *
FROM Numbers;

⸻

27. ROW_NUMBER()

Mais cobrada.

SELECT
employee,
salary,
ROW_NUMBER()
OVER(ORDER BY salary DESC)
FROM Employees;

⸻

28. RANK()

Empates pulam posições.

100
100
90
Resultado
1
1
3
RANK()
OVER(ORDER BY salary DESC)

⸻

29. DENSE_RANK()

Empates NÃO pulam.

100
100
90
Resultado
1
1
2
DENSE_RANK()
OVER(ORDER BY salary DESC)

⸻

30. PARTITION BY

SELECT
employee,
department,
salary,
ROW_NUMBER()
OVER(
PARTITION BY department
ORDER BY salary DESC
)
FROM Employees;

⸻

31. LAG()

Valor anterior.

SELECT
month,
sales,
LAG(sales)
OVER(ORDER BY month)
FROM Sales;

⸻

32. LEAD()

Próximo valor.

LEAD(sales)
OVER(ORDER BY month)

⸻

33. FIRST_VALUE()

FIRST_VALUE(price)
OVER(
ORDER BY price DESC
)

⸻

34. LAST_VALUE()

LAST_VALUE(price)
OVER(...)

⸻

35. SUM OVER

Acumulado.

SELECT
month,
sales,
SUM(sales)
OVER(
ORDER BY month
)
FROM Sales;

⸻

36. AVG OVER

AVG(salary)
OVER()

⸻

37. Percentual

salary /
SUM(salary) OVER()

⸻

38. Top N por grupo

Muito comum.

WITH Ranking AS
(
SELECT *,
ROW_NUMBER()
OVER(
PARTITION BY department
ORDER BY salary DESC
) rn
FROM Employees
)
SELECT *
FROM Ranking
WHERE rn<=3;

⸻

39. Segundo maior salário

Pergunta clássica.

SELECT MAX(salary)
FROM Employees
WHERE salary <
(
SELECT MAX(salary)
FROM Employees
);

Outra forma:

SELECT salary
FROM
(
SELECT
salary,
DENSE_RANK()
OVER(
ORDER BY salary DESC
) r
FROM Employees
)x
WHERE r=2;

⸻

40. SELF JOIN

SELECT
e.name,
m.name manager
FROM Employees e
LEFT JOIN Employees m
ON e.manager_id=m.id;

⸻

41. Strings

Maiúsculas

UPPER(name)

Minúsculas

LOWER(name)

Tamanho

LENGTH(name)

Substring

SUBSTRING(name,1,3)

Concatenar

CONCAT(first,last)

⸻

42. Datas

Ano

YEAR(order_date)

Mês

MONTH(order_date)

Dia

DAY(order_date)

Diferença

DATEDIFF(end,start)

Adicionar dias

DATE_ADD(order_date,INTERVAL 30 DAY)

⸻

43. Problema clássico: cliente que mais comprou

SELECT
customer_id,
SUM(total)
FROM Orders
GROUP BY customer_id
ORDER BY SUM(total) DESC
LIMIT 1;

⸻

44. Produtos nunca vendidos

SELECT p.*
FROM Products p
LEFT JOIN Orders o
ON p.id=o.product_id
WHERE o.product_id IS NULL;

⸻

45. Média por departamento

SELECT
department,
AVG(salary)
FROM Employees
GROUP BY department;

⸻

46. Funcionários acima da média

SELECT *
FROM Employees
WHERE salary >
(
SELECT AVG(salary)
FROM Employees
);

⸻

47. Total acumulado

SELECT
date,
sales,
SUM(sales)
OVER(
ORDER BY date
)
FROM Sales;

⸻

48. Ranking por departamento

SELECT
employee,
department,
salary,
RANK()
OVER(
PARTITION BY department
ORDER BY salary DESC
)
FROM Employees;

⸻

49. Diferença para o mês anterior

SELECT
month,
sales,
sales-
LAG(sales)
OVER(ORDER BY month)
AS difference
FROM Sales;

⸻

50. Ordem de execução do SQL (muito importante)

FROM
JOIN
WHERE
GROUP BY
HAVING
SELECT
DISTINCT
WINDOW FUNCTIONS
ORDER BY
LIMIT

⸻

Padrões de questões mais frequentes

Top 3 salários por departamento

ROW_NUMBER()
PARTITION BY

⸻

Encontrar duplicados

GROUP BY
HAVING COUNT(*)>1

⸻

Encontrar registros sem correspondência

LEFT JOIN
IS NULL

ou

NOT EXISTS

⸻

Ranking

ROW_NUMBER
RANK
DENSE_RANK

⸻

Total acumulado

SUM()
OVER()

⸻

Comparação com linha anterior

LAG()

⸻

Comparação com próxima linha

LEAD()

⸻

Pegadinhas comuns

* WHERE não pode usar funções de janela (ROW_NUMBER(), RANK(), etc.). Para filtrá-las, use uma CTE ou subconsulta.
* HAVING filtra grupos; WHERE filtra linhas antes da agregação.
* COUNT(*) conta todas as linhas, enquanto COUNT(coluna) ignora valores NULL.
* UNION remove duplicatas; UNION ALL mantém todas as linhas.
* ROW_NUMBER() sempre gera valores únicos; RANK() e DENSE_RANK() tratam empates de forma diferente.
* NOT IN pode produzir resultados inesperados se a subconsulta retornar NULL; em muitos casos, prefira NOT EXISTS.
* Em consultas com LEFT JOIN, colocar condições da tabela da direita no WHERE pode transformar o resultado em um INNER JOIN. Quando necessário, mova essas condições para a cláusula ON.

Estratégia para a prova

* Resolva primeiro as questões de agregação (GROUP BY, HAVING) e JOIN, que costumam ser mais rápidas.
* Quando precisar filtrar por ranking (como “top 3 por categoria”), pense imediatamente em ROW_NUMBER() ou DENSE_RANK() dentro de uma CTE.
* Para comparar uma linha com a anterior ou calcular acumulados, use funções de janela (LAG(), LEAD(), SUM() OVER), evitando subconsultas complexas.
* Antes de escrever a consulta, identifique o padrão do problema: agregação, ranking, existência (EXISTS), registros sem correspondência (LEFT JOIN ... IS NULL), ou comparação com média/total. Isso acelera muito a resolução.

Se você dominar esses padrões e souber identificar rapidamente qual técnica aplicar, estará preparado para a grande maioria das questões da certificação SQL (Intermediate) do HackerRank.