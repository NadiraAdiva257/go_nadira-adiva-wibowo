Resume Materi "Concurrent Programming"
- terdapat 3 jenis cara mengeksekusi programming:
1. sequential atau berurutan adalah mengeksekusi program secara satu per satu dan bergantian dengan 1 core
2. parallel adalah mengeksekusi program secara bersamaan dengan multicore. waktu selesainya akan serentak bersama
3. concurrent adalah mengeksekusi program secara hampir bersamaan tetapi independen dengan 1 core. waktu selesainya tidak bisa ditebak antara satu program dengan program lain
- tentang goroutine:
1. goroutine digunakan untuk menerapkan concurrent programming dalam golang dengan cara menambahkan keyword go ketika melakukan pemanggilan suatu fungsi
2. goroutine juga dapat digunakan secara multiple dan lebih ringan daripada thread
3. function main yang biasanya digunakan untuk menjalankan berbagai code program di golang juga merupakan goroutine sehingga butuh blocking main untuk mencegah function main selesai duluan sebelum goroutine lainnya dijalankan. salah satu caranya dengan membuat time.Sleep()
4. terdapat juga gomaxprocs untuk mengontrol mau menggunakan berapa banyak procs sehingga sisanya dapat digunakan untuk database, restAPI, dkk
- tentang channel:
1. channel digunakan sebagai jembatan komunikasi antar goroutine
2. unbuffered channel adalah channel yang tidak memiliki size. syntaxnya adalah make(chan tipeData). prosesnya adalah data dikirim kemudian ditangkap untuk dikeluarkan. proses tersebut dilakukan secara bergantian sehingga hanya dapat menampung 1 isian
3. bisa terjadi deadlock pada unbuffered channel jika channel tidak ada pengirimnya dan tidak ada penangkapnya
4. buffered channel adalah channel yang memiliki size. syntaxnya adalah make(chan tipeData, n). prosesnya adalah data dikirim kemudian ditampung dulu sesuai size baru kemudian ditangkap untuk dikeluarkan. nanti kalau yang ditampung ternyata melebihi size, maka akan ditangkap beberapa. prosesnya mirip seperti FIFO. tidak memerlukan goroutine untuk mengirim dan menangkap data
5. close channel digunakan supaya tidak menerima data lagi dengan syarat ditaruh di pengirim. bisa dipakai sebagai alternatif mengeluarkan channel secara dinamis
- tentang select:
1. select digunakan untuk mengontrol channel2 yang saling berkomunikasi
2. dengan select, maka data dari channel2 tersebut bisa untuk diproses lagi jika memang diinginkan
- tentang race condition:
1. race condition adalah kelemahan dari goroutine. terjadi ketika 2 atau lebih goroutine mengakses memori yang sama, dimana pada saat bersamaan terdapat memori yang sedang diassign atau dimodif
2. race condition tidak menyebabkan error tetapi output bisa tidak sesuai yang diharapkan
3. go run --race file.go untuk mengetahui apakah ada race condition atau tidak
4. terdapat 3 cara mengatasi race condition yaitu wait groups, channel blocking, dan mutex. ketiganya memiliki prinsip yang sama yaitu yang menyebabkan race condition dilock dulu (biasanya proses write) sampai prosesnya selesai baru nanti diunlock untuk ke proses selanjutnya (biasanya proses read)
- tentang wait groups
1. wg.Add() menambahkan berapa banyak proses goroutine yang akan dijalankan
2. wg.Done() memberi sinyal ke wg.Wait() kalau proses goroutine sudah selesai dijalankan
3. wg.Wait() menahan untuk lanjut ke proses di luar goroutine sampai wg.Done() memberi sinyal
- tentang channel blocking
1. proses sederhananya adalah akan berusaha menyinkronisasikan antar goroutine
2. biasanya deklar var dulu dengan make(chan bool)
3. var <- true, fungsinya sama seperti wg.Done()
4. <- var, fungsinya sama seperti wg.Wait()
- tentang mutex
1. deklarasi var dulu dengan sync.Mutex
2. var.Lock(), fungsinya sama seperti wg.Done()
3. var.Unlock(), fungsinya sama seperti wg.Wait()