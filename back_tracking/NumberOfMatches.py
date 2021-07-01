# back tracking ? 
class Solution:
    def numberOfMatches(self, n: int) -> int:
        res =0
        def bt(n,res):
            if n==0 or n==1:
                return 0,res

            if n%2 ==0:
                res += n/2
            else:
                res += (n-1)/2+1
            
            return bt(n//2,res)
        
        return bt(n,res)[1]
            
if __name__ == "__main__":
    s = Solution()
    print(s.numberOfMatches(14))