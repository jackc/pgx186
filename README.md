# pgx issue 186 test

```
$ createdb pgx186
$ psql -f structure.sql pgx186
Expanded display is used automatically.
Timing is on.
Null display is "∅".
CREATE TABLE
Time: 16.140 ms
$ PGHOST=/tmp PGDATABASE=pgx186 go run main.go
$ psql -c 'select * from projects' pgx186
     id     |  user_id   |                                                                                  data                                                                                  |          created_at
------------+------------+------------------------------------------------------------------------------------------------------------------------------------------------------------------------+-------------------------------
 84679d5109 | 28e20760c7 | {"id": "", "name": "", "site": "", "type": 1, "email": "", "name2": "", "tasks": [], "device": 1, "status": 1, "created_at": "0001-01-01T00:00:00Z", "browser_bar": 0} | 2016-09-27 08:39:38.827125-05
(1 row)
```
