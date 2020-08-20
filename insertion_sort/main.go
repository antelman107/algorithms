package main

import "fmt"

func insertionSort(a []int) (res []int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j > 0 && j < len(a) && a[j] < a[j-1]; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}

	return a
}

func main() {
	fmt.Println(insertionSort([]int{10, 4, 2, 3, 5}))
}
