package main

import "fmt"

func main() {
	nums1 := []int{0, 7, 2, 3, 4, 8, 6, 3, 8, 9}
	fmt.Println(split(nums1))
}

func split(Arr []int) []int {
	if len(Arr) == 1 {
		return Arr
	}
	ArrOne := split(Arr[:len(Arr)/2])
	ArrTwo := split(Arr[len(Arr)/2:])
	return mergesort(ArrOne, ArrTwo)
}

func mergesort(nums1 []int, nums2 []int) []int {
	Arr := []int{}
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			Arr = append(Arr, nums1[i])
			i++
		} else {
			Arr = append(Arr, nums2[j])
			j++
		}
	}
	for ; i < len(nums1); i++ {
		Arr = append(Arr, nums1[i])
	}
	for ; j < len(nums2); j++ {
		Arr = append(Arr, nums2[j])
	}
	return Arr
}
