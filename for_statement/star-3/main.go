package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	for i := num; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Printf("\n")
	}
}
