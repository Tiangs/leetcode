
class Solution:
    def reverse(self, x: int) -> int:

        if x<-2**31 or x > 2**31:
            return 0
        s = str(abs(x))

        stack = []
        res = []
        for i in range(len(s)):
            stack.append(s[i])
        
        while stack:
            res.append(stack.pop())    

        final_res = int(''.join(res))

        if x < 0 :
            return final_res * -1
        else:
            return final_res

