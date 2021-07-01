class Solution:
    def dailyTemperatures(self, T):

        stack = []
        res = [0] * len(T)
        print(res)

        for i in range(len(T)):
            print(i)

            while stack and T[i] > T[stack[-1]]:
                prev = stack.pop()
                res[prev]  = i-prev
            stack.append(i)

        return res

if __name__ == "__main__":
    s = Solution()
    print(s.dailyTemperatures([73, 74, 75, 71, 69, 72, 76, 73]))

