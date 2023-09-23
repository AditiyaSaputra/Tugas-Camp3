package main

import (
	"fmt"
)

// Struct pertama: Manusia
type Manusia struct {
	Nama         string
	Usia         int
	JenisKelamin string
}

// Method pada struct Manusia
func (m Manusia) Sapa() {
	fmt.Printf("Halo, nama saya %s. Saya berusia %d tahun.\n", m.Nama, m.Usia)
}

// Struct kedua: Kendaraan
type Kendaraan struct {
	Merek string
	Model string
	Tahun int
}

// Method pada struct Kendaraan
func (k Kendaraan) Informasi() {
	fmt.Printf("Ini adalah %s %s tahun %d.\n", k.Merek, k.Model, k.Tahun)
}

func main() {
	// Membuat objek Manusia
	manusia := Manusia{
		Nama:         "Aditiya",
		Usia:         20,
		JenisKelamin: "Laki-laki",
	}

	// Memanggil method Sapa pada objek Manusia
	manusia.Sapa()

	// Membuat objek Kendaraan
	kendaraan := Kendaraan{
		Merek: "Honda",
		Model: "Mega Pro",
		Tahun: 2012,
	}

	// Memanggil method Informasi pada objek Kendaraan
	kendaraan.Informasi()
}
