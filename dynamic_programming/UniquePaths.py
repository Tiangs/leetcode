class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        if m == 0 and n ==0 :
            return 0
        if m == 1 or n== 1 :
            return 1
        
        dp = [[0]*n for _ in range(m)]
        dp[0][0] = 0
        dp[0][1] = 1
        dp[1][0] = 1

        # fill in the dp array
        # each point is the sum of the left point([i][j-1]) and the upper point([i-1][j])
        for i in range(m):
            for j in range(n):
                up   = 0
                left = 0
                
                if (dp[i][j] != 0): # for the first 3 points already initialized 
                    continue
                if (i ==0 and j==0): # for the fist point 
                    continue

                elif i ==0:# if first row
                    left = dp[i][j-1]
                    dp[i][j] = left
                elif j ==0: # if first column
                    up = dp[i-1][j]
                    dp[i][j] = up
                else: # general case. 
                    up = dp[i-1][j]
                    left = dp[i][j-1]
                    dp[i][j] = up + left

        # return the finish point 
        return dp[m-1][n-1]

       

        



if __name__ == '__main__':
    s = Solution()
    r = s.uniquePaths(3,7)
    print(r)