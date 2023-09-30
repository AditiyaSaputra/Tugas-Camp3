package main

import (
	"fmt"
)

func main() {
	// Membuat map dengan tipe key string dan tipe value int
	data := make(map[string]int)

	// Menambahkan data ke dalam map
	data["Yeremia"] = 21
	data["Alice"] = 18
	data["Aditiya"] = 20

	// Mengakses data dalam map
	umurJohn := data["Yeremia"]
	umurJane := data["Alice"]
	umurBob := data["Aditiya"]

	// Cetak data dalam map
	fmt.Printf("Umur Yeremia: %d\n", umurJohn)
	fmt.Printf("Umur Alice: %d\n", umurJane)
	fmt.Printf("Umur Aditiya: %d\n", umurBob)

	// Menghapus data dari map
	delete(data, "Alice")

	// Cetak data setelah menghapus
	fmt.Println("Data setelah menghapus Alice:", data)
}
