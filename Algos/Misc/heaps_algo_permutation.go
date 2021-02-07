package main

import (
	"fmt"
)

func generate(k int, arr []int) [][]int {
	// var perms [][]int
	perms := make([][]int, 1)
	fmt.Printf("address of a slice perms %p , perms: %v\n", &perms, perms)
	fmt.Printf("address of a 0th arr element which slice perms points %p \n", perms[0])
	generateRecursive(k, arr, &perms)
	fmt.Println(perms)
	fmt.Println(append(perms, []int{1, 2}))

	fmt.Printf("address of a slice perms %p \n", &perms)
	fmt.Printf("address of a 0th arr element which slice perms points %p \n", perms)

	return perms

}

func generateRecursive(k int, arr []int, perms *[][]int) {

	// fmt.Printf("genRec -> address of a slice perms %p \n", &perms)
	// fmt.Printf("genRec -> address of a 0th arr elem slice perms %p \n", &perms)
	/*
	   	procedure generate(k : integer, A : array of any):
	       if k = 1 then
	           output(A)
	       else
	           // Generate permutations with kth unaltered
	           // Initially k == length(A)
	           generate(k - 1, A)

	           // Generate permutations for kth swapped with each k-1 initial
	           for i := 0; i < k-1; i += 1 do
	               // Swap choice dependent on parity of k (even or odd)
	               if k is even then
	                   swap(A[i], A[k-1]) // zero-indexed, the kth is at k-1
	               else
	                   swap(A[0], A[k-1])
	               end if
	               generate(k - 1, A)

	           end for
	       end if
	*/

	if k == 1 {
		*perms = append(*perms, arr)
		// fmt.Println(perms, "*", &perms)
		fmt.Printf("genRec After append -> address of a slice perms %p ,perms: %v\n", &perms, perms)
		// fmt.Printf("genRec After append -> address of a 0th arr elem slice perms %p \n", &perms)
		// return perms
	} else {
		generateRecursive(k-1, arr, perms)

		for i := 0; i < k-1; i += 1 {

			if k%2 != 0 {
				arr[i], arr[k-1] = arr[k-1], arr[i]
			} else {
				arr[0], arr[k-1] = arr[k-1], arr[0]
			}

			generateRecursive(k-1, arr, perms)
		}
	}
	// return perms
}

func main() {

	arr := []int{1, 2, 3}
	s := generate(len(arr), arr)
	fmt.Println(s)
}
