Resume Materi "System Design"
- terdapat 5 hal yang dibahas dalam system design:
1. diagram design
- digunakan untuk memudahkan dalam membuat system design
- diagram sendiri adalah simbol untuk merepresentasikan informasi yang akan disampaikan dalam system
- tools diagram antara lain adalah draw io, lucid chart, dkk
- terdapat 4 jenis diagram:
(1) flowchart adalah representasi 1 alur system dengan menunjukkan prosesnya
(2) use case diagram adalah rangkuman dari system dan interaksi antar aktor dengan system
(3) entity relationship diagram adalah design dalam pembuatan database
(4) high level architecture adalah keseluruhan dari system design (arsitektur) atau tingkatan diagram yang paling tinggi dalam system design
2. distributed system
- distributed secara sederhananya dapat digambarkan dalam ystem yang sudah besar atau kompleks. dimana biasanya membutuhkan banyak server supaya kinerjanya powerfull karena eksekusi tugasnya dapat dibagi
- terdapat 5 karakteristik:
1. scalability
- misalnya adalah system yang biasanya menerima 25 item/day sedangkan servernya hanya mengeksekusi 20 item/day --> maka servernya dapat di incerease ability nya sehingga server dapat mengeksekusi 30 item/day
- misalnya lagi, kemudian system menerima 50 item/day --> maka server yang mengeksekusi 30 item/day dapat diduplikasi
- vertical scale adalah menaikkan kapasitas server
- horizontal scale adalah menduplikasi server
2. reliability adalah kehandalan system untuk memberikan pelayanan sehingga solusinya adalah distributed system (multi server)
3. availability
4. efficiency
- salah satu contohnya adalah efisiensi dalam waiting respond. dimana maksimal 2 detik
5. serviceability / manageability adalah kondisi dimana system harus bisa dimaintenance & diservice sehingga clean code dibutuhkan untuk memudahkan hal tersebut
3. job queue dan services
- job queue adalah antrian task yang dijalankan secara async
- load balancing adalah pembagian traffic supaya ideal sesuai inputan real time. contohnya digunakan antar user dengan web server, web server dengan internal platform layer (apk server atau cache server), internal platform layer dengan database, dkk
- monolithic adalah servis system yang dijalankan dalam 1 server. kurang aman kalau untuk system yang sudah besar
- microservices adalah servis system yang dijalankan dalam banyak server sehingga eksekusinya lebih ringan. keuntungannya lagi jika salah satu servis sedang tidak bisa dijalankan, servis lain tetap bisa dijalankan dan diakses oleh user
4. SQL dan NO SQL
- sql adalah bahasa untuk RDBMS yang terstruktur, seperti mysql, postgree, sql lite
- no sql adalah bahasa untuk DBMS yang dinamis sehingga lebih fleksibel, seperti redis (key & value), cassandra (column-family), neo4j (graf)
- manfaat RDBMS adalah bisa untuk segala keperluan fitur, standarnya jelas, punya byk tool (administrasi, reporting, framework)
- terdapat 4 prinsip RDBMS:
1. atomicity adalah transaksi terjadi semua atau tidak sama sekali
2. consistency berarti databasenya sesuai dengan yang sudah ditentukan karena aturannya jelas
3. isolation adalah transaksi dieksekusi secara sekuensial walaupun request bersamaan (concurrent)
4. durability adalah jaminan bahwa transaksi yg dihapus akan tetap tersimpan di database dengan soft delete
- no sql dapat digunakan ketika menghindari prinsip ACID, menghindari kompleksitas sql karena join yang terlalu banyak, menghindari design schema di depan, butuh schema yang fleksibel, data logging (json), data sementara (cache), dkk
- kelebihan no sql adalah schema less, fast dev, dkk
5. caching dan indexing
- cache adalah penyimpanan data sementara sehingga jika ada request yang sudah tersimpan di cache, maka bisa menggunakan itu lagi sehingga prosesnya tidak perlu diulang
- cache disimpan di RAM
- database replication adalah cadangan db. distributed tetapi tetap saling terhubung. mendukung reliability di distributed system sehingga tiap server juga punya database sendiri
- database indexing --> database yang diurutkan / sorting berdasarkan indexing yang diinginkan sehingga lebih powerful dalam select, insert, delete. complexitynya adalah O log(n)