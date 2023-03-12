Resume Materi "Clean Code"
- clean code adalah penulisan kode yang baik sehingga mudah dibaca, dipahami, dan dimodifikasi oleh programmer lainnya
- clean code dibutuhkan dalam work collaboration, feature development, dan faster development
- terdapat 9 karakteristik dalam clean code:
1. penamaan yang mudah dipahami. gunakan nama2 yang biasanya umum digunakan
2. penamaan yang mudah dieja dan dicari
3. penamaan yang singkat tapi mampu mendeskripsikan konteks
4. penamaan yang konsisten. jika menggunakan metode snake case, maka seterusnya memakai itu. jika menggunakan metode camel case, maka seterusnya memakai itu. dan sebagainya
5. hindari penambahan konteks yang tidak perlu. hal tersebut akan membuat tidak efektif ketika dibaca
6. gunakan komentar secara efektif dan efisien. tidak semua syntax coding perlu dijelaskan menggunakan komentar. tetapi jika perlu, maka komentar ditulis di atas syntax coding
7. good function. karena function dapat memiliki banyak parameter, ada baiknya ketika memanggil function tersebut objek parameternya juga kita tulis
8. gunakan konvensi. karena setiap bahasa programming memiliki gaya penulisannya masing-masing
9. saran formatting. beberapa diantaranya adalah baris code yang berhubungan saling berdekatan, dekatkan variabel dengan penggunanya dan dekatkan fungsi dengan pemanggilnya, perhatikan indentansi, gunakan prettier atau formatter, dkk
- terdapat 2 prinsip dalam clean code:
1. KISS (keep it to simple). lebih mengarah mengenai efektifitas suatu fungsi. fungsi yang dibuat harus jelas digunakan untuk satu tugas saja. jangan terlalu banyak argumen pada fungsi. dan sebagainya
2. DRY (don't repeat yourself). jika mempunyai tugas yang sama dan akan banyak sering digunakan, jadikan tugas tersebut sebagai fungsi sehingga hanya perlu memanggil fungsinya saja dan dapat menghindari duplikasi
- refactoring adalah proses restrukturisasi kode yang dibuat tanpa mengubah outputnya sehingga lebih menjadi clean code. gunakan 2 prinsip KISS dan DRY