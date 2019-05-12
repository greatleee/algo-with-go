package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	for i := 1; i <= 9; i++ {
		fmt.Printf("%d * %d = %d\n", number, i, i*number)
	}
}
