package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(number, "is Even")
		} else {
			fmt.Println(number, "is Odd")
		}
	}
}
