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
insert into tbl (k, v) values (11, '11string');
insert into tbl (k, v) values (11, '11string'), (2, '2str'), (3, '3str');
update tbl set k = 22 where k = 11;
delete from tbl where k = 22;
select * from tbl;
update tbl set v = '44str' where k = 44 returning *;
table tbl;
```


# ShortCut
## Goland
`Double ⇧Shift` - Search everywhere
