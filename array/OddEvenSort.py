class Solution:
    def sortArrayByParityII(self, A ) :
        even = 0
        odd = 1
        res = [0]* len(A)

        for i in range(len(A)):
            if A[i] %2 ==0:
                res[even] = A[i]
                even +=2
            else:
                res[odd] = A[i]
                odd +=2
            
        return res

class Solution_Two_Pointer:
    pass

if __name__ == "__main__":
    s = Solution()
    print(s.sortArrayByParityII([4,2,5,7]))