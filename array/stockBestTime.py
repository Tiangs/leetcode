class Solution:
    def maxProfit(self, prices) -> int:
        max_val = 0
        n = len(prices)

        i = 0
        while i < n-1:
            j = i + 1

            while j < n:
                if prices[j] <= prices[i]:
                    i = j
                    break
                
                t = prices[j] - prices[i]
                if t > max_val:
                    max_val = t

                if j == n-1:
                    return max_val
                j+=1
            
            

        return max_val