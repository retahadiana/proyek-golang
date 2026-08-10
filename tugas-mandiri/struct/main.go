package main

import "fmt"

type Student struct {
	ID       int
	Name     string
	Grade    float64
	IsActive bool
}

// Value receiver: Hanya membaca data (tidak mengubah struct)
func (s Student) GetInfo() string {
	return fmt.Sprintf("ID: %d | Name: %s | Grade: %.2f | Active: %v", s.ID, s.Name, s.Grade, s.IsActive)
}

// Pointer receiver: Memperbarui nilai Grade pada struct asli
func (s *Student) UpdateGrade(grade float64) {
	s.Grade = grade
}

// Pointer receiver: Mengubah status aktif pada struct asli
func (s *Student) Activate() {
	s.IsActive = true
}

// Pointer receiver: Mengubah status non-aktif pada struct asli
func (s *Student) Deactivate() {
	s.IsActive = false
}

func main() {
	mhs := Student{ID: 101, Name: "Rian", Grade: 3.20, IsActive: false}

	fmt.Println("Awal       :", mhs.GetInfo())

	mhs.Activate()
	mhs.UpdateGrade(3.75)

	fmt.Println("Setelah Update:", mhs.GetInfo())
}
