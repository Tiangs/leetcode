class Solution:
    def generate(self, numRows: int) :
        res =[]
        for i in range(numRows):
            if i == 0:
                res.append([1])
                continue
            if i == 1:
                res.append([1,1])
                continue
            t = [0] * (i+1)
            t[0] =1
            for j in range(1,i):
                t[j] = res[i-1][j-1] + res[i-1][j]
            t[-1] = 1
            res.append(t)
        return res

if __name__ == "__main__":
    s = Solution()
    print(s.generate(5))
        