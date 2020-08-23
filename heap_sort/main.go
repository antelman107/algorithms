package main

import "fmt"

func siftDown(a []int, elem, max int) {
	current := elem
	for {
		// Left child of current node
		child := 2*current + 1

		// Child is out of slice bound - finish
		if child >= max {
			return
		}

		// Find greatest child
		if child+1 < max && a[child+1] > a[child] {
			child++
		}

		// Child is greater than top element
		if a[current] >= a[child] {
			return
		}

		a[current], a[child] = a[child], a[current]
		current = child
	}

}

func heapSort(a []int) (res []int) {
	// Build heap
	for i := (len(a) - 1) / 2; i >= 0; i-- {
		siftDown(a, i, len(a)-1)
	}

	// Put element from i to end, rebuild remained heap
	for i := len(a) - 1; i >= 0; i-- {
		a[0], a[i] = a[i], a[0]
		siftDown(a, 0, i)
	}

	return a
}

func main() {
	fmt.Println(heapSort([]int{10, 4, 5, 3, 2}))
}
