package main

import (
	"chx-passport/api"
	"chx-passport/setup"
)

func main() {
	setup.Init()
	api.RunApi()
}

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	pivot := arr[0]
	left := make([]int, 0)
	right := make([]int, 0)
	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
}
