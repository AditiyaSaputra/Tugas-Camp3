package main

import (
	"fmt"
)

func main() {
	// Membuat map dengan tipe key string dan tipe value int
	data := make(map[string]int)

	// Menambahkan data ke dalam map
	data["John"] = 21
	data["Jane"] = 18
	data["Bob"] = 20

	// Mengakses data dalam map
	umurJohn := data["John"]
	umurJane := data["Jane"]
	umurBob := data["Bob"]

	// Cetak data dalam map
	fmt.Printf("Umur John: %d\n", umurJohn)
	fmt.Printf("Umur Jane: %d\n", umurJane)
	fmt.Printf("Umur Bob: %d\n", umurBob)

	// Menghapus data dari map
	delete(data, "Jane")

	// Cetak data setelah menghapus
	fmt.Println("Data setelah menghapus Jane:", data)
}
