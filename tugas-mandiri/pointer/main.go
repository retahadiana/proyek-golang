package main

import "fmt"

// Menukar nilai dua integer melalui pointer
func swap(a, b *int) {
	*a, *b = *b, *a
}

// Menambahkan item baru ke slice melalui pointer
func updateSlice(s *[]string, newItem string) {
	*s = append(*s, newItem)
}

// Demo Pass by Value vs Pass by Pointer
func ubahNilaiBiasa(val int) {
	val = 100
}

func ubahNilaiPointer(val *int) {
	*val = 100
}

func main() {
	// 1. Uji swap
	x, y := 10, 20
	fmt.Printf("Sebelum swap: x=%d, y=%d\n", x, y)
	swap(&x, &y)
	fmt.Printf("Setelah swap : x=%d, y=%d\n\n", x, y)

	// 2. Uji updateSlice
	buah := []string{"Apel", "Jeruk"}
	fmt.Println("Sebelum updateSlice:", buah)
	updateSlice(&buah, "Mangga")
	fmt.Println("Setelah updateSlice :", buah)

	// 3. Perbandingan Pass by Value vs Pointer
	n := 50
	ubahNilaiBiasa(n)
	fmt.Println("\nSetelah ubahNilaiBiasa  :", n) // tetap 50
	ubahNilaiPointer(&n)
	fmt.Println("Setelah ubahNilaiPointer:", n) // berubah 100
}
