package main

import (
	"fmt"
)

// Struct: Mahasiswa
type Mahasiswa struct {
	Nama     string
	NIM      string
	Semester int
}

// Method pada struct Mahasiswa untuk mengganti nilai semester menggunakan pointer
func (m *Mahasiswa) UbahSemester(semesterBaru int) {
	m.Semester = semesterBaru
}

func main() {
	// Membuat objek Mahasiswa
	mahasiswa := Mahasiswa{
		Nama:     "Aditiya",
		NIM:      "H1A021011",
		Semester: 4,
	}

	// Cetak data mahasiswa sebelum perubahan
	fmt.Printf("Data mahasiswa sebelum perubahan:\nNama: %s\nNIM: %s\nSemester: %d\n", mahasiswa.Nama, mahasiswa.NIM, mahasiswa.Semester)

	// Mengubah nilai semester menggunakan method UbahSemester
	mahasiswa.UbahSemester(5)

	// Cetak data mahasiswa setelah perubahan
	fmt.Printf("\nData mahasiswa setelah perubahan:\nNama: %s\nNIM: %s\nSemester: %d\n", mahasiswa.Nama, mahasiswa.NIM, mahasiswa.Semester)
}
