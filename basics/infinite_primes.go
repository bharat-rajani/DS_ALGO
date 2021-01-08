package main

import (
	"fmt"
	"math"
)

func main() {

	for i := 2; ; i++ {
		sq := math.Sqrt(float64(i))
		// fmt.Println(int(sq), sq)
		flag := true
		for j := 2; j <= int(sq); j++ {
			// fmt.Println(i, j)
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(i)
		}
	}
}
