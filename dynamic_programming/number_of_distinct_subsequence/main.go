package main
import "fmt"


func numDistinct(s, t string) int {
    m, n := len(t), len(s)
    // if m < n {
    //     return 0
    // }
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
        // dp[i][n] = 1

    }

	//first line should be 1
	for i:= range dp[0]{
		dp[0][i]=1
	}
	// fmt.Println(dp)

	for i:=1;i<m+1;i++{
		for j:= 0;j<n+1;j++{
			if j == 0{
				dp[i][j] = 0
				continue

			}

			if t[i-1] != s[j-1]{
				dp[i][j] = dp[i][j-1]

			}else{
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			}

		}
	}
	// fmt.Println(dp)
	return dp[m][n]
 
}
func main(){
	fmt.Println(numDistinct("rabbbit","rabbit"))


}