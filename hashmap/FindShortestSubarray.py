class Solution:
    def findShortestSubArray(self, nums) -> int:
        m = {}

        for i in range(len(nums)):
            if nums[i] not in m.keys():
                #init
                m[nums[i]] = [i,i,1]
                continue
                
            m[nums[i]][1]  = i
            m[nums[i]][2] +=1
        maxFreq = sorted(m.items(),key = lambda x : (-x[1][-1],x[1][1] - x[1][0]) )[0]
        return maxFreq[1][1] - maxFreq[1][0] + 1

if __name__ == "__main__":

    s = Solution()
    print(s.findShortestSubArray([2,1]))