Resume Materi "Configuration Management and CLI"
- CLI (Command Line Interface) adalah baris perintah untuk menjalankan suatu tugas pada komputer dan kecepatannya lebih cepat daripada GUI
- terdapat 4 alasan memakai command line:
1. kontrol granular dari aplikasi atau OS
2. manajemen lebih cepat untuk sistem operasi yang jumlahnya besar
3. mampu untuk membuat dan menyimpan script yang berisikan baris-baris perintah
4. membantu menyelesaikan masalah. contohnya adalah masalah koneksi jaringan
- shell atau CLI adalah tempat untuk menjalankan command line. contohnya adalah command prompt di Windows, terminal di MacOS
- command line yang akan saya resume di sini adalah command di UNIX shell
- saat membuka UNIX shell, akan muncul tulisan berupa your_username@your_hostname:~$ atau :~#
- ~ adalah current path (home)
- $ adalah shell untuk normal user yang hanya dapat memanipulasi /home/username directory
- # adalah shell untuk root user yang dapat memanipulasi semua directory
- normal user + sudoers dapat memanipulasi seperti halnya root user
- beberapa command yang terkait directory:
1. pwd --> menampilkan directory yang aktif
2. ls --> menampilkan isian direktori
3. ls -l --> menampilkan isian direktori secara lengkap. terdiri dari permission, count of file, owned by (user), owned by (group), size, create date, file/directory name
4. mkdir --> membuat directory baru
5. cd --> berpindah directory
6. rmdir --> menghapus directory yang masih kosong
7. rm -r --> menghapus directory beserta semua isian di dalamnya
8. cp --> mengcopy directory ke directory lain
9. mv --> memindahkan directory ke directory lain
10. ln --> membuat file yang akan mengelink ke file asli. ada soft link dan hard link. dimana soft link kalau file aslinya dihapus, maka file tersebut tidak akan bisa dibuka. sedangkan hard link kebalikannya
11. chmod --> merubah permission directory. penjelasan ada di chmod file
- beberapa command yang terkait file:
1. touch --> membuat file baru
2. head, cat, tail --> menampilkan isian dari file. dimana cat akan menampilkan keseluruhannya. sedangkan head hanya akan menampilkan isian pada barisan awalan dan tail menampilkan isian pada barisan akhiran
3. vim, nano --> menambah, memodifikasi, menghapus isian file
4. chown --> merubah owner file
5. chmod --> merubah permission file. dimana dibagi menjadi 3 owner, yaitu user, group, dan other. di setiap owner tersebut bisa membaca (r), menulis (w), dan mengeksekusi (x). untuk mengaktifkan rwx, maka chmod nya harus 7 dalam desimal. untuk r, chmod nya bernilai 4 dalam desimal. untuk w, chmod nya bernilai 2 dalam desimal. untuk x, chmod nya bernilai 1
6. diff --> membandigkan isian antar file
7. rm --> menghapus file dan isiannya
8. cp --> mengcopy file ke directory lain
9. mv -> memindahkan file ke directory lain
- beberapa command yang terkait network:
1. ping --> mengetes koneksi ke sebuah target apakah ada internetnya atau tidak
2. ssh --> media transport untuk koneksi ke remote server
3. netstat --> mengetahui aktivitas di dalam jaringan tersebut
4. nmap --> menganalisis port
5. ifconfig
6. wget, curl --> mengunduh sebuah file dari url yang ditentukan
7. telnet
8. netcat
- beberapa command yang terkait utility:
1. man --> mengetahui informasi terkait suatu command
2. env
3. echo --> menampilkan sesuatu yang ditulis. echo juga bisa sekalian memasukkan tulisanya ke dalam file dengan echo ... > namaFile. untuk ke baris selanjutnya di dalam file, maka echo ... >> namaFile
4. date --> mengetahui real time
5. which
6. watch --> monitoring sebuah program
7. sudo --> bertindak seperti root user
8. history --> menampilkan semua command yang pernah dilakukan
9. grep --> mendapatkan file yang mengandung suatu data
10. locate --> melakukan pencarian
- shell script adalah suatu script yang menyimpan berbagai command sehingga ketika script tersebut dijalankan maka otomatis menjalankan berbagai command di dalamnya
- permission untuk semua ownernya harus bisa mengeksekusi
- ./namaScript.sh --> enjalankan script tersebut
- vim atau nano namaScript.sh --> mengisi command-command di dalamnya