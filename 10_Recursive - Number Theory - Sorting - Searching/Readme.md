Resume Materi "Recursive, Number Theory, Sorting, Searching"
- tentang recursive
1. recursive adalah function yang memanggil dirinya sendiri sehingga prosesnya seperti stuck
2. recursive memudahkan untuk menyelesaikan masalah dan memperpendek penulisan code
3. akan tetapi, jika recursive yang dipanggil terlalu banyak dapat menyebabkan stuck overflow
4. contoh penggunaan recursive adalah factorial dan fibonacci
- tentang number theory
1. number theory lebih berisi hal-hal yang berkaitan dengan matematika
2. contoh dari number theory adalah bilangan prima, FPB, KPK, faktorial, dkk
- tentang searching
1. searching adalah proses pencarian yang diinginkan dari kumpulan banyaknya data
2. terdapat beberapa teknik searching, yaitu linear search, binary search, dkk
3. untuk linear search berarti mencari yang diinginkan dengan cara mengecek satu per satu kumpulan datanya sehingga time complexity nya adalah O(n) dan space complexity nya adalah O(1)
4. untuk binary search berarti mencari yang diinginkan dengan cara membagi dua kumpulan data yang sudah terurut. misal yang ingin dicari ternyata ada di sebelah kanan setelah kumpulan datanya dibagi 2, berarti fokus pencariannya ada di sebelah kanan. begitu seterusnya sampai akhir. time complexity nya adalah O(log n) dan space complexity nya adalah O(1)
- tentang sorting
1. sorting adalah pengurutan kumpulan data yang sebelumnya acak sehingga menjadi terurut
2. terdapat beberapa teknik sorting, yaitu selection sort, counting sort, merge sort, insertion sort, bubble sort, quick sort, dkk
3. untuk selection sort berarti membandingkan antar nilai secara urut. jika ketemu angka yang lebih kecil, maka posisinya akan ditukar. time complexity nya adalah O(n^2) dan space comlexity nya adalah O(1)
4. untuk counting sort berarti membuat array penampung untuk menghitung ada berapa kali angka pada data yang muncul untuk ditulis sesuai dengan posisi index pada umumnya. nantinya array tersebut hanya sepanjang index terakhir yang ada. baru kemudian diurutkan dari yang terkecil. time complexity nya adalah O(n+k) dan space complexity nya adalah O(k)
5. untuk merge sort berarti membagi kumpulan data menjadi dua sampai data berdiri sendiri. baru kemudian digabungkan secara bertahap dengan tipa penggabungannya diurutkan. time comlexity nya adalah O(log n) dan space complexity nya adalah O(1)