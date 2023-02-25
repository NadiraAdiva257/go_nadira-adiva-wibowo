Resume Materi "Time Complexity and Space Complexity"
- time complexity adalah lamanya waktu yang digunakan untuk menjalankan sebuah algoritma
- space complexity adalah besaran memori yang digunakan untuk menjalankan sebuah algoritma
- dalam time complexity, biasanya ditentukan mana yang termasuk operasi dominan dan seberapa banyak operasi dominan tersebut dilakukan
- BigO notation adalah metode untuk mengetahui time complexity dan space complexity
- terdapat 7 jenis BigO notation dalam time complexity:
1. O(1) / constant time adalah operasi yang hanya dilakukan sekali walaupun jumlah inputannya banyak. termasuk operasi yang paling cepat
2. O(n) / linear complexity adalah banyaknya operasi yang dilakukan sesuai dengan jumlah inputannya. operasi ini biasanya paling banyak digunakan
3. O(n+m) / linear time adalah 2 operasi dengan bentuk dominan yang sama
4. O(log n) / logarithmic time adalah banyaknya operasi yang dimisalkan dalam 2^x = n. dimana n adalah jumlah inputan dan x adalah operasi yang dilakukan. ciri-cirinya adalah n/2. sering digunakan untuk binary search. termasuk operasi yang efektif
5. O(n^2) / quadratic time adalah operasi bersarang yang dilakukan sesuai dengan jumlah inputannya dikuadratkan. ciri-cirinya adalah nested if dan nested loop
6. O(n!) / factorial time adalah banyaknya operasi didapatkan dari seperti factorial dalam matematika
7. O(2^n) / exponential time adalah banyaknya operasi didapatkan dari seperti exponential dalam matematika
- urutan paling cepat dari jenis-jenis BigO notation dalam time complexity:
1. O(1) / constant time
2. O(log n) / logarithmic time
3. O(n) / linear complexity
4. O(n^2) / quadratic time
5. dst
- space complexity biasanya dapat dilihat dari banyaknya variabel pada suatu operasi algoritma
- beberapa jenis BigO notation dalam space complexity:
1. O(1) / constant time adalah banyaknya variabel tetap sama walaupun jumlah inputannya banyak
2. O(n) / linear complexity adalah banyaknya value yang didapatkan dari append yang disesuaikan dengan jumlah inputannya
- kunci utamanya adalah bagaimana menghasilkan output dari jalannya algoritma dengan cepat dan hemat memori