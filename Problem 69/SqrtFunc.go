package main

import (
	"fmt"
	"math"
)

func main() {
	x := 0
	fmt.Scan(&x)
	x = mysqrt(x)
	fmt.Println(x)
}
func mysqrt(x int) int {
	X := float64(x)
	X = math.Sqrt(float64(X))
	x = int(X)
	return x
}
