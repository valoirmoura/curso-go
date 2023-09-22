package main

import "fmt"

func main() {

	s := []int{2, 4, 6, 8, 10}
	fmt.Printf("len=%d cap=%d %v\n\n", len(s), cap(s), s)

}
