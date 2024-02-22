package main

import "fmt"

func sliceChange(a []int) {
	a[0] = 0
	a = append(a, 4) // nothing happens
}

func sliceChangeByPointer(b *[]int) {
	*b = append(*b, 4)
}

func main() {
	a := make([]int, 0)
	a = append(a, 1, 2, 3)
	fmt.Println(a) // 1, 2, 3
	sliceChange(a)
	fmt.Println(a) // 0, 2, 3
	sliceChangeByPointer(&a)
	fmt.Println(a) // 0, 2, 3, 4
}
