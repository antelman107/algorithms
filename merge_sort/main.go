package main

import "fmt"

func mergeSort(a []int) []int {
	if len(a) == 1 {
		return a
	}

	return merge(mergeSort(a[:len(a)/2]), mergeSort(a[len(a)/2:]))
}

func merge(a, b []int) (res []int) {
	res = make([]int, len(a)+len(b))

	i := 0
	iA := 0
	iB := 0
	for iA < len(a) && iB < len(b) {
		if a[iA] < b[iB] {
			res[i] = a[iA]
			iA++
		} else {
			res[i] = b[iB]
			iB++
		}
		i++
	}

	for j := iA; j < len(a); j++ {
		res[i] = a[j]
		i++
	}

	for j := iB; j < len(b); j++ {
		res[i] = b[j]
		i++
	}

	return
}

func main() {
	fmt.Println(mergeSort([]int{1, 4, 6, 7, 8, 9, 2, 3, 5, 10}))
}
