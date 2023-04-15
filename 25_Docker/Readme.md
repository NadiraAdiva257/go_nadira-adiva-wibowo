Resume Materi "Docker"
- docker adalah aplikasi untuk menyatukan berbagai file software dan pendukungnya dalam sebuah container untuk kemudian divirtualisasi di server supaya aplikasi tersebut bisa berjalan di berbagai platform dengan konfigurasi hardware yang berbeda-beda
- docker vs virtual machine:
1. docker --> sederhananya semua aplikasi dijalankan di atas 1 host operating system yang sama sehingga eksekusinya lebih cepat dan tidak memakan storage yang banyak
2. virtual machine --> sederhananya setiap aplikasi dijalankan di atas masing-masing operating system sehingga eksekusinya lebih lama dan memakan storage yang banyak
- terdapat beberapa istilah yang harus diketahui di docker:
1. docker image --> kumpulan file untuk membangun aplikasi. dimana 1 docker image dapat dimiliki banyak container
2. docker container --> wadah untuk mengemas dan menjalankan aplikasi. dimana hanya bisa mengakses resource yang telah ditentukan dari docker image dan 1 container hanya bisa memiliki 1 docker image
3. docker registry --> wadah untuk menyimpan docker image dengan perintah docker push. dimana docker registry berada di docker hub
4. docker hub --> layanan untuk menemukan dan berbagi docker image
5. docker daemon --> proses pengelolaan docker images, container, network, dan storage volumes. dimana docker daemon menerima request dari docker API dan akan memprosesnya
6. docker engine rest API --> untuk berinteraksi dengan docker daemon dan bisa diakses oleh client melalui HTTP
7. docker client --> pengguna dapat mengirimkan perintah seperti docker build, docker pull, dan docker run ke docker daemon
8. docker host --> bertanggung jawab menerima perintah yang diberikan docker client
- untuk infrastruktur docker antara lain adalah client mengirimkan perintah ke docker host. dimana di dalam docker host terdapat docker daemon yang mengelola docker images dan docker container. docker container nantinya akan menjalankan docker images. dan docker images akan disimpan di docker registry
- buat Dockerfile untuk membuat 1 docker images. terdapat beberapa syntax yang diperlukan dalam Dockerfile tersebut:
1. from --> pull image dari registry
2. workdir --> membuat direktori untuk instruksi semacam run, cmd, copy, dan add
3. copy --> copy file-file local ke workdir tersebut
4. run --> mengeksekusi bash command tersebut
5. add --> copy file dengan beberapa proses lainnya
6. expose --> menentukan port untuk container saat runtime
7. cmd --> eksekusi command tersebut saat container dijalankan
- sedangkan jika ingin membuat beberapa docker images secara langsung dalam 1 file, dapat membuat file docker-compose.yml
