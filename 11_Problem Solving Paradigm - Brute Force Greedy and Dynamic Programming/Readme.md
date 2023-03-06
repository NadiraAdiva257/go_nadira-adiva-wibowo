Resume Materi "Problem Solving Paradigm"
- problem solving paradigm adalah cara untuk menyelesaikan sebuah amsalah dengan berbagai prinsip
- terdapat 4 prinsip yang paling sering digunakan yaitu brute force (complete search), divide and conquer, greedy approach, dan dynamic programming
- tentang brute force
1. cara kerja brute force dengan mengecek semua data satu per satu sehingga time complexity nya adalah O(n)
2. kurang efektif digunakan sehingga menjadi opsi terakhir jika prinsip yang lain tidak bisa digunakan
3. biasanya untuk pencarian yang searah dan untuk program yang unlimited time
- tentang divide and conquer
1. cara kerjanya dengan memecah kelompok data menjadi bagian kecil-kecil
2. biasanya dibagi menjadi 2 atau biner
3. contoh yang sering dipakai adalah binary search dengan syaratnya data harus urut
4. cara kerja binary search pada umumnya adalah jika yang dicari berada di barisan sebelah kanan setelah urutan datanya dibagi 2, maka hanya fokus pada sebelah kanan. begitu seterusnya sampai yang dicari ketemu
- tentang greedy approach
1. pendekatan dengan mencari nilai local optimal, bukan global optimal
2. biasanya untuk data yang berbentuk graph, map, dkk
3. cara kerjanya adalah akan mencari nilai optimal dari local walaupun kedepannya belum tentu jalur yang dipilih adalah yang optimal
4. contoh selain graph dan map adalah huffman coding, activity selection, dijkstra algorith, dkk
- tentang dynamic programming
1. cara kerja umumnya adalah akan menyimpan seluruh kemungkinan kemudian dicari yang paling optimal
2. karakteristiknya adalah overlapping subproblems dan optimal substructure property
3. terdapat 2 method yaitu topdown dan bottom up
4. topdown berarti menyelesaikan masalah dari yang paling besar ke yang lebih kecil. biasanya dengan rekursif
5. bottom up berarti menyelesaikan masalah dari yang paling kecil ke yang lebih besar. biasanya memakai n-dimensional table