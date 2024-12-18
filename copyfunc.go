package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{8, 7, 8, 9, 10}
	m := 3
	n := 3
	fmt.Println(copyfunc(arr1, m, arr2, n))
}

func copyfunc(arr1 []int, m int, arr2 []int, n int) []int {
	arr := append(arr1[:m], arr2[:n]...)
	return arr
}
