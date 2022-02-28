# Data

## Mysql Database Connection Source
```text
root:123456@tcp(127.0.0.1:3306)/kratos_shop?charset=utf8mb4&parseTime=True&loc=Local
```

## 创建数据库
```mysql
mysql> CREATE DATABASE kratos_shop;
Query OK, 1 row affected (0.01 sec)

```

## 创建数据库后，启动程序 ent 会自动创建数据库表(太牛逼了)

```text
mysql> CREATE DATABASE kratos_shop;
Query OK, 1 row affected (0.01 sec)

mysql> show tables;
+-----------------------+
| Tables_in_kratos_shop |
+-----------------------+
| addresses             |
| cards                 |
| users                 |
+-----------------------+
3 rows in set (0.00 sec)

mysql> desc users;
+---------------+--------------+------+-----+---------+----------------+
| Field         | Type         | Null | Key | Default | Extra          |
+---------------+--------------+------+-----+---------+----------------+
| id            | bigint(20)   | NO   | PRI | NULL    | auto_increment |
| username      | varchar(191) | NO   | UNI | NULL    |                |
| password_hash | varchar(255) | NO   |     | NULL    |                |
| created_at    | datetime     | NO   |     | NULL    |                |
| updated_at    | datetime     | NO   |     | NULL    |                |
+---------------+--------------+------+-----+---------+----------------+
5 rows in set (0.01 sec)

mysql> desc cards;
+------------+--------------+------+-----+---------+----------------+
| Field      | Type         | Null | Key | Default | Extra          |
+------------+--------------+------+-----+---------+----------------+
| id         | bigint(20)   | NO   | PRI | NULL    | auto_increment |
| name       | varchar(255) | NO   |     | NULL    |                |
| card_no    | varchar(255) | NO   |     | NULL    |                |
| ccv        | varchar(255) | NO   |     | NULL    |                |
| expires    | varchar(255) | NO   |     | NULL    |                |
| created_at | datetime     | NO   |     | NULL    |                |
| updated_at | datetime     | NO   |     | NULL    |                |
| user_cards | bigint(20)   | YES  |     | NULL    |                |
+------------+--------------+------+-----+---------+----------------+
8 rows in set (0.00 sec)

mysql> desc addresses;
+----------------+--------------+------+-----+---------+----------------+
| Field          | Type         | Null | Key | Default | Extra          |
+----------------+--------------+------+-----+---------+----------------+
| id             | bigint(20)   | NO   | PRI | NULL    | auto_increment |
| name           | varchar(255) | NO   |     | NULL    |                |
| mobile         | varchar(255) | NO   |     | NULL    |                |
| address        | varchar(255) | NO   |     | NULL    |                |
| post_code      | varchar(255) | NO   |     | NULL    |                |
| created_at     | datetime     | NO   |     | NULL    |                |
| updated_at     | datetime     | NO   |     | NULL    |                |
| user_addresses | bigint(20)   | YES  |     | NULL    |                |
+----------------+--------------+------+-----+---------+----------------+
8 rows in set (0.00 sec)

```