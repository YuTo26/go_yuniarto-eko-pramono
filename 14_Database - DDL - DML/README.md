# (14) Database Schema - DDL - DML

*******************

1/ Database schema merupakan struktur yang digunakan untuk merepresentasikan hubungan antara tabel dalam sebuah database. Dalam database schema, terdapat deskripsi dari tabel, kolom, relasi antar tabel, dan constraint pada tabel.

*******************

2/ DDL (Data Definition Language) digunakan untuk mengatur struktur tabel dan database seperti membuat, mengubah, dan menghapus tabel, kolom, atau constraint. Contoh sintaksis DDL adalah:

- CREATE TABLE untuk membuat tabel baru
- ALTER TABLE untuk mengubah tabel yang sudah ada
- DROP TABLE untuk menghapus tabel dari database
  
*******************

3/ DML (Data Manipulation Language) digunakan untuk mengelola data dalam tabel seperti menambah, mengubah, dan menghapus data dalam tabel. Contoh sintaksis DML adalah:

- SELECT untuk menampilkan data dari tabel
- INSERT untuk menambahkan data ke dalam tabel
- UPDATE untuk mengubah data dalam tabel
- DELETE untuk menghapus data dari tabel

5 / contoh sintak basic DDL dan DML

DDL:

``` sql

CREATE TABLE users (
  id INT PRIMARY KEY,
  name VARCHAR(50),
  email VARCHAR(100) UNIQUE,
  password VARCHAR(100)
);

ALTER TABLE users
ADD COLUMN age INT;

DROP TABLE users;
```

DML:

```sql

SELECT * FROM users;

INSERT INTO users (id, name, email, password)
VALUES (1, 'John Doe', 'john@example.com', 'password123');

UPDATE users
SET name = 'Jane Doe'
WHERE id = 1;

DELETE FROM users
WHERE id = 1;
```
