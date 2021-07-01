class Solution:
    def maximumWealth(self, accounts) -> int:
        res = []
        for i in range(len(accounts)):

            sum = 0
            for j in range(len(accounts[i])):
                sum += accounts[i][j]
            
            res.append(sum)

        return max(res)


        