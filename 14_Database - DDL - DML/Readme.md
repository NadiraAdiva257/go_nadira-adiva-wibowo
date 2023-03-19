Resume Materi "Database - DDL - DML"
- database adalah sekumpulan data yang tersimpan secara terstruktur
- ERD (Entity Relationship Diagram) adalah model yang memudahkan untuk membuat database yang yang datanya saling berhubungan dalam bentuk desain
- terdapat 4 komponen dalam ERD:
1. entitas adalah individu yang mewakili sesuatu yang nyata. entitas disini mewakili tabel-tabel dalam setiap database. setiap entitas mengandung instance
2. atribut adalah karakteristik dari entitas tersebut. atribut disini mewakili kolom-kolom dalam setiap tabel
3. UID (Unique Identifier) adalah atribut khusus yang unik sebagai penanda dari setiap tabel
4. relationship adalah hubungan antar 2 entitas atau lebih. terdapat 3 jenis relationship:
(1) 1 to 1 yang artinya setiap anggota entitas A hanya boleh berhubungan dengan 1 anggota entitas B, begitu juga sebaliknya. contohnya adalah mahasiswa dengan nim
(2) 1 to n yang artinya setiap anggota entitas A dapat berhubungan dengan lebih dari 1 anggota entitas B, tetapi tidak sebaliknya. contohnya adalah mahasiswa dengan buku
(3) m to n yang artinya setiap anggota entitas A dapat berhubungan dengan lebih dari 1 anggota entitas B, begitu juga sebaliknya. contohnya adalah mahasiswa dengan mata kuliah
- RDBMS (Relational Database Manajemen System) adalah aplikasi untuk memanajemen database. jenis RDBMS antara lain adalah mysql, postgree, oracle. bentuknya terstruktur seperti tabel di excel
- sql adalah salah satu bahasa untuk mengelola DBMS
- terdapat 4 jenis perintah sql:
1. DDL (Data Definition Language) adalah kumpulan operasi untuk memodifikasi struktur database, tabel, dan sebagainya
2. DML (Data Manipulation Language) adalah kumpulan operasi untuk memodifikasi data pada tabel
3. DCL (Data Control Language) adalah kumpulan operasi untuk menentukan siapa yang bisa mengakses
4. DTL (Data Transaction Language) adalah kumpulan operasi untuk memanajemen perubahan yang dilakukan DML
- beberapa operasi di DDL:
1. CREATE DATABASE namaDB --> membuat database
2. SHOW DATABASES --> melihat daftar database
3. USE namaDB --> masuk ke dalam database tersebut
4. CREATE TABLE namaTabel --> membuat tabel
5. ALTER TABLE namaTabel ADD atribut tipeData(panjangnya) --> menambahkan atribut atau kolom ke tabel
6. RENAME TABLE namaTabelLama to namaTabelBaru --> mengganti nama tabel
7. DROP TABLE namaTabel --> menghapus tabel
- beberapa operasi di DML:
1. INSERT INTO namaTabel (atribut) VALUES (isianAtribut) --> memasukkan data ke atribut atau kolom
2. SELECT * namaTabel --> menampilkan semua data atribut di tabel tersebut
3. SELECT * namaTabel WHERE --> menampilkan semua data atribut di tabel tersebut berdasarkan suatu kondisi
4. UPDATE namaTabel SET atribut = isianAtribut WHERE --> mengubah data atribut di tabel tersebut berdasarkan suatu kondisi
5. DELETE FROM namaTabel WHERE --> menghapus data atribut di tabel tersebut berdasarkan suatu kondisi
- beberapa statement di DML:
1. LIKE --> menampilkan data berdasarkan data string tertentu
2. BETWEEN --> menampilkan data pada rentang tertentu
3. IN --> menampilkan data yang ada di dalam IN
4. AND --> menampilkan data yang kedua kondisinya bernilai benar
5. OR --> menampilkan data yang salah satu kondisinya benar
6. ORDER BY --> menampilkan data sesuai urutan. ASC berarti urutan dari rendah ke tinggi atau kecil ke besar. DSC sebaliknya
7. DISTINCT --> menampilkan data yang unik
8. LIMIT --> menampilkan data sesuai banyak LIMIT
9. AS --> mengubah sementara nama tabel
10. GROUP BY --> mengelompokkan data untuk setiap atributnya
 