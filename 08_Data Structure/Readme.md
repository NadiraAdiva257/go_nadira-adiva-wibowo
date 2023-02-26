Resume Materi "Data Structure"
- terdapat 3 data structure yang akan dibahas dalam resume ini:
1. array
2. slice
3. map
- tentang array
1. array adalah kumpulan data dengan satu tipe data yang sama
2. panjang array atau daya simpan array harus ditentukan dari awal
3. syntax pembuatan array secara tidak langsung adalah var namaArray [panjangArray]tipeDataArray
4. syntax pembuatan array secara langsung adalah var namaArray = []tipeDataArray{data}
4. syntax pemberian atau modifikasi data dari array adalah namaArray[index] = data
5. syntax memanggil data dari array adalah namaArray[index]
6. dengan menggunakan looping, memanggil data dari array lebih mudah karena tidak perlu menuliskan index nya satu-satu
7. terdapat 3 syntax untuk looping tersebut:
(1) for i, element := range namaArray {
		fmt.Print("index", i, "=", element)
	}
(2) for i := 0; i < len(namaArray); i++ {
		fmt.Println(namaArray[i])
	}
(3) i := 0
	for range namaArray {
		fmt.Println(namaArray[i])
		i++
	}
8. syntax mengetahui panjang array adalah len(namaArray)
9. terdapat fungsi append untuk menambahkan data ke dalam array pada index paling akhir. syntaxnya adalah append(namaArray, data)
- tentang slice
1. slice adalah potongan data yang diambil dari array atau bisa disebut reference type dari array
2. panjang slice atau daya simpan slice dapat diubah
3. terdapat 3 tipe data dalam slice:
(1) pointer adalah penunjuk data. kalau data yang di array berubah, maka data yang di slice juga berubah
(2) len adalah panjang slice dari data yang sudah tersedia
(3) cap adalah maksimal kapasitas panjang slice
4. syntax pembuatan slice secara tidak langsung adalah var namaSlice []tipeDataSlice
5. syntax pembuatan slice secara langsung adalah var namaSlice = []tipeDataSlice{data}
5. syntax pembuatan slice dengan keyword make adalah make([]tipeDataSlice, len, cap) 
6. syntax pemberian atau modifikasi data dari slice adalah namaSlice[index] = data
7. syntax memanggil data dari slice adalah namaSlice[index]
8. dengan menggunakan looping, memanggil data dari slice lebih mudah karena tidak perlu menuliskan index nya satu-satu. syntax untuk looping sama seperti di array
9. terdapat fungsi copy untuk mengambil slice yang sama persis dari slice lain. syntaxnya adalah copy(destinationSlice, sourceSlice). pastikan bahwa kapasitas keduanya sama
10. terdapat fungsi append untuk menambahkan data ke dalam slice pada index paling akhir. syntaxnya adalah append(namaSlice, data)
11. kapasitas dalam slice akan bertambah 2x lipat jika panjang slice bertambah 1 dari maksimalnya kapasitas
- tentang map
1. map adalah kumpulan key dan value. dimana key harus unik
2. tipe data untuk key tidak terbatas hanya integer. bisa juga string
3. syntax pembuatan map secara tidak langsung adalah var namaMap map[tipeDataKey]tipeDataValue
4. syntax pembuatan map secara langsung adalah var namaMap = map[tipeDataKey]tipeDataValue{key:value}
4. syntax pembuatan map dengan keyword make adalah make(map[tipeDataKey]tipeDataValue)
5. syntax pemberian atau modifikasi data dari map adalah namaMap[key] = value
6. syntax memanggil data dari map adalah namaMap[key]
7. dengan menggunakan looping, memanggil data dari map lebih mudah karena tidak perlu menuliskan key nya satu-satu. syntax untuk looping sama seperti di array