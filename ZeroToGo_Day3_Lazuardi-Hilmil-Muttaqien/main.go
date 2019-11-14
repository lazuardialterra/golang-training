package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Lazuardi Hilmil Muttaqien - System Analyst
	// Challenge Day 3

	// var jumlah int
	// fmt.Printf("jumlah push up ")
	// fmt.Scanf("%v", &jumlah)

	// if jumlah > 10 {
	// 	fmt.Println("SUPER KEREN COI")
	// } else if jumlah > 5 && jumlah <= 10 {
	// 	fmt.Println("KEREN BINGIT")
	// } else {
	// 	fmt.Println("CUPU")
	// }

	//task mbuh
	// var Nmaneh int
	// fmt.Printf("Masukkan angka ")
	// fmt.Scanf("%v", &Nmaneh)
	// for i := 1; i <= Nmaneh; i++ {

	// 	if i == 8 {

	// 		fmt.Println("KETEMU YANG DELAPAN!!!")
	// 		continue
	// 	}
	// 	if i >= 9 {
	// 		fmt.Println("HOOOP MANDEK DI SEMBILAN GAK ISOK LANJOT!!!")
	// 		break
	// 	} else {
	// 		fmt.Println(i)
	// 	}

	// }
	// var clue string
	// fmt.Printf("Katakan Peta: ")
	// fmt.Scanf("%v\n", &clue)
	// if clue == "PETA" {
	// 	for i := 0; i < len(clue); i++ {
	// 		fmt.Println(string(clue[i]))
	// 	}
	// } else {
	// 	fmt.Println("KANDANI KATAKAN PETA")
	// }

	// task 1 bilangan faktor
	var N int
	fmt.Println("Bilangan Faktor")
	fmt.Println("Masukkan nilai ")
	fmt.Scanf("%v", &N)

	fmt.Printf("Hasil faktor \n")
	for i := 1; i <= N; i++ {
		if N%i == 0 {
			fmt.Printf("%d\n", i)
		}
	}
	fmt.Println(" ")
	//task 2 prima
	var L int
	fmt.Println("Bilangan Prima")
	var prima bool = true
	fmt.Println("Masukkan nilai ")
	fmt.Scanf("%v\n", &L)
	if L == 1 {
		prima = false
	} else {
		for i := 2; i < L; i++ {
			//fmt.Println(i)
			if L%i == 0 {
				//fmt.Println(L)
				prima = false

			}
		}
	}
	if prima == true {
		println("PRIMA!!")
	} else {
		println("BUKAN PRIMA!!")
	}
	fmt.Println(" ")
	//task 3 palindrom
	fmt.Println("Cek Palindrom")
	var text string
	var pali string
	// fmt.Println("Masukkan text ")
	// fmt.Scanf("%v\n", &text)

	xscannerkey := bufio.NewScanner(os.Stdin)
	xscannerkey.Scan()
	text = xscannerkey.Text()

	for z := len(text); z > 0; z-- {
		//fmt.Println(string(text[z-1]))
		pali += string(text[z-1])
	}
	//fmt.Println(pali)
	if text == pali {
		fmt.Println("PALINDROM")
	} else {
		fmt.Println("BUKAN PALINDROM")
	}

	//task 4 asteriks
	fmt.Println("\nASTERIKS")
	var xasteriks, xasteriks2 int
	fmt.Println("Masukkan nilai ")
	fmt.Scanf("%v\n", &xasteriks)
	//fmt.Println(xasteriks)
	xasteriks2 = xasteriks + 1
	for a := xasteriks; a > 0; a-- {
		for i := 1; i <= a; i++ {
			fmt.Printf(" ")
		}
		for a1 := xasteriks; a1 > a; a1-- {
			fmt.Printf("*")
		}
		for a2 := xasteriks2; a2 > a; a2-- {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
