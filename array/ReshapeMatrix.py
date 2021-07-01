class Solution:
    def matrixReshape(self, nums, r , c) :
        eles = [j for i in nums for j in i]
        if r * c != len(eles):
            return nums

        res = []
        for i in range(r):
            t =[]
            for j in range(c):
                t.append(eles.pop(0))
            res.append(t)
        
        return res

if __name__ == "__main__":
    s = Solution()
    print(s.matrixReshape([[1,2],[3,4]],1,4))