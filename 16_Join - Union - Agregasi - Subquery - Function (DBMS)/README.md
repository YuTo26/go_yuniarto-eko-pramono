# (16) Join - Union - Agregasi - Subquery - Function (DBMS)

*******************
1/ Join => Join digunakan untuk menggabungkan dua tabel atau lebih berdasarkan kolom yang sama di antara mereka. Berikut adalah jenis join yang umum digunakan:

- Inner Join: mengembalikan baris yang memiliki nilai yang cocok di kedua tabel.
- Left Join: mengembalikan semua baris dari tabel kiri dan baris yang cocok dari tabel kanan.
- Right Join: mengembalikan semua baris dari tabel kanan dan baris yang cocok dari tabel kiri.
- Full Outer Join: mengembalikan semua baris dari kedua tabel dan mengisi nilai NULL jika tidak ada nilai yang cocok di antara mereka.
Sintaks dasar join:

```sql
SELECT column_name(s)
FROM table1
JOIN table2
ON table1.column_name = table2.column_name;
```

*******************
2/ Union => Union digunakan untuk menggabungkan hasil dari dua atau lebih SELECT statements ke dalam satu set hasil yang tidak memiliki duplikat. Setiap SELECT statement harus memiliki jumlah dan tipe data kolom yang sama.
Sintaks dasar union:

```sql

SELECT column_name(s) FROM table1
UNION
SELECT column_name(s) FROM table2;
```

*******************
3/ Agregasi => Agregasi adalah operasi yang digunakan untuk menghitung nilai statistik pada satu atau lebih kolom dari tabel, seperti COUNT, SUM, AVG, MIN, dan MAX. Seringkali digunakan dengan GROUP BY untuk menghitung nilai agregat berdasarkan kelompok.
Sintaks dasar agregasi:

```sql
SELECT COUNT(column_name)
FROM table_name
WHERE condition;
```

*******************
4/ Subquery => Subquery adalah query yang tertanam di dalam query utama dan digunakan untuk mendapatkan informasi dari tabel lain. Subquery dapat digunakan di SELECT, FROM, WHERE, dan HAVING statements.
Sintaks dasar subquery:

```sql

SELECT column1, column2, ...
FROM table1
WHERE column_name IN (SELECT column_name FROM table2 WHERE condition);
```

*******************
5/ Function (DBMS) => Function adalah sebuah objek database yang digunakan untuk menghasilkan nilai-nilai kembalian berdasarkan input yang diberikan. Function dapat digunakan dalam operasi SQL seperti SELECT, WHERE, dan JOIN.

```sql

CREATE FUNCTION function_name(parameter1 data_type, parameter2 data_type, ...)
RETURNS return_type
BEGIN
SQL statements;
END;

SELECT function_name(parameter1, parameter2, ...);
```

*******************
