class Solution:
    def maxSatisfied(self, customers , grumpy , X: int) -> int:
        if X == len(customers):
            return sum(customers)

        total_val = 0
        for i,j in zip(customers,grumpy):
            if j == 0:
                total_val += i
        max_retrieve = 0

        i = 0
        while i < len(customers)-X+1:
            j = i
            max_retrieve_temp = 0
            # while j < i+X:
            for j in range(i,i+X):
                if grumpy[j] ==1 :
                    max_retrieve_temp  += customers[j]
                # j+=1
            
            max_retrieve = max(max_retrieve,max_retrieve_temp)
            i+=1

        return total_val +max_retrieve


class Official_Solution:

    def maxSatisfied(self, customers: List[int], grumpy: List[int], X: int) -> int:
        n = len(customers)
        total = sum(c for c, g in zip(customers, grumpy) if g == 0)
        maxIncrease = increase = sum(c * g for c, g in zip(customers[:X], grumpy[:X]))
        for i in range(X, n):
            increase += customers[i] * grumpy[i] - customers[i - X] * grumpy[i - X]
            maxIncrease = max(maxIncrease, increase)
        return total + maxIncrease


if __name__ == "__main__":
    s = Solution()
    print(s.maxSatisfied([1,0,1,2,1,1,7,5],[0,1,0,1,0,1,0,1],3))
    print(s.maxSatisfied([1],[0],1))
    print(s.maxSatisfied([2,6,6,9],[0,0,1,1],1))

