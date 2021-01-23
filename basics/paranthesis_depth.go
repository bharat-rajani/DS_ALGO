package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println("hi")
	dl := []int{12, 65, 24, 243, 14, 62, 68, 3, 4, 5, 113}
	sort.Ints(dl)
	fmt.Println(dl)
	fmt.Println(dl)
	for i := 0; i < 3; i++ {
		fmt.Printf("%p\n", &dl)
		e, dl := maxLessThanX(dl, 113)
		fmt.Println(dl, e)
	}

}

func maxLessThanX(arr []int, x int) (int, []int) {

	fmt.Println("Recvd->", arr)
	lmx := int(math.Inf(-1))
	limx := -1
	gmx := lmx
	gimx := limx

	for idx, i := range arr {
		if i < x {
			lmx = i
			limx = idx
		}
		if lmx >= gmx {
			gmx = lmx
			gimx = limx
		}

	}
	// fmt.Println("->", lmx, gmx, x, limx, gimx)
	fmt.Println("->", gimx)
	return gmx, append(arr[:gimx], arr[gimx+1:]...)
}
