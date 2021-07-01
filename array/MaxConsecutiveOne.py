class Solution:
    def findMaxConsecutiveOnes(self, nums) -> int:
        maxCount =0 
        count =0

        for i in range(len(nums)):
            if nums[i] ==1:
                count +=1
                maxCount = max(maxCount,count)
            else:
                count =0
            
        return maxCount

if __name__ == "__main__":
    s = Solution()
    print(s.findMaxConsecutiveOnes([1,1,0,0,1,1]))