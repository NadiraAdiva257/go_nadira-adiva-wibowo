Resume Materi "Clean and Hexagonal Architecture"
- arsitektur yang bagus adalah yang terpisahkan secara layer untuk membuat sebuah aplikasi yang modular, scalable, maintanable, dan testable
- MVC adalah model, view, controller. dimana user berinteraksi melalui view dan controller akan menghubungkan antara model dan view
- MVVM adalah model, view model, dan view
- ada permasalahan yang bisa muncul jika tidak menggunakan salah satu architecture, yaitu setiap tim yang biasanya mempunyai business domain dan unique code structure sendiri akan kesusahan untuk menjelaskan dan memahamkan kepada tim lain (mobility issue). terlebih lagi terkait dengan unit test yang berhubungan dengan database (another issue). selain itu, terlalu bergantung dengan framework juga tidak bagus
- clean architecture adalah salah satu macam arsitektur yang dapat menyelesaikan permasalan tersebut
- clean architecture dibagi menjadi 4 layer dari yang paling dalam:
https://medium.com/golangid/mencoba-golang-clean-architecture-c2462f355f41
1. entities / model layer --> menyimpan business rule yang dipakai pada domain lainnya. bisa berupa object dengan method atau struktur data. layer ini dapat diakses oleh semua layer dan semua domain. code di layer ini jarang berubah karena jika berubah, maka layer di luarnya akan terpengaruh juga. dan sebaliknya, perubahan di layer luar tidak boleh berpengaruh pada code di layer ini
2. repository layer --> menyimpan database handler (querying, inserting, deleting). tidak ada business logic karena yang ada hanya fungsi standar untuk input output dari datastore. selain itu, juga menentukan datastore yang digunakan (RDBMS atau NoSQL)
3. usecase layer --> menangani business logic (penjumlahan, pengurangan, dkk) pada setiap domain dan memilih repository layer yang digunakan. selain itu, juga penghubung antara repository layer dan delivery layer dan sebaliknya. datanya akan direpresentasikan dalam bentuk yang dapat dibaca oleh kedua layer. jika ada perubahan di repository layer, maka layer ini juga akan terpengaruh
4. delivery layer --> menjadi output dari aplikasi. menentukan metode penyampaian yang dipakai, misalnya dengan rest API, HTML, gRPC, file, dsb. selain itu, juga penghubung antara user dan sistem
- jadi intinya adalah jika ada perubahan di layer yang dalam, maka layer luarnya akan terpengaruh juga
- selain clean architecture, terdapat hexagonal architecture, onion architecture, screaming architecture, DCI from agile development, BCE from object oriented software, dkk. dimana memiliki tujuan yang sama yang memisahkan yang harus dipisahkan dengan cara membaginya menjadi beberapa layer. dan pasti memiliki layer untuk business rules, business logic, dan interfaces
- terdapat beberapa tujuan mengimplementasi clean architecture:
1. independent of frameworks --> memungkinkan untuk menggunakan framework sebagai alat
2. testable --> business rules dapat dites tanpa membutuhkan suatu hal yang lain
3. independent of UI --> user interface dapat diubah dengan mudah tanpa berpengaruh pada yang lainnya --> misal dari android ke IOS
4. independent of database --> business rules tidak terikat dengan database sehingga database dapat diubah dengan mudah --> misal dari MySQL ke MongoDB
5. independent of any external --> tidak ketergantungan dengan any external
- terdapat beberapa keuntungan menggunakan clean architecture:
1. karena terstruktur sehingga memudahkan untuk mencari yang dibutuhkan
2. perkembangannya lebih cepat dalam jangka waktu yang panjang
3. ketergantungan mocking menjadi hal yang biasa dalam unit testing
4. perubahan dapat dengan mudah dilakukan dari prototype ke solusi yang tepat. misalnya mengubah penyimpanan dari memori ke database SQL atau dari monolith ke microservices