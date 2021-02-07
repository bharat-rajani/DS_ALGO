


/*
Problem: https://leetcode.com/problems/ones-and-zeroes/
Status : Accepted

This is the best I could do (BICD) solution.
BICD problems are solely written by Bharat Rajani without going through discussions,solutions or googling,
therefore these can be dumb :)

Concepts learnt:
    GO: https://dave.cheney.net/2018/05/29/how-the-go-runtime-implements-maps-efficiently-without-generics

Notes: Bharat Rajani is not a competitive programmer.
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestSubset(cntArr [][2]int, m, n, l int, mem map[[3]int]int) int {
	if l == 0 || (m == 0 && n == 0) {
		return 0
	}

	if elem, ok := mem[[3]int{l, m, n}]; ok {
		return elem
	}

	if cntArr[l-1][0] <= m && cntArr[l-1][1] <= n {
		mem[[3]int{l, m, n}] = max(1+largestSubset(cntArr, m-cntArr[l-1][0], n-cntArr[l-1][1], l-1, mem), largestSubset(cntArr, m, n, l-1, mem))
		return mem[[3]int{l, m, n}]
	}
	// if cntArr[l][0] > m || cntArr[l][1] > n{
	//     return largestSubset(cntArr,m,n,l-1)
	// }
	mem[[3]int{l, m, n}] = largestSubset(cntArr, m, n, l-1, mem)
	return mem[[3]int{l, m, n}]

}

func findMaxForm(strs []string, m int, n int) int {

	zoArr := make([][2]int, len(strs))
	cntZero, cntOne := 0, 0

	for idx, str := range strs {
		// fmt.Println(str)
		for _, rune := range str {
			if rune == '0' {
				cntZero += 1
			} else if rune == '1' {
				cntOne += 1
			}
		}

		zoArr[idx][0], zoArr[idx][1] = cntZero, cntOne
		cntZero, cntOne = 0, 0
	}

	fmt.Println(zoArr)

	/*

	   [[1 1] [3 1] [2 4] [0 1] [1 0]]

	   [1 0]
	   - 5 0s and 3 1s -> 4 0s and 3 1s
	   -

	   if n==0 and m<=0 and n<=0
	       return 0
	   if elemnt[0]<=m and elemnt[1]<=n
	       take an element -> max(1+(further operation) + (further operations))
	   else
	       return (further operations n-1)


	*/

	mem := make(map[[3]int]int, 0)
	return largestSubset(zoArr, m, n, len(strs), mem)
}