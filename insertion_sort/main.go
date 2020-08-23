package main

import "fmt"

func insertionSort(a []int) (res []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && a[j] < a[j-1]; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}

	return a
}

func main() {
	fmt.Println(insertionSort([]int{10, 4, 2, 3, 5}))
}
