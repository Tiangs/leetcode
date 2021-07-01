package main

import "fmt"

func main() {
	mat := [][]byte{
		{'0', '1', '1'},
		{'0', '1', '1'},
	}
	fmt.Println(maximalSquare(mat))

}

func maximalSquare(matrix [][]byte) int {
	maxCount := 0

	M, N := len(matrix), len(matrix[0])

	//init dp[][]
	dp := make([][]int, M) // initialize a slice of dy slices
	for i := 0; i < M; i++ {
		dp[i] = make([]int, N) // initialize a slice of dx unit8 in each of dy slices
	}

	for i := 0; i < M; i++ {

		for j := 0; j < N; j++ {
			//1 row or 1 cols, return 1
			if i == 0 || j == 0 {
				dp[i][j] = int(matrix[i][j] - '0')
                maxCount = max(maxCount, dp[i][j])
				continue
			}

			//dp[i][j] = min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1] )+1
			if matrix[i][j] == '0' {
				dp[i][j] = 0
			} else {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}

			maxCount = max(maxCount, dp[i][j])

		}

	}


	//1 row 1 col remains as it is

	return maxCount * maxCount
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
