package main

import "fmt"

func main() {
	// Deklarasi 5 variabel dengan tipe berbeda
	var nama string = "Ahmad"
	var umur int = 21
	var ipk float64 = 3.85
	var isLulus bool = true
	var hobi []string = []string{"Coding", "Membaca"}

	fmt.Printf("Nama: %s, Umur: %d, IPK: %.2f, Status: %t, Hobi: %v\n\n", nama, umur, ipk, isLulus, hobi)

	// Map data mahasiswa (Nama -> Nilai)
	mahasiswa := make(map[string]float64)

	// 1. Menambah data
	mahasiswa["Budi"] = 85.5
	mahasiswa["Siti"] = 92.0
	mahasiswa["Andi"] = 78.0

	// 2. Membaca dengan pengecekan keberadaan
	if nilai, ada := mahasiswa["Siti"]; ada {
		fmt.Printf("Nilai Siti: %.2f\n", nilai)
	}

	// 3. Menghapus data
	delete(mahasiswa, "Andi")

	// 4. Menelusuri seluruh isi map
	fmt.Println("\nDaftar Mahasiswa Tersisa:")
	for name, score := range mahasiswa {
		fmt.Printf("- %s : %.2f\n", name, score)
	}
}
