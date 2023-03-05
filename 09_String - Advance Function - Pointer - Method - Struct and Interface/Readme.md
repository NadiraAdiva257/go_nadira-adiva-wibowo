Resume Materi "String - Advance Function - Pointer - Method - Struct and Interface"
- tentang string
1. terdapat library strings yang isinya mencakup berbagai function tentang string
2. panjang string dapat diketahui dengan len()
3. antar string dapat dibandingkan dengan menggunakan operator ==
4. apakah terdapat string di suatu string lain dapat diketahui dengan strings.Contains()
5. mengambil beberapa string dari suatu string lain dengan index yang ditentukan [index:index]
6. mengganti huruf atau kata tertentu dari string dengan strings.Replace()
7. memasukkan huruf atau kata tertentu ke suatu string dengan index yang ditentukan [index:index]
8. mengganti semua string menjadi huruf kecil dengan strings.ToLower()
9. mengganti semua string menjadi huruf besar dengan strings.ToUpper()
10. memotong cutset di awal dan akhir string dengan strings.Trim()
11. memotong string berdasarkan separator yang hasilnya adalah array dengan strings.Split()
- tentang advance function
1. variadic function adalah function yang parameternya bisa menerima banyak inputan tanpa harus mendeklarasikan sebagai slice. syntax parameternya adalah (var ...tipeData)
2. anonymous function adalah function tanpa nama yang biasanya untuk mengelompokkan kumpulan code dalam function sebaris.
3. terdapat beberapa syntax untuk anonymous function:
(1) func(){
    ...
}()
(2) var := func(){
    ...
}
var()
(3) func(var tipeData){
    ...
}(param)
(4) var := func(var tipeData){
    ...
}
var(param)
4. closure adalah anonymous function yang mereference variabel yang dideklarasi di luar function tersebut sehingga reference tersebut akan ikut berubah ketika ada pergantian nilai
5. defer function adalah function yang akan selalu dipanggil terakhir setelah semua code dijalankan. prosesnya seperti stuck di LIFO (last in first out)
- tentang pointer
1. default variable di golang adalah passing by value sehingga untuk mengubah menjadi passing by reference dibutuhkan pointer
2. pointer adalah variable yang disimpan di dalam memori yang sama sehingga jika ada perubahan akan mempengaruhi semuanya
3. operator * adalah dereferencing untuk mendeklarasi variable pointer dan mengakses nilai dari address di memori
4. operator & adalah referencing untuk mengakses dan mengembalikan address dari variable pointer
5. keyword new dalam deklarasi pointer hanya akan mengembalikan pointer ke data kosong
6. zero value pointer nil juga hanya akan mengembalikan pointer ke data kosong
- tentang struct
1. struct adalah type yang menyimpan berbagai field dengan tipe data yang berbeda. mirip sebagai objek
2. syntax untuk membuat struct adalah type namaStruct struct {var tipeData}
3. terdapat 3 cara untuk deklarasi dan akses struct:
(1) var namaObjek namaStruct
namaObjek.field = ...
(2) namaObjek := namaStruct{
    field: ...
}
(3) namaObjek := namaStruct{...}
(4) namaObjek := new(namaStruct)
namaObjek.field = ...
4. terdapat struct method yang receiver dari struct. contohnya adalah:
type namaStruct struct {var tipeData}
func (ns namaStruct) SayHello(){}
5. untuk mengakses struct method dengan namaObjek.namaMethod
- tentang method
1. method adalah bagian dari function yang menempel pada suatu type seperti struct dan interface
2. method membantu penulisan object oriented style, menghindari konflik penamaan, dan mudah dibaca dipahami
3. syntax dari method adalah func (receiver typeStruct) namaMethod(parameterMethod) (return tipeData){}
- tentang interface
1. interface adalah suatu kontrak yang berisi method2 yang mendefinisikan behavior umum dari struct. misalnya method di dalam interface adalah berlari() sehingga method tersebut dapat diaplikasikan dengan type struct Manusia dan Hewan karena keduanya sama2 memiliki perilaku berlari
2. syntax untuk membuat interface adalah type namaInterface interface {namaMethod tipeData}
3. method di dalam interface tersebut nanti dibuat dengan kontrak yang sama persis
4. interface kosong adalah interface yang tidak punya method apapun di dalamnya sehingga cocok untuk dynamic value
- tentang package
1. package adalah koleksi dari kumpulan data dan function
2. package dapat diakses oleh package lain jika penulisan variablenya diawali dengan huruf besar
- tentang error handling
1. panic adalah kondisi jika terjadinya error. panic akan menghentikan semua code kecuali defer sehingga perlu ada pencegahan
2. recover adalah pencegahan terhadap panic sehingga program tetap jalan dengan pesan error yang diberikan dan pastikan recover disimpan di dalam defer