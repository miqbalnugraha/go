package main

import "fmt"

type BangunDatar interface {
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

func main() {
	persegiPanjang := PersegiPanjang{Panjang: 6, Lebar: 5}
	luasA := SeberapaLuas(persegiPanjang)
	fmt.Println("Luas Persegi Panjang:", luasA)

	persegi := Persegi{Sisi: 4}
	luasB := SeberapaLuas(persegi)
	fmt.Println("Luas Persegi:", luasB)
}

func SeberapaLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}
