package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x)
	fmt.Scan(&y)

	total := 0
	for i := 1; i < x; i++ {
		switch {
		case i == 1:
			fallthrough
		case i == 3:
			fallthrough
		case i == 5:
			fallthrough
		case i == 7:
			fallthrough
		case i == 8:
			fallthrough
		case i == 10:
			fallthrough
		case i == 12:
			total += 31
		case i == 4:
			fallthrough
		case i == 6:
			fallthrough
		case i == 9:
			fallthrough
		case i == 11:
			total += 30
		case i == 2:
			total += 28
		}
	}
	total += y

	nameoji := total % 7

	var day string
	switch {
	case nameoji == 0:
		day = "SUN"
	case nameoji == 1:
		day = "MON"
	case nameoji == 2:
		day = "TUE"
	case nameoji == 3:
		day = "WED"
	case nameoji == 4:
		day = "THU"
	case nameoji == 5:
		day = "FRI"
	case nameoji == 6:
		day = "SAT"
	}
	fmt.Println(day)
}
