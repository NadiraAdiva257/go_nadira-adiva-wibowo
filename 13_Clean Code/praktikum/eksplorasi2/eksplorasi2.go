package main

type kendaraan struct {
	// penamaan field tidak mudah dieja karena saling tersambung dengan huruf kecil
	// bisa menggunakan teknik camelcase menjadi totalRoda dan kecepatanPerJam
	totalroda       int
	kecepatanperjam int
}

type mobil struct {
	kendaraan
}

func (m *mobil) berjalan() {
	m.tambahkecepatan(10)
}

// penamaan fungsi dan parameternya tidak mudah dieja karena saling tersambung dengan huruf kecil
// bisa menggunakan teknik camelcase menjadi tambahKecepatan dan kecepatanBaru
func (m *mobil) tambahkecepatan(kecepatanbaru int) {
	// supaya lebih efektif, assingment bisa diganti menjadi m.kecepatanPerJam += kecepatanBaru
	m.kecepatanperjam = m.kecepatanperjam + kecepatanbaru
}

func main() {
	// penamaan objek dari struct tidak mudah dieja karena saling tersambung dengan huruf kecil
	// bisa menggunakan teknik camelcase menjadi mobilCepat
	mobilcepat := mobil{}
	// supaya lebih efektif, pemanggilan fungsi berjalan bisa menggunakan looping sesuai banyak yang diinginkan
	mobilcepat.berjalan()
	mobilcepat.berjalan()
	mobilcepat.berjalan()

	// penamaan objek dari struct tidak mudah dieja karena saling tersambung dengan huruf kecil
	// bisa menggunakan teknik camelcase menjadi mobilLamban
	mobillamban := mobil{}
	mobillamban.berjalan()
}
