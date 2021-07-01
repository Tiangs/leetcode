class Solution:
    def longestOnes(self, A , K: int) -> int:
        res = 0
        left , right = 0, 0
        zeros = 0

        while right < len(A):
            if A[right] == 0 :
                zeros +=1
            
            while zeros > K:
                if A[left] == 0 :

                    # revert the variable zero
                    zeros -=1

                left +=1
            
            res = max(res,right - left +1)
            right +=1
    
        return  res
        


if __name__ == "__main__":
    s = Solution()
    print(s.longestOnes([1,1,1,0,0,0,1,1,1,1,0],2))