package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)

	for i := 0; i < len(input); i++ {
		fmt.Printf("%c", input[i])
		if i%10 == 9 {
			fmt.Printf("\n")
		}
	}
	fmt.Println("")
}
