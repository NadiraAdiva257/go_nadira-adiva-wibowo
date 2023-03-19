Resume Materi "Join - Union - Agregasi - Subquery - Function (DBMS)"
- JOIN adalah klausa untuk mengombinasikan data dari dua atau lebih tabel
- INNER JOIN adalah klausa untuk mengombinasikan data dari dua atau lebih tabel yang memenuhi kondisi
- LEFT JOIN adalah klausa untuk mengombinasikan data dari dua atau lebih tabel yang akan mengembalikan seluruh data dari tabel kiri dan data yang memenuhi kondisi dari tabel kiri dan kanan
- RIGHT JOIN adalah klausa untuk mengombinasikan data dari dua atau lebih tabel yang akan mengembalikan seluruh data dari tabel kanan dan data yang memenuhi kondisi dari tabel kiri dan kanan
- UNION adalah klausa untuk menggabungkan semua data dari beberapa atribut pada beberapa tabel
- terdapat beberapa fungsi agregasi dari sql:
1. AVG --> mencari rata2 dari data di tabel
2. COUNT --> mencari jumlah dari data di tabel
3. MAX --> mencari data maksimal dari tabel
4. MIN --> mencari data minimum dari tabel
5. SUM --> mencari jumlah total nilai dari data di tabel 
6. HAVING --> mencari data berdasarkan suatu kondisi yang memiliki fungsi agregasi
- terdapat urutan penulisan operator dalam sql
SELECT ... FROM ... WHERE ... GROUP BY ... HAVING ... ORDER BY ...
- subquery adalah query di dalam query, dimana query yang berada di dalam akan dieksekusi terlebih dahulu
- subquery ditulis dalam tanda kurung dan posisinya di sebelah kanan pada suatu kondisi perbandingan
- jika hasil dari subquery hanya 1, maka operator perbandingan yang dapat digunakan antara lain adalah =, >, <, >=, <=, !=
- jika hasil dari subquery lebih dari 1, maka operator perbandingan yang dapat digunakan yaitu IN, ANY, ALL
- function dalam sql akan mengembalikan sebuah nilai ketika dipanggil
- syntax pembuatan function seperti berikut:
DELIMITER $
CREATE FUNCTION namaFunction RETURNS tipeData DETERMINISTIC
BEGIN
DECLARE varReturn
operasiSql
RETURN varReturn
END$$
DELIMITER ;
- deterministic menentukan siapa yang bisa menggunakan function tersebut. antara user pembuatnya saja (deterministic) atau user siapa saja (not deterministic)
- trigger function adalah function yang sebelum atau sesudah dijalankan akan memicu operasi lain
- syntax pembuatan trigger function seperti berikut:
DELIMITER $$
CREATE TRIGGER namaTrigger
BEFORE/AFTER operasiSql FOR EACH ROW
BEGIN
operasiSql
END$$
DELIMITER ;
- terdapat keyword OLD dan NEW dalam trigger. dimana OLD menyimpan nilai dari dalam data yang sedang berinteraksi atau dipanggil. sedangkan NEW menyimpan nilai dari data yang baru masuk atau sedang diinput