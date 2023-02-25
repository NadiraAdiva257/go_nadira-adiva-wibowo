Resume Materi "Control Version (Git)"
- versioning adalah pengaturan dan penyimpanan versi dari source code projek
- version control system adalah sistem yang membantu mengontrol versioning
- terdapat 3 perkembangan dari version control system:
1. single user adalah penyimpanan versi di local masing-masing sehingga sesama komputer belum saling menyambung. contohnya adalah RCS
2. centralized adalah penyimpanan versi hanya di 1 server pusatnya sehingga tidak ada yang tersimpan di local. contohnya adalah CVS
3. distributed adalah penyimpanan versi di local masing-masing dan 1 server yang sama sehingga saling tersinkronisasi. contohnya adalah GIT
- dari ketiga version control system tersebut yang paling populer adalah GIT, dimana GIT digunakan untuk mengembangkan software bersama-sama
- git repository adalah penyimpanan di local, yang berisi folder dan file projek serta folder .git
- remote repository adalah penyimpanan di github, yang berisi folder dan file projek juga
- install git terlebih dahulu di sistem operasi yang sesuai
- git init untuk membuat git repository di local
- git remote add untuk menambahkan remote repository ke local
- working directory adalah tempat untuk membuat, memodifikasi, dan menghapus projek
- staging adalah tempat penyimpanan sementara
- git add untuk menyimpan perubahan sementara dari working directory ke staging area
- git commit untuk menyimpan perubahan dari staging area ke repository dengan memberi pesan riwayat perubahan
- git status untuk mengecek status git repository di local, seperti pemberitahuan bahwa ada file yang belum di add, di commit atau ada file yang tertinggal dari remote repository
- git log --graph untuk melihat workflow perubahan
- git diff untuk melihat detail perubahan
- git stash untuk berpindah sementara ke staging area
- git stash apply untuk mengambil kembali dari staging area
- .gitignore untuk penyimpanan yang perubahannya diabaikan
- git checkout untuk berpindah ke branch lain
- git reset untuk kembali ke commit tertentu
- git push untuk mengirim perubahan yang telah di commit ke remote repository
- git fetch untuk mengambil metadata dari remote repository atau maksudnya mengambil perubahan dari remote repository tanpa di merge
- git pull untuk mengambil perubahan dari remote repository ke git repository dan langsung di merge
- git branch untuk melihat list branch
- git branch nama_branch untuk membuat branch baru
- git merge untuk menggabungkan antar branch
- pull request untuk mereview perubahan yang dibuat oleh branch lain sebelum dimerge ke branch main
- kegunaan pull request supaya menghindari conflict
- conflict adalah perubahan pada lokasi yang sama oleh 2 branch yang berbeda
- penyelesaian conflict dengan memilih code mana yang sebenarnya digunakan
- cara terbaik untuk berkolaborasi dengan git yaitu
1. buat 3 branch, yaitu branch main, develop, dan feature-feature
2. branch main hanya untuk menyimpan perubahan yang sudah fix banget. perubahan didapatkan dari branch develop
3. branch develop untuk menyimpan perubahan sebelum di merge ke branch main. perubahan didapatkan dari branch feature-feature
4. branch feature-feature untuk mengembangkan atau mengedit perubahan sesuai dengan feature masing-masing. perubahan tersebut akan di merge ke branch develop.