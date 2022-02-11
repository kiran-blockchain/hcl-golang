package main

import "fmt"

func main() {
	var arr = []int{13, 23, 3, 5, 6, 819}
	fmt.Println("Before Sorting arr", arr)
	final := qSort(arr, 0,2)
	fmt.Println("after Sorting arr", final)
}

/*
pass the array
pass low value
pass high value
*/

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
func qSort(arr []int, low, high int) []int {
	// if low is less than high then swap
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = qSort(arr, low, p-1)
		arr = qSort(arr, p+1, high)
	}
	return arr

}
