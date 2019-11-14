package main

import "fmt"

func main() {
	//Day 2
	//by Lazuardi Hilmil Muttaqien - System Analyst

	// opo kowe lali
	//var a = "anak presiden"
	//fmt.Println("Hello, saya adalah anak gembala")
	//fmt.Println("Atau saya adalah " + a)

	// var name, gender string = "bor ", "L"
	// fmt.Println("nama saya " + name + "jenis kelamin saya " + gender)
	// fmt.Printf("nama saya %s jenis kelamin saya %s \n", name, gender)

	//boolean
	// isLogin := false

	// //int
	// var a int = 15
	// b := 20

	// //float
	// c, d := 2.1, 30

	// //string
	// s := "hello!"

	// //const
	// const pi = 3.14

	// fmt.Println(isLogin, a, b, c, d, s, pi)

	//task 1
	var A, T, Luas float32
	fmt.Println("*************")
	fmt.Println("MENGHITUNG LUAS SEGITIGA")
	fmt.Print("masukkan panjang alas: ")
	fmt.Scanf("%v", &A)
	fmt.Print("masukkan tinggi: ")
	fmt.Scanf("%v", &T)
	// fmt.Println(inputan)
	// A = 20.0
	// T = 25.0
	Luas = (A * T) / 2

	fmt.Printf("Luas segitiga biru adalah %.2f \n", Luas)
	fmt.Println("*************")
	// a := 10
	// b := 15
	// luas2 := (a * b) / 2
	// fmt.Println(luas2)

	//task 2
	fmt.Println("*************")
	fmt.Println("MENGHITUNG LUAS TABUNG")
	var luastabung, r, tinggi float32
	const pi float32 = 3.14

	fmt.Print("masukkan jari-jari: ")
	fmt.Scanf("%v", &r)
	fmt.Print("masukkan tinggi ")
	fmt.Scanf("%v", &tinggi)
	luastabung = 2 * pi * r * (r + tinggi)
	fmt.Printf("Luas tabung adalah %.2f \n", luastabung)
	fmt.Println("*************")
	// r = 4.0
	// tinggi = 20.0
	// const pi = 3.14

	// voltabung = (2.0 * r * r * pi) * tinggi
	// fmt.Println(voltabung)

	// var inputan int
	// fmt.Print("masukkan data: ")
	// fmt.Scanf("%v", &inputan)
	// fmt.Println(inputan)
}
