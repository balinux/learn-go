package main

import "fmt"

func main() {
	// if, else if, & else
	var nilai int8 = 8

	if nilai == 10 {
		fmt.Println("anda lulus dengan sempurna")
	} else if nilai > 7 {
		fmt.Print("anda lulus\n")
	} else if nilai < 7 {
		fmt.Println("anda tidak lulus")
	} else {
		fmt.Printf(" anda tidak lulus dengan nilai %d\n", nilai)
	}

	// Variabel Temporary Pada if - else
	point := 8890.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s Good!\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// Seleksi Kondisi Menggunakan Keyword switch - case
	pointx := 6

	switch pointx {
	case 8:
		fmt.Println("good")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("biasa aja")
	}

	// Pemanfaatan case Untuk Banyak Kondisi
	pointy := 6

	switch pointy {
	case 8, 9, 10:
		fmt.Println("good")
	case 5, 6, 7:
		fmt.Println("awesome")
	default:
		fmt.Println("biasa aja")
	}

	// Kurung Kurawal Pada Keyword
	pointa := 1

	switch pointa {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	// Switch Dengan Gaya if - else
	pointb := 9

	switch {
	case pointb == 8:
		fmt.Println("perfect")
	case (pointb < 8) && (pointb > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("selamat")
		}
	}

	// Penggunaan Keyword fallthrough Dalam switch
	pointc := 7

	switch {
	case pointc == 8:
		fmt.Println("perfect")
	case (pointc < 8) && (pointc > 3):
		fmt.Println("awesome")
		fallthrough
	default:
		{
			fmt.Println("selamat fallthrough")
		}
	}
}
