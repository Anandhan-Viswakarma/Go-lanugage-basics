package main

import "fmt"

func main() {

	var b [5]int
	k := 0
	for i := 0; i < 5; i++ {
		b[i] = k
		k++
	}
	for j := 0; j < 5; j++ {
		fmt.Println("\n", b[j])
	}
	var A [5][5]int

	for p := 0; p < 5; p++ {
		for q := 0; q < 5; q++ {
			A[p][q] = 1
			fmt.Printf("%d",A[p][q])

		}
		fmt.Println("\n")

	}

}
