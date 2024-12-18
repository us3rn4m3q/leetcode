package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m := 3
	n := 3
	fmt.Println(merge(nums1, m, nums2, n))
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	copy(nums1[m:], nums2[:n])
	Arr := mergesort(nums1)
	copy(nums1, Arr)
	return Arr
}

func mergesort(Arr []int) []int {
	if len(Arr) == 1 {
		return Arr
	}
	ArrOne := mergesort(Arr[:len(Arr)/2])
	ArrTwo := mergesort(Arr[len(Arr)/2:])
	return sort(ArrOne, ArrTwo)
}

func sort(nums1 []int, nums2 []int) []int {
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
