package main

import "fmt"

func main() {
	x := 0
	fmt.Scan(&x)
	x = mySqrt(x)
	fmt.Println(x)
}

func mySqrt(x int) int {
	max, min, mid := x, 0, 0
	for max > min {
		mid = (max + min) / 2
		if mid*mid < x {
			min = mid - 1
		} else if mid*mid > x {
			max = mid + 1
		} else {
			x = mid
		}
	}
	return x
}
