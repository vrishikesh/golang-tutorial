package main

import "fmt"

func main() {
	slice := make([]int, 11)

	for i := range slice {
		if i%2 == 0 {
			fmt.Println(i, " is even")
		} else {
			fmt.Println(i, " is odd")
		}
	}

	fmt.Println(slice)
}
