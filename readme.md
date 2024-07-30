# Содержание
- [docker](#docker)
- [Postgres SQL](#postgres-sql)
- [Shortcut](#shortcut)
  - [goland](#goland)


# DOCKER
[tutorial on youtube docker](https://www.youtube.com/watch?v=3c-iBn73dDE)

создать контейнер с базой данных
```shell
docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
```
```shell
docker run -d -p 3306:3306 --name www -e MYSQL_ROOT_PASSWORD=417149 mysql
```
docker Status
```shell
docker ps
```
Export BD SQL
```shell
mysqldump -u root -p www > www.sql
```
coonect to docker containers
```shell
docker exec -it 81a871529576 /bin/bash
```
connect to psql
```shell
psql -U postgres
```


# Postgres SQL

```sql
--порядок использования
FROM > WHERE > GROUP BY > HAVING > ORDER BY > DISTINCT > OFFSET > LIMIT > FETCH
```

```sql
insert into tbl (k, v) values (11, '11string');
insert into tbl (k, v) values (11, '11string'), (2, '2str'), (3, '3str');
update tbl set k = 22 where k = 11;
delete from tbl where k = 22;
select * from tbl;
update tbl set v = '44str' where k = 44 returning *;
table tbl;
```

```sql
--Сортировка по нескольким столбцам:
SELECT column1, column2 FROM table1 ORDER BY column2 DESC, column1 ASC;
```

```sql
SELECT * FROM employees limit 3 offset 3;
id | name  |       position       | department_id
----+-------+----------------------+---------------
4 | Diana | Marketing Specialist |             4
5 | Eve   | Accountant           |             5
6 | Frank | IT Support           |             6
(3 rows)
```




```sql
-- выборкой без учета регистра ILIKE
SELECT * FROM employees WHERE name ILIKE 'eve';
 id | name |  position  | department_id 
----+------+------------+---------------
  5 | Eve  | Accountant |             5
(1 row)
```

```sql
--Подсчет количества сотрудников в каждом департаменте
SELECT department_id, COUNT(*) as employee_count
FROM employees
GROUP BY department_id
ORDER BY department_id;

department_id | employee_count 
---------------+----------------
             1 |              1
             2 |              1
             3 |              1
             4 |              1
             5 |              1
             6 |              1
             7 |              1
             8 |              1
             9 |              1
            10 |              1
(10 rows)
```








# ShortCut
## Goland
`Double ⇧Shift` - Search everywhere


# ffmpeg
Вытянуть субтитры
```shell
ffmpeg -i MovieName.mkv -map 0:s:0 subs.srt
```