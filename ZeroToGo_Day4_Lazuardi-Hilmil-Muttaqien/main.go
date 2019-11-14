package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//task 1

	var xinputtext1, xinputtext2 string
	var text1 []string
	var text2 []string
	fmt.Println("ARRAY MERGE")
	fmt.Println("Masukkan Text 1 : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	xinputtext1 = scanner.Text()
	text1 = strings.Split(strings.Trim(xinputtext1, " "), ",")

	fmt.Println("Masukkan Text 2 : ")
	xscanner := bufio.NewScanner(os.Stdin)
	xscanner.Scan()
	xinputtext2 = xscanner.Text()
	text2 = strings.Split(strings.Trim(xinputtext2, " "), ",")

	//text1 := []string{"jan", "mar", "may", "jul", "sep"}
	//fmt.Println(odd)

	//text2 := []string{"jan", "feb", "mar", "apr", "des"}
	xcount := len(text2)
	for i := 0; i < xcount; i++ {
		text1 = (append(text1, text2[i]))
	}

	maptext := make(map[string]string)
	newtext := make([]string, 0)
	for _, val := range text1 {
		maptext[val] = "key"
		//fmt.Println(val)
	}
	//fmt.Println(maptext)
	for xletter := range maptext {
		newtext = append(newtext, xletter)
		//fmt.Printf("%v ", xletter)
	}
	//fmt.Println(check)
	fmt.Println(newtext)

	//task 2
	fmt.Println("\nMEAN and MEDIAN")
	var xmean, xmedian float64

	number := [...]int{3, 5, 8, 5, 3}
	countnumber := len(number)
	xbagi := countnumber / 2
	var xtotal int
	for _, val := range number {
		xtotal += val
	}
	xmean = float64(xtotal) / float64(countnumber)
	fmt.Printf("mean e %v\n", xmean)

	if countnumber%2 == 0 {
		xmedian = (float64(number[xbagi]) + float64(number[xbagi-1])) / 2
	} else {
		xmedian = (float64(number[xbagi]))
	}
	fmt.Printf("median e %v\n", xmedian)

	//task 3
	fmt.Println("\nCSV to MAP")
	var xkey, xvalue string
	var xkeysplit []string
	var xvaluesplit []string
	//var mapbaru = map[string]string{}

	fmt.Println("masukkan keys : ")
	xscannerkey := bufio.NewScanner(os.Stdin)
	xscannerkey.Scan()
	xkey = xscannerkey.Text()
	fmt.Println("masukkan value : ")
	xscannervalue := bufio.NewScanner(os.Stdin)
	xscannervalue.Scan()
	xvalue = xscannervalue.Text()

	xkeysplit = strings.Split(xkey, ",")
	xvaluesplit = strings.Split(xvalue, ",")

	//textx := []string{"firstname", "lastname", "age"}

	//textz := []string{"akunya", "kamu", "17"}

	csvmap := make(map[string]string)
	// for i := 0; i < len(textx); i++ {
	// 	csvmap[textx[i]] = textz[i]
	// }
	for i := 0; i < len(xkeysplit); i++ {
		csvmap[xkeysplit[i]] = xvaluesplit[i]
	}
	fmt.Println(csvmap)

	//task 4
	// create slice untuk menampung jawaban
	// split input kedalam slice per command
	// split yang didalam slice menjadi 3

	fmt.Println("\nSlice Command")
	var xcmdinput string
	var xarray1 []string
	//var xarray2 []string

	fmt.Println("Masukkan Command : ")
	xcmdscanner := bufio.NewScanner(os.Stdin)
	xcmdscanner.Scan()
	xcmdinput = xcmdscanner.Text()
	xarray1 = strings.Split(xcmdinput, ",")
	//fmt.Println(len(xarray1))
	//fmt.Println(xarray1)
	//var xcountarray1 int = len(xarray1)
	xoutputslice := make([]int, 3, 3)

	for _, xz := range xarray1 {
		xarray2 := strings.Fields(xz)
		switch xarray2[0] {
		case "insert":
			fmt.Println("insert")
			xposisi, _ := strconv.Atoi(xarray2[1])
			xvalue, _ := strconv.Atoi(xarray2[2])
			xoutputslice[xposisi-1] = xvalue
			fmt.Println(xoutputslice)

		case "remove":
			fmt.Println("remove")
			xposisi, _ := strconv.Atoi(xarray2[1])
			xoutputslice[xposisi-1] = xoutputslice[len(xoutputslice)-1]
			//fmt.Println(xoutputslice[xposisi-1])
			//fmt.Println(xoutputslice[len(xoutputslice)-1])
			xoutputslice[len(xoutputslice)-1] = 0
			//fmt.Println(xoutputslice[len(xoutputslice)-1])
			xoutputslice = xoutputslice[:len(xoutputslice)-1]

			fmt.Println(xoutputslice)
		case "append":
			fmt.Println("append")
			xvalue, _ := strconv.Atoi(xarray2[1])
			//fmt.Println(xvalue)
			xoutputslice = append(xoutputslice, xvalue)
			fmt.Println(xoutputslice)

		case "sort":
			fmt.Println("sort")
			sort.Ints(xoutputslice)
			fmt.Println(xoutputslice)

		case "reverse":
			fmt.Println("reverse")
			sort.Ints(xoutputslice)
			sort.Slice(xoutputslice, func(i, j int) bool {
				return xoutputslice[i] > xoutputslice[j]
			})
			fmt.Println(xoutputslice)

		case "pop":
			xoutputslice[len(xoutputslice)-1] = 0
			xoutputslice = xoutputslice[:len(xoutputslice)-1]
			fmt.Println(xoutputslice)
		}

		// xsplit2 = append(xsplit2, xzsplit)
		//spew.Dump(xzsplit)
	}

	fmt.Println("Final Output ", xoutputslice)
	//spew.Dump(xsplit2)

}
