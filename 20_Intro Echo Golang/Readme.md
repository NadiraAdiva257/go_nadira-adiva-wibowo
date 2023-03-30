Resume Materi "Intro Echo Golang"
- echo golang adalah framework bahasa golang yang membantu untuk mengembangkan aplikasi berbasis web
- echo golang sendiri memiliki high performance, extensible, dan minimalist untuk dilakukan
- kekurangan dari menggunakan framework adalah menyebabkan ketergantungan. dimana ketika library dalam framework diupdate, maka projeknya juga harus diupdate
- terdapat beberapa keuntungan dari menggunakan echo golang adalah:
1. optimized router, karena menggunakan lokasi dynamic memori yang nilainya adalah 0
2. data rendering, berarti untuk memberi response dari suatu request akan cukup mudah dengan menggunakan custom handler
3. data binding, berarti dapat melakukan request ke API dalam bentuk apapun (form, JSON, query, dkk) sehingga memudahkan untuk mendapatkan data tersebut
4. middleware, berarti sebuah blok kode yang dipanggil sebelum atau sesudah http request diproses. middleware biasanya dibuat per fungsi nya, contohnya: middleware autentikasi, middleware untuk logging, middleware untuk gzip compression, dkk
5. scalable, berarti projek mudah dikembangkan dalam rest API yang berskala kecil sampai besar
- jadi echo golang minimalist but extensible. dimana minimalist karena walaupun sederhana tapi powerfull. kemudian dia extensible karena dapat collab dengan library lain, contohnya dengan gorm untuk operasi databasenya, dengan go jwt untuk authentication, dkk
- untuk menggunakan echo golang, maka perlu menginstall framework tersebut melalui command go get github.com/labstack/echo/v4
- pada func main di file main.go, deklarasi variable echo dengan e := echo.New()
- buat route dengan yang diinginkan sesuai dengan method request HTTP yang diinginkan juga. misal e.GET("/students", funcController/funcHandler)
- buat juga syntax untuk mulai menjalankan server dengan e.Start(":portYangDiinginkan")
- handler pada echo golang bertugas menerima request yang masuk. dimana selanjutnya akan diolah pada controller
- data rendering adalah merespons dalam bentuk JSON dari suatu request
- terdapat beberapa bentuk retrieve data:
1. URL params, dengan syntax Param("keyYangDiinginkan") dan route dengan "/:keyYangDiinginkan"
2. query params, dengan syntax QueryParam("keyYangDiinginkan") dan route dengan "/?keyYangDiinginkan=value"
3. form value, dengan menuliskan form:"keyYangDiinginkan" di struct. kemudian dengan syntax FormValue("keyYangDiinginkan")
- untuk testing keseluruhan API tersebut dapat menggunakan aplikasi Postman
