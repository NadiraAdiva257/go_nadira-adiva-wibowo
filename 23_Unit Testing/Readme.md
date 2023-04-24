Resume Materi "Unit Testing"
- testing adalah proses analisis software untuk mendeteksi perbedaan fitur yang diharapkan dengan fitur yang sudah dibuat serta memastikan apakah software berjalan dengan baik tanpa adanya error atau bug
- terdapat beberapa tujuan dari testing:
1. mencegah regresi --> dimana kondisi software sebelumnya sudah berjalan dengan baik, tapi kemudian ada perubahan sehingga software menjadi eror dan perlu ditesting ulang
2. refactoring --> merubah code tanpa mengubah fungsionalitasnya
3. meningkatkan code design --> membangun code yang bisa ditesting dengan membuat fungsi-fungsi kecil
4. code documentation --> memudahkan untuk mengetahui input, proses, dan output sistem tanpa perlu melihat detail code nya
5. memperluas team --> memudahkan tim lain untuk memahami dengan testing yang berhasil
- terdapat 3 level dalam testing:
1. UI --> testing keseluruhan dengan interaksi secara langsung terhadap user interfacenya
2. integrarion --> testing spesifik modul atau subsistem dengan melalui API
3. unit --> testing unit terkecil atau terspesifik dengan melalui fungsi-fungsi atau method-method yang berisi logical operation
- untuk testing di golang sendiri sudah disediakan. dan untuk tambahan frameworknya dapat menggunakan framework testify
- terdapat 2 pola testing yang biasanya digunakan:
1. centralize --> semua file testing dijadikan dalam 1 folder tersendiri
2. together --> file testing disimpan dalam setiap folder yang di dalamnya juga terdapat file yang mau ditesting
- test file adalah koleksi dari test suites
- test suites adalah koleksi dari test cases
- test fixtures adalah proses untuk memastikan environment yang dilakukan untuk testing tetap konsisten
- test cases adalah kumpulan code uji positif dan negatif untuk menjalankan testing dari setiap fungsi dengan memiliki input, proses, dan output
- mocking adalah semacam duplikasi dari suatu fungsi yang simulasi perilakunya mirip dengan yang asli sehingga testing lebih independen sehingga tingkat resikonya lebih rendah
- beberapa hal yang tidak perlu ditesting yaitu 3rd party modules, database, 3rd party API, object class yang sudah ditesting sebelumnya
- coverage adalah alat ukur untuk menunjukkan source code program testing sudah dieksekusi dan memastikan semua logical operation sudah ditesting
- coverage akan muncul saat testing dijalankan
- beberapa format coverage report adalah CLI, XML, HTML, dan LCOV
- untuk membuat sebuah test file, penamaan di akhir harus bertuliskan xxx_test.go
- untuk membuat sebuah test function, penamaan di awal harus bertuliskan TestXxx(t *testing.T)
- untuk menjalankan testing, dapat menggunakan command go test .pathFolder -cover
- u/ menjalankan testing dengan coverage, dapat menggunakan command go test .pathFolder -cover - versiReport=cover.out
