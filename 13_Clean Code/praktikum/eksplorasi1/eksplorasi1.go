package main

type user struct {
	id int
	// tipe data seharusnya string karena username termasuk nama
	username int
	password int
}

// penamaan struct tidak mudah dieja karena saling tersambung dengan huruf kecil
// bisa menggunakan teknik camelcase menjadi userService
type userservice struct {
	// penamaan field tidak mudah dipahami. apa itu t?
	t []user
}

// penamaan fungsi tidak mudah dieja karena saling tersambung dengan huruf kecil
// bisa menggunakan teknik camelcase menjadi getAllUsers
// penamaan u untuk struct userservice sudah singkat namun tidak mendeskripsikan konteks
// seharusnya us diambil dari (u)ser(s)ervice
// kalau penamannya u konteksnya bisa ke struct user
func (u userservice) getallusers() []user {
	// return field tidak mudahi dipahami. apa itu t?
	return u.t
}

// penamaan fungsi tidak mudah dieja karena saling tersambung dengan huruf kecil
// bisa menggunakan teknik camelcase menjadi getUserById
func (u userservice) getuserbyid(id int) user {
	// penamaan r untuk value dari slice tersebut sudah singkat namun tidak mendeskripsikan konteks
	// seharusnya value supaya lebih jelas
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}

func main() {

}
