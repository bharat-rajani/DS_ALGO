package main

import (
	"fmt"
)

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func mergeSort(a []int64, n int64) []int64 {

	if len(a) < 2 {
		fmt.Println(len(a))
		return a
	}

	// middle := len(a)/2
	// fmt.Println(type(middle))

	var cur, start int64
	for cur = 1; cur < n; cur = 2 * cur {
		for start = 0; start < n; start += 2 * cur {
			// fmt.Printf("start,i %d %d; arr ->", start, min(start+(i*2), n))
			fmt.Printf("%v \n", a[start:min(start+(cur*2), n)])
			mid := min(start+(cur-1), n-1)
			end := min(start+(2*cur)-1, n-1)
			merge(a, start, mid, end)
			fmt.Println(a)

		}
		fmt.Printf("\n")
	}

	return a
}

func merge(a []int64, start, mid, end int64) {
	fmt.Println("#######  start merge", start, mid, end)
	sL := mid - start + 1
	sR := end - mid
	L := make([]int64, sL, sL)
	R := make([]int64, sR, sR)

	for i := int64(0); i < sL; i++ {
		L[i] = a[start+i]
	}

	for i := int64(0); i < sR; i++ {
		R[i] = a[mid+1+i]
	}

	fmt.Println(L, R)
	// fmt.Println("end merge")

	lIdx, rIdx, idx := int64(0), int64(0), start
	for lIdx < sL && rIdx < sR {

		if L[lIdx] <= R[rIdx] {
			a[idx] = L[lIdx]
			lIdx++
		} else if L[lIdx] > R[rIdx] {
			fmt.Println("->>>", mid-lIdx+1)
			a[idx] = R[rIdx]
			rIdx++
		}
		idx++

	}

	for lIdx < sL {
		a[idx] = L[lIdx]
		lIdx++
		idx++
	}

	for rIdx < sR {
		a[idx] = R[rIdx]
		rIdx++
		idx++
	}

}

func main() {
	fmt.Println("hi")
	// a := make([]int64, 5)
	a := []int64{2, 4, 3, 1, 5}
	fmt.Printf("Array-> %v\nSize-> %d \nStarting sorting...\n", a, len(a))
	fmt.Println("Result: ", mergeSort(a, int64(len(a))))
}
