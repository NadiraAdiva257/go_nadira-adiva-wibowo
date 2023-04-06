Resume Materi "ORM and Code Structure (MVC)"
- notes: path orm-mvc/prioritas1 isinya jawaban dari prioritas1 sekaligus prioritas2
- ORM (object relational mapping) adalah teknik pemrograman di komputer untuk mengkonversi data yang ada pada suatu relational database menjadi sebuah object di memory atau sebaliknya supaya mudah diakses dan dimanipulasi dengan menggunakan bahasa pemrograman berorientasi objek. Dimana keduanya akan mempunyai property yang sama persis
- dalam golang, ORM bisa digunakan pada library GORM
- terdapat beberapa kelebihan menggunakan ORM:
1. query lebih singkat --> query yang panjang bisa disingkat menjadi lebih ringkas
2. otomatis fetch data ke use object -> data dari relational database akan otomatis dikonversi ke object struct di memory. tanpa perlu mapping 1 per 1
3. screening data sebelum disimpan di database --> mudah untuk memvalidasi data sebelum dimasukkan ke relational database
4. punya fitur cache query --> mengeksekusi lebih cepat suatu query yang sebelumnya pernah dipanggil karena tersimpan di cache
- terdapat beberapa kekurangan mengunakan ORM:
1. biaya prosesnya cukup overhead
2. mengeksekusi data atau query yang tidak dibutuhkan atau diinginkan
3. jika terdapat query yang complex, biasanya saat lebih dari 10 tabel yang join, maka eksekusinya akan sangat lama
4. terdapat beberapa sql khusus yang tidak terdapat di ORM sehingga harus menggunakan library lain
- database migration adalah cara untuk memperbarui skema database atau mengembalikan skema database ke skema sebelumnya
- terdapat beberapa faktor kenapa database migration:
1. update database simpel
2. rollback database simpel
3. track perubahan pada struktur database
4. history struktur database ditulis dalam bentuk code
5. selalu kompatibel dengan versi aplikasi berapapun
- untuk menggunakan GORM, maka harus menginstall GORM dan driver databasenya (MySQL)
$ go get -u gorm.io/gorm
$ go get -u gorm.io/driver/mysql
 - untuk terkoneksi dengan database, maka dapat menggunakan syntax <username>:<password>@/<db_name>?charset=utf8&parseTime=True&loc=Local
- buat juga skema yang sama dengan yang ada di database menggunakan struct dan buat auto migration untuk skema tersebut
- untuk melakukan CRUD dan segala operasional database bisa diakses melalui gorm.io/docs. beberapa syntax diantaranya adalah DB.Save untuk memasukkan data, DB.First() untuk mengambil data, DB.Updates() untuk mengupdate data, DB.Delete() untuk menghapus data
- untuk MVC sendiri adalah struktur code yang mengorganisasi dan memisahkan antara model, view, dan controller
- model adalah code yang mengorganisasikan data yang ada di database
- view adalah code untuk menampilkan data yang telah diolah tersebut
- controller adalah adalah code yang menghubungkan dan mengatur antara model dan view. dimana biasanya berisi fungsi-fungsi untuk mengolah data tersebut
- MVC dibutuhkan untuk mencapai aplikasi modular, menerapkan pemisahan kepentingan, dan mengurangi konflik saat pembuatan

