package main

import "fmt"

//     ''a b c d e
// 	0 0 0 0 0 0
// a   0 1 1 1 1 1
// c.  0 1 1 2 2 2.
// e.  0 1 1 2 2 3

//     a b c
// d.  0 0 0
// e.  0 0 0
// f.  0 0 0

//     a b c
// a.  1 1 1
// b.  1 2 2
// c.  1 2 3

func main() {
	// fmt.Println(longestCommonSubsequence("bsbininm", "jmjkbkjkv"))
	fmt.Println(longestCommonSubsequence("abcde", "ace"))

}

func longestCommonSubsequence(text1 string, text2 string) int {

	n, m := len([]byte(text1))+1, len([]byte(text2))+1

	if n == 1 || m == 1 {
		return 0
	}

	//init dp matrix
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	fmt.Println(dp)

	for i := 0; i < m; i++ {

		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
				continue
			}

			if text1[j-1] == text2[i-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}

		}
	}

	fmt.Println(dp)
	return dp[m-1][n-1]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}
