class Solution:
    def climbStairs(self, n: int) -> int:
        dp = [0] * (n+1)

        for i in range(n+1):
            
            if i ==0:
                dp[i] = 0
            elif i ==1:
                dp[i] =1
                print(dp)
            elif i ==2:
                dp[i] =2
            else:
                dp[i] = dp[i-1] + dp[i-2]
                print(dp)
            
        print(dp)
        return dp[-1]
        

if __name__ == "__main__":
    s = Solution()
    print(s.climbStairs(3))
            