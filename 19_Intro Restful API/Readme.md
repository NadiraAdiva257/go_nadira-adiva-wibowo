Resume Materi "Intro Restful API"
- API (Application Programming Interface) adalah software yang memiliki kumpulan fungsi dan prosedur yang memungkinkan pembuatan aplikasi yang dapat mengakses fitur atau data dari sistem operasi, aplikasi, dan layanan lain
- cara kerja API secara sederhana adalah menerima request dari client untuk diberikan kepada server. kemudian server akan memberikan respons ke client
- beberapa fungsi dari API:
1. memperluas fungsionalitas aplikasi
2. memudahkan integrasi multiplatform
3. menyediakan fleksibilitas
4. menyediakan akses terstandarisasi
- REST (Representational State Transfer) adalah gaya arsitektur yang memudahkan sistem berkomunikasi satu sama lain
- RESTful berarti sistem yang sifatnya stateless dan dapat memisahkan masalah client dan server serta memhamai pesan antar client dan server
- REST menggunakan protocol HTTP
- protocol HTTP sendiri adalah kumpulan perintah dalam komunikasi antar jaringan. contohnya adalah komputer client request mengakses IP address atau URL, lalu web server akan mengelola permintaan tersebut
- url HTTP sendiri antara lain terdiri dari schema (http atau https), authority (domain:port). terdapat juga parameter untuk memfilter, sorting, dkk
- terdapat beberapa method request HTTP:
1. GET untuk meminta data yang ditentukan
2. POST untuk mengirim data baru
3. PUT untuk mengupdate data baru yang ditentukan
4. DELETE untuk menghapus data yang ditentukan
- terdapat beberapa code response HTTP:
1. 200 artinya OK atau request berhasil
2. 400 artinya bad request
3. 404 artinya request not found
4. 405 artinya method not allowed terhadap request
5. 500 artinya internal server error
- format untuk request & response yaitu JSON, XML, SOAP, dkk
- JSON (JavaScript Object Notation) adalah format sederhana untuk menyimpan dan mentransfer data. dimana strukturnya adalah kumpulan value yang saling berpasangan (object) dan daftar value yang berurutan (array)
- beberapa API tools yang digunakan adalah postman, katalon, apache jmeter, soapUI, dkk
- postman sendiri adalah client HTTP yang memudahkan untuk menguji, mengembangkan, dan mendokumentasikan API dengan memungkinkan para user menggabungkan request HTTP yang sederhana dan kompleks dengan cepat
- terdapat beberapa aturan dalam pembuatan API:
1. gunakan kata benda, jangan kata kerja
2. gunakan kata plural, jangan singular
3. gunakan resource nesting untuk membuat hierarki
4. gunakan dash untuk pemisah string
5. gunakan lowercase
6. gunakan method crud yang disediakan HTTP sesuai dengan request yang spesifik
- dalam golang, pembuatan API dengan menggunakan package net/http. dimana di dalamnya dibuat routing dan controller secara manual