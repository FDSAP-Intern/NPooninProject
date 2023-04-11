package main

import "fmt"

func main() {
	arr := [5]int{4, 1, 3, 1, 1}

	for x := 0; x < len(arr); x++ {
		for y := 0; y < arr[x]; y++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

